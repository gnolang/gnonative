package gnonative

import (
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"

	mdb "github.com/gnolang/gno/tm2/pkg/db"
)

var errShort = fmt.Errorf("chunk blob: short buffer")

// NativeDB is implemented in the native (Kotlin/Swift) layer.
type NativeDB interface {
	Get([]byte) []byte
	Has(key []byte) bool
	Set([]byte, []byte)
	SetSync([]byte, []byte)
	Delete([]byte)
	DeleteSync([]byte)
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

// decodeChunkBlob parses a single binary blob produced by NativeDB.ScanChunk.
//
// Blob layout (all integers are big-endian):
//
//	+---------+-------------------+---------------------------------------+--------------------------+------------------------+
//	| Offset  | Field             | Description                           | Type/Size                | Notes                  |
//	+---------+-------------------+---------------------------------------+--------------------------+------------------------+
//	| 0       | flags             | bit0 = hasMore (1 => more pages)      | uint8 (1 byte)           | other bits reserved    |
//	| 1       | count             | number of K/V pairs that follow       | uint32 (4 bytes, BE)     | N                      |
//	| 5       | pairs[0..N-1]     | repeated K/V frames:                  |                          |                        |
//	|         |  - klen           | key length                            | uint32 (4 bytes, BE)     |                        |
//	|         |  - key            | key bytes                             | klen bytes               |                        |
//	|         |  - vlen           | value length                          | uint32 (4 bytes, BE)     |                        |
//	|         |  - value          | value bytes                           | vlen bytes               |                        |
//	| ...     | nextSeekLen       | length of the nextSeek key            | uint32 (4 bytes, BE)     | 0 if empty             |
//	| ...     | nextSeek          | nextSeek key bytes                    | nextSeekLen bytes        |                        |
//	+---------+-------------------+---------------------------------------+--------------------------+------------------------+
//
// Semantics:
//   - The iterator uses 'hasMore' to know if additional pages exist.
//   - 'nextSeek' is typically the last key of this page; pass it back as 'seekKey' (exclusive)
//     on the next ScanChunk call to continue from the next item.
//   - Keys/values are raw bytes; ordering and range checks are done on the raw key bytes.
//
// On decode errors (short buffer / lengths out of range), the function returns errShort.
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
