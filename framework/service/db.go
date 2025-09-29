package gnonative

import (
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"

	mdb "github.com/gnolang/gno/tm2/pkg/db"
)

// NativeDB is implemented in the native (Kotlin/Swift) layer.
type NativeDB interface {
	Get([]byte) []byte
	Has(key []byte) bool
	Set([]byte, []byte)
	SetSync([]byte, []byte)
	Delete([]byte)
	DeleteSync([]byte)
	// GetAllKeys() [][]byte // needed for iteration support
	ScanChunk(start, end, seekKey []byte, limit int, reverse bool) ([]byte, error)
}

type db struct {
	NativeDB

	closed atomic.Bool
	mu     sync.RWMutex // used for optional safety around Stats/Print
}

func (d *db) Close() error {
	d.closed.Store(true)
	return nil
}

func (d *db) ensureOpen() {
	if d.closed.Load() {
		panic("db: use after Close")
	}
}

func (db *db) Iterator(start, end []byte) mdb.Iterator {
	db.ensureOpen()
	it := &iterator{
		db:         db,
		start:      append([]byte(nil), start...),
		end:        append([]byte(nil), end...),
		reverse:    false,
		chunkLimit: 256,
	}
	it.fill()
	return it
}

func (db *db) ReverseIterator(start, end []byte) mdb.Iterator {
	db.ensureOpen()
	it := &iterator{
		db:         db,
		start:      append([]byte(nil), start...),
		end:        append([]byte(nil), end...),
		reverse:    true,
		chunkLimit: 256,
	}
	it.fill()
	return it
}

// Iterator creates a forward iterator over a domain of keys
// func (d *db) Iterator(start, end []byte) mdb.Iterator {
// 	d.ensureOpen()
// 	return d.createIterator(start, end, false)
// }

// Iterator creates a forward iterator over a domain of keys
// func (d *db) ReverseIterator(start, end []byte) mdb.Iterator {
// 	d.ensureOpen()
// 	return d.createIterator(start, end, true)
// }

func (d *db) Print() {
	d.mu.RLock()
	defer d.mu.RUnlock()
	// With only point ops exposed, we can't enumerate keys here.
	// Keep as a stub or log something useful for debugging:
	fmt.Println("db.Print(): NativeDB has no range API; nothing to print")
}

func (d *db) Stats() map[string]string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	// Return whatever you track on the Go side. Native side has no stats here.
	return map[string]string{
		"closed": fmt.Sprintf("%v", d.closed.Load()),
	}
}

func (d *db) NewBatch() mdb.Batch {
	d.ensureOpen()
	return &batch{db: d}
}

// --- Batch implementation (pure Go) ---

type batch struct {
	db  *db
	ops []op
}

type op struct {
	del   bool
	key   []byte
	value []byte // nil for delete
}

func (b *batch) Set(key, value []byte) {
	b.db.ensureOpen()
	// Defensive copies: gomobile & callers must not mutate later.
	k := append([]byte(nil), key...)
	v := append([]byte(nil), value...)
	b.ops = append(b.ops, op{del: false, key: k, value: v})
}

func (b *batch) Delete(key []byte) {
	b.db.ensureOpen()
	k := append([]byte(nil), key...)
	b.ops = append(b.ops, op{del: true, key: k})
}

// Write applies ops using async variants.
func (b *batch) Write() {
	b.db.ensureOpen()
	for _, o := range b.ops {
		if o.del {
			b.db.Delete(o.key)
		} else {
			b.db.Set(o.key, o.value)
		}
	}
	// Clear buffer to allow reuse if desired.
	b.ops = b.ops[:0]
}

// WriteSync applies ops using sync variants.
func (b *batch) WriteSync() {
	b.db.ensureOpen()
	for _, o := range b.ops {
		if o.del {
			b.db.DeleteSync(o.key)
		} else {
			b.db.SetSync(o.key, o.value)
		}
	}
	b.ops = b.ops[:0]
}

func (b *batch) Close() {
	// Drop buffered ops; allow GC.
	b.ops = nil
	b.db = nil
}

type kv struct {
	k []byte
	v []byte
}

type iterator struct {
	db         *db
	start      []byte
	end        []byte
	reverse    bool
	seekKey    []byte
	chunk      []kv
	i          int
	hasMore    bool
	closed     bool
	chunkLimit int
}

func (it *iterator) Domain() (start, end []byte) { return it.start, it.end }

func (it *iterator) Valid() bool {
	if it.closed {
		return false
	}
	for it.i >= len(it.chunk) && it.hasMore {
		it.fill()
	}
	return it.i < len(it.chunk)
}

func (it *iterator) Next() {
	if !it.Valid() {
		return
	}
	cur := it.chunk[it.i]
	it.i++
	// keep seekKey strictly at the last returned key
	it.seekKey = append(it.seekKey[:0], cur.k...)
}

func (it *iterator) Key() []byte {
	if !it.Valid() {
		return nil
	}
	return it.chunk[it.i].k
}

func (it *iterator) Value() []byte {
	if !it.Valid() {
		return nil
	}
	return it.chunk[it.i].v
}
func (it *iterator) Close() { it.closed = true; it.chunk = nil }

func (it *iterator) fill() {
	if it.closed {
		return
	}
	blob, err := it.db.ScanChunk(it.start, it.end, it.seekKey, it.chunkLimit, it.reverse)
	if err != nil {
		it.chunk, it.i, it.hasMore = nil, 0, false
		return
	}
	pairs, nextSeek, hasMore, err := decodeChunkBlob(blob)
	if err != nil {
		it.chunk, it.i, it.hasMore = nil, 0, false
		return
	}
	it.chunk = pairs
	it.i = 0
	it.hasMore = hasMore
	if len(nextSeek) > 0 {
		it.seekKey = append(it.seekKey[:0], nextSeek...)
	}
}

// --- framing decode ---

func decodeChunkBlob(b []byte) (pairs []kv, nextSeek []byte, hasMore bool, err error) {
	if len(b) < 1+4 {
		return nil, nil, false, errShort
	}
	flags := b[0]
	hasMore = (flags & 0x01) != 0
	b = b[1:]

	count := int(binary.BigEndian.Uint32(b[:4]))
	b = b[4:]

	pairs = make([]kv, 0, count)
	for i := 0; i < count; i++ {
		if len(b) < 4 {
			return nil, nil, false, errShort
		}
		klen := int(binary.BigEndian.Uint32(b[:4]))
		b = b[4:]
		if klen < 0 || len(b) < klen {
			return nil, nil, false, errShort
		}
		k := append([]byte(nil), b[:klen]...)
		b = b[klen:]

		if len(b) < 4 {
			return nil, nil, false, errShort
		}
		vlen := int(binary.BigEndian.Uint32(b[:4]))
		b = b[4:]
		if vlen < 0 || len(b) < vlen {
			return nil, nil, false, errShort
		}
		v := append([]byte(nil), b[:vlen]...)
		b = b[vlen:]

		pairs = append(pairs, kv{k: k, v: v})
	}

	if len(b) < 4 {
		return nil, nil, false, errShort
	}
	nlen := int(binary.BigEndian.Uint32(b[:4]))
	b = b[4:]
	if nlen < 0 || len(b) < nlen {
		return nil, nil, false, errShort
	}
	if nlen > 0 {
		nextSeek = append([]byte(nil), b[:nlen]...)
	}
	return pairs, nextSeek, hasMore, nil
}

var errShort = fmt.Errorf("chunk blob: short buffer")

// Old implementation of Iterator (in-memory, inefficient).

// type dbIterator struct {
// 	db      *db
// 	keys    [][]byte
// 	index   int
// 	start   []byte
// 	end     []byte
// 	reverse bool
// 	valid   bool
// }
//
// func (d *db) createIterator(start, end []byte, reverse bool) mdb.Iterator {
// 	// Get all keys from the native database
// 	allKeys := d.NativeDB.GetAllKeys()
//
// 	// Filter keys within the domain
// 	var filteredKeys [][]byte
// 	for _, key := range allKeys {
// 		if d.keyInDomain(key, start, end) {
// 			filteredKeys = append(filteredKeys, key)
// 		}
// 	}
//
// 	// Sort keys
// 	sort.Slice(filteredKeys, func(i, j int) bool {
// 		if reverse {
// 			return bytes.Compare(filteredKeys[i], filteredKeys[j]) > 0
// 		}
// 		return bytes.Compare(filteredKeys[i], filteredKeys[j]) < 0
// 	})
//
// 	iterator := &dbIterator{
// 		db:      d,
// 		keys:    filteredKeys,
// 		index:   0,
// 		start:   copyBytes(start),
// 		end:     copyBytes(end),
// 		reverse: reverse,
// 		valid:   len(filteredKeys) > 0,
// 	}
//
// 	return iterator
// }
//
// func (d *db) keyInDomain(key, start, end []byte) bool {
// 	// Handle nil start (empty byteslice)
// 	if start != nil && bytes.Compare(key, start) < 0 {
// 		return false
// 	}
//
// 	// Handle nil end (no upper limit)
// 	if end != nil && bytes.Compare(key, end) >= 0 {
// 		return false
// 	}
//
// 	return true
// }
//
// // Domain returns the start and end limits of the iterator
// func (it *dbIterator) Domain() (start []byte, end []byte) {
// 	return it.start, it.end
// }
//
// // Valid returns whether the current position is valid
// func (it *dbIterator) Valid() bool {
// 	return it.valid && it.index >= 0 && it.index < len(it.keys)
// }
//
// // Next moves the iterator to the next sequential key
// func (it *dbIterator) Next() {
// 	if !it.Valid() {
// 		panic("iterator is not valid")
// 	}
//
// 	it.index++
// 	if it.index >= len(it.keys) {
// 		it.valid = false
// 	}
// }
//
// // Key returns the key of the cursor
// func (it *dbIterator) Key() []byte {
// 	if !it.Valid() {
// 		panic("iterator is not valid")
// 	}
//
// 	return it.keys[it.index]
// }
//
// // Value returns the value of the cursor
// func (it *dbIterator) Value() []byte {
// 	if !it.Valid() {
// 		panic("iterator is not valid")
// 	}
//
// 	key := it.keys[it.index]
// 	return it.db.Get(key)
// }
//
// // Close releases the Iterator
// func (it *dbIterator) Close() {
// 	it.valid = false
// 	it.keys = nil
// }

// Helper function to copy byte slices
func copyBytes(src []byte) []byte {
	if src == nil {
		return nil
	}
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}
