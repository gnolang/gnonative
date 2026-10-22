package service

import (
	"errors"
	"testing"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// The response of the "vm/qeval_json" query for "gno.land/r/gnoland/boards2/v0.GetBoard(1)".
const testQEvalJSONBoard = `{"results":[{"T":{"@type":"/gno.RefType","ID":"gno.land/p/gnoland/boards/exts/hub/v0.Board"},"V":{"@type":"/gno.StructValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"},"Fields":[{"T":{"@type":"/gno.PrimitiveType","value":"65536"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"openboard1"}},{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"},"Vrd":false}},{"T":{"@type":"/gno.PrimitiveType","value":"4"}},{"T":{"@type":"/gno.PrimitiveType","value":"32"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"32"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.RefType","ID":".uverse.address"},"V":{"@type":"/gno.StringValue","value":"g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5"}},{"T":{"@type":"/gno.PrimitiveType","value":"1024"},"N":"5v2zagAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"1024"}}]}},{"T":{"@type":"/gno.PrimitiveType","value":"4"},"N":"AQAAAAAAAAA="}]}`

// The responses of the "vm/qtype_json" query by type ID.
var testQTypeJSON = map[string]string{
	"gno.land/p/gnoland/boards/exts/hub/v0.Board": `{"typeid":"gno.land/p/gnoland/boards/exts/hub/v0.Board","type":{"@type":"/gno.DeclaredType","PkgPath":"gno.land/p/gnoland/boards/exts/hub/v0","Name":"Board","Base":{"@type":"/gno.StructType","Fields":[{"Name":"id","Type":{"@type":"/gno.PrimitiveType","value":"65536"}},{"Name":"name","Type":{"@type":"/gno.PrimitiveType","value":"16"}},{"Name":"aliases","Type":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"}}},{"Name":"readonly","Type":{"@type":"/gno.PrimitiveType","value":"4"}},{"Name":"threadCount","Type":{"@type":"/gno.PrimitiveType","value":"32"}},{"Name":"memberCount","Type":{"@type":"/gno.PrimitiveType","value":"32"}},{"Name":"creator","Type":{"@type":"/gno.DeclaredType","PkgPath":".uverse","Name":"address","Base":{"@type":"/gno.PrimitiveType","value":"16"}}},{"Name":"createdAt","Type":{"@type":"/gno.PrimitiveType","value":"1024"}},{"Name":"updatedAt","Type":{"@type":"/gno.PrimitiveType","value":"1024"}}]}}}`,
	".uverse.address":                `{"typeid":".uverse.address","type":{"@type":"/gno.DeclaredType","PkgPath":".uverse","Name":"address","Base":{"@type":"/gno.PrimitiveType","value":"16"}}}`,
	"gno.land/r/demo/boards.BoardID": `{"typeid":"gno.land/r/demo/boards.BoardID","type":{"@type":"/gno.DeclaredType","PkgPath":"gno.land/r/demo/boards","Name":"BoardID","Base":{"@type":"/gno.PrimitiveType","value":"65536"}}}`,
	"gno.land/r/sys/users.UserData":  `{"typeid":"gno.land/r/sys/users.UserData","type":{"@type":"/gno.DeclaredType","PkgPath":"gno.land/r/sys/users","Name":"UserData","Base":{"@type":"/gno.StructType","Fields":[{"Name":"addr","Type":{"@type":"/gno.DeclaredType","PkgPath":".uverse","Name":"address","Base":{"@type":"/gno.PrimitiveType","value":"16"}}},{"Name":"username","Type":{"@type":"/gno.PrimitiveType","value":"16"}},{"Name":"deleted","Type":{"@type":"/gno.PrimitiveType","value":"4"}}]}}}`,
}

// The responses of the "vm/qobject_json" query by object ID.
var testQObjectJSON = map[string]string{
	// A gno.land/r/sys/users.UserData
	"0ed754d80e40da2645666583e74d4ab093970a6a:535": `{"objectid":"0ed754d80e40da2645666583e74d4ab093970a6a:535","value":{"@type":"/gno.StructValue","ObjectInfo":{"ID":"0ed754d80e40da2645666583e74d4ab093970a6a:535","Hash":"c71b7fdf1ea86d0d240dbee17e52c795e78ee716","OwnerID":"0ed754d80e40da2645666583e74d4ab093970a6a:534","ModTime":"0","RefCount":"1","LastObjectSize":"396"},"Fields":[{"T":{"@type":"/gno.RefType","ID":".uverse.address"},"V":{"@type":"/gno.StringValue","value":"g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5"}},{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"test1"}},{"T":{"@type":"/gno.PrimitiveType","value":"4"}}]}}`,
	// The array of a []string{"a", "b", "c"}
	"0dc5eff18974d66f6720bc387fc7e5c63514d961:21": `{"objectid":"0dc5eff18974d66f6720bc387fc7e5c63514d961:21","value":{"@type":"/gno.ArrayValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:21","ModTime":"0","RefCount":"1","LastObjectSize":"0"},"List":[{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"a"}},{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"b"}},{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"c"}}],"Data":null}}`,
	// A heap item which only has a ref to another object
	"0dc5eff18974d66f6720bc387fc7e5c63514d961:23": `{"objectid":"0dc5eff18974d66f6720bc387fc7e5c63514d961:23","value":{"@type":"/gno.HeapItemValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:23","ModTime":"0","RefCount":"1","LastObjectSize":"0"},"Value":{"T":{"@type":"/gno.RefType","ID":"gno.land/r/foo.Node"},"V":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:24","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"}}}}`,
}

// testQuery returns a fake of the client's Query for the "vm/qtype_json" and "vm/qobject_json" queries, using
// testQTypeJSON and testQObjectJSON. Add the data of each query (the type ID or object ID) to queried (if not nil).
func testQuery(t *testing.T, queried *[]string) func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error) {
	return func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error) {
		if queried != nil {
			*queried = append(*queried, string(cfg.Data))
		}
		var responses map[string]string
		switch cfg.Path {
		case "vm/qtype_json":
			responses = testQTypeJSON
		case "vm/qobject_json":
			responses = testQObjectJSON
		default:
			t.Errorf("unexpected query path %q", cfg.Path)
		}
		res, ok := responses[string(cfg.Data)]
		if !ok {
			return nil, errors.New("not found: " + string(cfg.Data))
		}
		return &ctypes.ResultABCIQuery{Response: abci.ResponseQuery{ResponseBase: abci.ResponseBase{Data: []byte(res)}}}, nil
	}
}

// testConvertQEvalJSON converts data, using testQTypeJSON for the type queries. Add each queried type ID to queried.
func testConvertQEvalJSON(t *testing.T, data string, queried *[]string) *api_gen.QEvalJSONValuesResponse {
	t.Helper()
	converter := newGnoValueConverter(testQuery(t, queried), zap.NewNop())
	res, err := converter.convertQEvalJSON([]byte(data))
	require.NoError(t, err)
	return res
}

// assertGnoScalar asserts that the value has the type, kind and scalar.
func assertGnoScalar(t *testing.T, v *api_gen.GnoValue, typ string, kind string, scalar string) {
	t.Helper()
	assert.Equal(t, typ, v.Type)
	assert.Equal(t, kind, v.Kind)
	assert.Equal(t, scalar, v.Scalar)
}

func TestConvertQEvalJSONStruct(t *testing.T) {
	var queried []string
	res := testConvertQEvalJSON(t, testQEvalJSONBoard, &queried)

	require.Len(t, res.Results, 2)
	assert.Empty(t, res.Error)
	assertGnoScalar(t, res.Results[1], "bool", gnoKindBool, "true")

	board := res.Results[0]
	assert.Equal(t, "gno.land/p/gnoland/boards/exts/hub/v0.Board", board.Type)
	assert.Equal(t, gnoKindStruct, board.Kind)
	require.Len(t, board.Fields, 9)

	names := make([]string, len(board.Fields))
	for i, field := range board.Fields {
		names[i] = field.Name
	}
	assert.Equal(t, []string{"id", "name", "aliases", "readonly", "threadCount", "memberCount", "creator", "createdAt", "updatedAt"}, names)

	assertGnoScalar(t, board.Fields[0], "uint64", gnoKindNumber, "1")
	assertGnoScalar(t, board.Fields[1], "string", gnoKindString, "openboard1")
	// A nil slice
	assertGnoScalar(t, board.Fields[2], "[]string", gnoKindList, "")
	assert.Empty(t, board.Fields[2].Elems)
	assertGnoScalar(t, board.Fields[3], "bool", gnoKindBool, "false")
	assertGnoScalar(t, board.Fields[4], "int", gnoKindNumber, "1")
	assertGnoScalar(t, board.Fields[6], "address", gnoKindString, "g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5")
	assertGnoScalar(t, board.Fields[7], "int64", gnoKindNumber, "1790180838")
	assertGnoScalar(t, board.Fields[8], "int64", gnoKindNumber, "0")

	// The address already has its string value, so only the struct type is queried.
	assert.Equal(t, []string{"gno.land/p/gnoland/boards/exts/hub/v0.Board"}, queried)
}

func TestConvertQEvalJSONDeclaredPrimitive(t *testing.T) {
	// The values are in N, which needs the base type to interpret: a BoardID (uint64) of max uint64,
	// an empty address and a float64.
	var queried []string
	res := testConvertQEvalJSON(t, `{"results":[`+
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/r/demo/boards.BoardID"},"N":"//////////8="},`+
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/r/demo/boards.BoardID"},"N":"AgAAAAAAAAA="},`+
		`{"T":{"@type":"/gno.RefType","ID":".uverse.address"}},`+
		`{"T":{"@type":"/gno.PrimitiveType","value":"262144"},"N":"AAAAAAAA+D8="},`+
		`{"T":{"@type":"/gno.PrimitiveType","value":"64"},"N":"/wAAAAAAAAA="}]}`, &queried)

	require.Len(t, res.Results, 5)
	assertGnoScalar(t, res.Results[0], "gno.land/r/demo/boards.BoardID", gnoKindNumber, "18446744073709551615")
	assertGnoScalar(t, res.Results[1], "gno.land/r/demo/boards.BoardID", gnoKindNumber, "2")
	assertGnoScalar(t, res.Results[2], "address", gnoKindString, "")
	assertGnoScalar(t, res.Results[3], "float64", gnoKindNumber, "1.5")
	assertGnoScalar(t, res.Results[4], "int8", gnoKindNumber, "-1")
	// Each type is only queried once.
	assert.Equal(t, []string{"gno.land/r/demo/boards.BoardID", ".uverse.address"}, queried)
}

func TestConvertQEvalJSONContainers(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"}`
	const stringType = `{"@type":"/gno.PrimitiveType","value":"16"}`
	const intType = `{"@type":"/gno.PrimitiveType","value":"32"}`
	res := testConvertQEvalJSON(t, `{"results":[`+
		// []byte("hi")
		`{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"4096"},"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":null,"Data":"aGk="},"Offset":"0","Length":"2","Maxcap":"2"}},`+
		// []string{"a","b","c"}[1:2]
		`{"T":{"@type":"/gno.SliceType","Elt":`+stringType+`,"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":[{"T":`+stringType+`,"V":{"@type":"/gno.StringValue","value":"a"}},{"T":`+stringType+`,"V":{"@type":"/gno.StringValue","value":"b"}},{"T":`+stringType+`,"V":{"@type":"/gno.StringValue","value":"c"}}],"Data":null},"Offset":"1","Length":"1","Maxcap":"2"}},`+
		// map[string]int{"a":1}
		`{"T":{"@type":"/gno.MapType","Key":`+stringType+`,"Value":`+intType+`},"V":{"@type":"/gno.MapValue",`+objectInfo+`,"List":{"List":[{"Key":{"T":`+stringType+`,"V":{"@type":"/gno.StringValue","value":"a"}},"Value":{"T":`+intType+`,"N":"AQAAAAAAAAA="}}]}}},`+
		// A pointer to a persisted object
		`{"T":{"@type":"/gno.PointerType","Elt":{"@type":"/gno.RefType","ID":"gno.land/p/gnoland/boards/v0.storage"}},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:16","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"},"Index":"0"}},`+
		// A pointer to a new int
		`{"T":{"@type":"/gno.PointerType","Elt":`+intType+`},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.HeapItemValue",`+objectInfo+`,"Value":{"T":`+intType+`,"N":"BwAAAAAAAAA="}},"Index":"0"}},`+
		// A nil pointer and a nil []byte
		`{"T":{"@type":"/gno.PointerType","Elt":`+intType+`}},`+
		`{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"4096"},"Vrd":false}},`+
		// An untyped nil
		`{}]}`, nil)

	require.Len(t, res.Results, 8)

	assert.Equal(t, "[]uint8", res.Results[0].Type)
	assert.Equal(t, gnoKindBytes, res.Results[0].Kind)
	assert.Equal(t, []byte("hi"), res.Results[0].Data)

	assert.Equal(t, "[]string", res.Results[1].Type)
	assert.Equal(t, gnoKindList, res.Results[1].Kind)
	require.Len(t, res.Results[1].Elems, 1)
	assertGnoScalar(t, res.Results[1].Elems[0], "string", gnoKindString, "b")

	assert.Equal(t, "map[string]int", res.Results[2].Type)
	assert.Equal(t, gnoKindMap, res.Results[2].Kind)
	require.Len(t, res.Results[2].Keys, 1)
	require.Len(t, res.Results[2].Elems, 1)
	assertGnoScalar(t, res.Results[2].Keys[0], "string", gnoKindString, "a")
	assertGnoScalar(t, res.Results[2].Elems[0], "int", gnoKindNumber, "1")

	assert.Equal(t, "*gno.land/p/gnoland/boards/v0.storage", res.Results[3].Type)
	assert.Equal(t, gnoKindPointer, res.Results[3].Kind)
	assert.Equal(t, "0dc5eff18974d66f6720bc387fc7e5c63514d961:16", res.Results[3].ObjectId)
	assert.Equal(t, "gno.land/p/gnoland/boards/v0.storage", res.Results[3].TypeId)
	assert.Empty(t, res.Results[3].Elems)

	assert.Equal(t, "*int", res.Results[4].Type)
	assert.Equal(t, gnoKindPointer, res.Results[4].Kind)
	assert.Empty(t, res.Results[4].ObjectId)
	require.Len(t, res.Results[4].Elems, 1)
	assertGnoScalar(t, res.Results[4].Elems[0], "int", gnoKindNumber, "7")

	assertGnoScalar(t, res.Results[5], "*int", gnoKindNil, "")
	assertGnoScalar(t, res.Results[6], "[]uint8", gnoKindBytes, "")
	assertGnoScalar(t, res.Results[7], "", gnoKindNil, "")
}

func TestConvertQEvalJSONError(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"}`
	// The type is not in testQTypeJSON.
	res := testConvertQEvalJSON(t, `{"results":[`+
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/r/foo.myErr"},"V":{"@type":"/gno.StructValue",`+objectInfo+`,"Fields":[{"T":{"@type":"/gno.PrimitiveType","value":"32"},"N":"AwAAAAAAAAA="}]}},`+
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/r/foo.Unknown"},"N":"AwAAAAAAAAA="}],`+
		`"@error":"boom"}`, nil)

	assert.Equal(t, "boom", res.Error)
	require.Len(t, res.Results, 2)

	// The struct type can't be resolved, so the field doesn't have a name.
	assert.Equal(t, "gno.land/r/foo.myErr", res.Results[0].Type)
	assert.Equal(t, gnoKindStruct, res.Results[0].Kind)
	require.Len(t, res.Results[0].Fields, 1)
	assert.Empty(t, res.Results[0].Fields[0].Name)
	assertGnoScalar(t, res.Results[0].Fields[0], "int", gnoKindNumber, "3")

	// We can't interpret the value without the type.
	assert.Equal(t, gnoKindUnknown, res.Results[1].Kind)
	assert.Equal(t, []byte{3, 0, 0, 0, 0, 0, 0, 0}, res.Results[1].Data)
}

func TestResolveTypeLogsQueryError(t *testing.T) {
	// The type is not in testQTypeJSON, so the query returns an error, which is logged.
	core, logs := observer.New(zap.DebugLevel)
	var queried []string
	converter := newGnoValueConverter(testQuery(t, &queried), zap.New(core))
	assert.Nil(t, converter.resolveType("gno.land/r/foo.Unknown"))

	require.Equal(t, 1, logs.Len())
	entry := logs.All()[0]
	assert.Equal(t, zap.DebugLevel, entry.Level)
	assert.Equal(t, "gno.land/r/foo.Unknown", entry.ContextMap()["typeId"])
	assert.Equal(t, "not found: gno.land/r/foo.Unknown", entry.ContextMap()["error"])

	// The failure is cached, so there is no other query or log.
	assert.Nil(t, converter.resolveType("gno.land/r/foo.Unknown"))
	assert.Equal(t, []string{"gno.land/r/foo.Unknown"}, queried)
	assert.Equal(t, 1, logs.Len())
}

func TestConvertQEvalJSONMalformed(t *testing.T) {
	converter := newGnoValueConverter(func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error) {
		return nil, errors.New("unexpected")
	}, zap.NewNop())

	_, err := converter.convertQEvalJSON([]byte(`not JSON`))
	assert.Error(t, err)

	// No results
	res, err := converter.convertQEvalJSON([]byte(`{"results":[]}`))
	require.NoError(t, err)
	assert.Empty(t, res.Results)
}

func TestConvertQEvalJSONRefs(t *testing.T) {
	const refValue = `{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"}`
	res := testConvertQEvalJSON(t, `{"results":[`+
		// A slice of a persisted array
		`{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"},"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":`+refValue+`,"Offset":"2","Length":"3","Maxcap":"8"}},`+
		// A persisted struct
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/p/gnoland/boards/exts/hub/v0.Board"},"V":`+refValue+`}]}`, nil)

	require.Len(t, res.Results, 2)

	assert.Equal(t, "[]string", res.Results[0].Type)
	assert.Equal(t, gnoKindRef, res.Results[0].Kind)
	assert.Equal(t, "0dc5eff18974d66f6720bc387fc7e5c63514d961:7", res.Results[0].ObjectId)
	assert.Equal(t, uint64(2), res.Results[0].Offset)
	assert.Equal(t, uint64(3), res.Results[0].Length)
	assert.Empty(t, res.Results[0].TypeId)

	assert.Equal(t, gnoKindRef, res.Results[1].Kind)
	assert.Equal(t, "0dc5eff18974d66f6720bc387fc7e5c63514d961:7", res.Results[1].ObjectId)
	assert.Equal(t, "gno.land/p/gnoland/boards/exts/hub/v0.Board", res.Results[1].TypeId)
}

func TestConvertQEvalJSONBackrefs(t *testing.T) {
	// Objects which are not persisted are numbered by the export in the order they are first expanded. An object
	// which is visited again is an ExportRefValue with the synthetic ID ":N".
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"}`
	const intType = `{"@type":"/gno.PrimitiveType","value":"32"}`
	const mapType = `{"@type":"/gno.MapType","Key":{"@type":"/gno.PrimitiveType","value":"16"},"Value":` + intType + `}`
	const sliceType = `{"@type":"/gno.SliceType","Elt":` + intType + `,"Vrd":false}`
	const nodeType = `{"@type":"/gno.RefType","ID":"gno.land/r/demo/foo.Node"}`
	const nodePointerType = `{"@type":"/gno.PointerType","Elt":` + nodeType + `}`
	res := testConvertQEvalJSON(t, `{"results":[`+
		// m := map[string]int{"a": 1}; return m, m (:1 is the map)
		`{"T":`+mapType+`,"V":{"@type":"/gno.MapValue",`+objectInfo+`,"List":{"List":[{"Key":{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"a"}},"Value":{"T":`+intType+`,"N":"AQAAAAAAAAA="}}]}}},`+
		`{"T":`+mapType+`,"V":{"@type":"/gno.ExportRefValue","ObjectID":":1"}},`+
		// s := []int{1, 2}; return s, s[1:] (:2 is the array)
		`{"T":`+sliceType+`,"V":{"@type":"/gno.SliceValue","Base":{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":[{"T":`+intType+`,"N":"AQAAAAAAAAA="},{"T":`+intType+`,"N":"AgAAAAAAAAA="}],"Data":null},"Offset":"0","Length":"2","Maxcap":"2"}},`+
		`{"T":`+sliceType+`,"V":{"@type":"/gno.SliceValue","Base":{"@type":"/gno.ExportRefValue","ObjectID":":2"},"Offset":"1","Length":"1","Maxcap":"1"}},`+
		// n := &Node{}; n.Next = n; return n (:3 is the heap item, :4 is the struct)
		`{"T":`+nodePointerType+`,"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.HeapItemValue",`+objectInfo+`,"Value":{"T":`+nodeType+`,"V":{"@type":"/gno.StructValue",`+objectInfo+`,"Fields":[`+
		`{"T":`+nodePointerType+`,"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.ExportRefValue","ObjectID":":3"},"Index":"0"}}]}}},"Index":"0"}}]}`, nil)

	require.Len(t, res.Results, 5)
	assert.Equal(t, gnoKindMap, res.Results[0].Kind)
	assert.Equal(t, gnoKindList, res.Results[2].Kind)
	assert.Equal(t, gnoKindPointer, res.Results[4].Kind)
	require.Len(t, res.Results[4].Elems, 1)
	node := res.Results[4].Elems[0]
	assert.Equal(t, gnoKindStruct, node.Kind)
	require.Len(t, node.Fields, 1)

	// The map, the slice of the array and the pointer to the heap item are backrefs, which can't be fetched.
	for _, v := range []*api_gen.GnoValue{res.Results[1], res.Results[3], node.Fields[0]} {
		assert.Equal(t, gnoKindBackref, v.Kind)
		assert.Empty(t, v.ObjectId)
		assert.Empty(t, v.TypeId)
	}
	assert.Equal(t, "map[string]int", res.Results[1].Type)
	assert.Equal(t, "[]int", res.Results[3].Type)
	assert.Equal(t, "*gno.land/r/demo/foo.Node", node.Fields[0].Type)
}

func TestConvertQEvalJSONEmptySliceRef(t *testing.T) {
	const refValue = `{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"}`
	// An empty slice of a persisted array doesn't need the array, so it is not a ref.
	res := testConvertQEvalJSON(t, `{"results":[`+
		`{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"},"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":`+refValue+`,"Offset":"2","Length":"0","Maxcap":"8"}},`+
		`{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"4096"},"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":`+refValue+`,"Offset":"0","Length":"0","Maxcap":"8"}}]}`, nil)

	require.Len(t, res.Results, 2)
	assertGnoScalar(t, res.Results[0], "[]string", gnoKindList, "")
	assert.Empty(t, res.Results[0].ObjectId)
	assert.Empty(t, res.Results[0].Elems)
	assertGnoScalar(t, res.Results[1], "[]uint8", gnoKindBytes, "")
	assert.Empty(t, res.Results[1].ObjectId)
}

func TestConvertQEvalJSONPointerIndex(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"}`
	const intType = `{"@type":"/gno.PrimitiveType","value":"32"}`
	res := testConvertQEvalJSON(t, `{"results":[`+
		// A pointer to the second field of a persisted struct
		`{"T":{"@type":"/gno.PointerType","Elt":{"@type":"/gno.PrimitiveType","value":"16"}},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"},"Index":"1"}},`+
		// A pointer to the second element of a byte array which is not persisted, like &buf[1]
		`{"T":{"@type":"/gno.PointerType","Elt":{"@type":"/gno.PrimitiveType","value":"4096"}},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":null,"Data":"aGk="},"Index":"1"}},`+
		// The special index PointerIndexMap, with a persisted map
		`{"T":{"@type":"/gno.PointerType","Elt":`+intType+`},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:8","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"},"Index":"-2"}},`+
		// The special index PointerIndexBlockBlank, with a struct which is not persisted
		`{"T":{"@type":"/gno.PointerType","Elt":`+intType+`},"V":{"@type":"/gno.PointerValue","TV":null,"Base":{"@type":"/gno.StructValue",`+objectInfo+`,"Fields":[{"T":`+intType+`,"N":"AQAAAAAAAAA="}]},"Index":"-1"}}]}`, nil)

	require.Len(t, res.Results, 4)
	assert.Equal(t, gnoKindPointer, res.Results[0].Kind)
	assert.Equal(t, "0dc5eff18974d66f6720bc387fc7e5c63514d961:7", res.Results[0].ObjectId)
	assert.Equal(t, uint64(1), res.Results[0].Index)

	assert.Equal(t, "*uint8", res.Results[1].Type)
	assert.Equal(t, gnoKindPointer, res.Results[1].Kind)
	assert.Empty(t, res.Results[1].ObjectId)
	require.Len(t, res.Results[1].Elems, 1)
	assertGnoScalar(t, res.Results[1].Elems[0], "uint8", gnoKindNumber, "105")

	// A special index can't be used to get the value which is pointed to, so it is not field or element 0.
	for _, v := range res.Results[2:] {
		assert.Equal(t, "*int", v.Type)
		assert.Equal(t, gnoKindUnknown, v.Kind)
		assert.Empty(t, v.ObjectId)
		assert.Empty(t, v.Elems)
	}
}

// testConvertQObjectJSON converts the object in data for the ref, using testQTypeJSON for the type queries.
func testConvertQObjectJSON(t *testing.T, value string, ref *api_gen.GnoValue) (*api_gen.GnoValue, error) {
	t.Helper()
	converter := newGnoValueConverter(testQuery(t, nil), zap.NewNop())
	return converter.convertQObjectJSON([]byte(`{"objectid":"`+ref.ObjectId+`","value":`+value+`}`), ref)
}

func TestConvertQObjectJSONStruct(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","ModTime":"0","RefCount":"1","LastObjectSize":"0"}`
	ref := &api_gen.GnoValue{
		Name:     "board",
		Type:     "gno.land/p/gnoland/boards/exts/hub/v0.Board",
		Kind:     gnoKindRef,
		ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:7",
		TypeId:   "gno.land/p/gnoland/boards/exts/hub/v0.Board",
	}
	value, err := testConvertQObjectJSON(t, `{"@type":"/gno.StructValue",`+objectInfo+`,"Fields":[`+
		`{"T":{"@type":"/gno.PrimitiveType","value":"65536"},"N":"AQAAAAAAAAA="},`+
		`{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"openboard1"}}]}`, ref)
	require.NoError(t, err)

	// The result replaces the ref, so it has the same name and type.
	assert.Equal(t, "board", value.Name)
	assert.Equal(t, "gno.land/p/gnoland/boards/exts/hub/v0.Board", value.Type)
	assert.Equal(t, gnoKindStruct, value.Kind)
	require.Len(t, value.Fields, 2)
	assert.Equal(t, "id", value.Fields[0].Name)
	assertGnoScalar(t, value.Fields[0], "uint64", gnoKindNumber, "1")
	assert.Equal(t, "name", value.Fields[1].Name)
	assertGnoScalar(t, value.Fields[1], "string", gnoKindString, "openboard1")

	// Without the type ID, the fields don't have names.
	ref.TypeId = ""
	value, err = testConvertQObjectJSON(t, `{"@type":"/gno.StructValue",`+objectInfo+`,"Fields":[`+
		`{"T":{"@type":"/gno.PrimitiveType","value":"65536"},"N":"AQAAAAAAAAA="}]}`, ref)
	require.NoError(t, err)
	assert.Equal(t, gnoKindStruct, value.Kind)
	require.Len(t, value.Fields, 1)
	assert.Empty(t, value.Fields[0].Name)
	assertGnoScalar(t, value.Fields[0], "uint64", gnoKindNumber, "1")
}

func TestConvertQObjectJSONSlice(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","ModTime":"0","RefCount":"1","LastObjectSize":"0"}`
	const stringType = `{"@type":"/gno.PrimitiveType","value":"16"}`
	const array = `{"@type":"/gno.ArrayValue",` + objectInfo + `,"List":[` +
		`{"T":` + stringType + `,"V":{"@type":"/gno.StringValue","value":"a"}},` +
		`{"T":` + stringType + `,"V":{"@type":"/gno.StringValue","value":"b"}},` +
		`{"T":` + stringType + `,"V":{"@type":"/gno.StringValue","value":"c"}}],"Data":null}`

	// Only the elements of the slice
	value, err := testConvertQObjectJSON(t, array, &api_gen.GnoValue{
		Type: "[]string", Kind: gnoKindRef, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:7", Offset: 1, Length: 2,
	})
	require.NoError(t, err)
	assert.Equal(t, "[]string", value.Type)
	assert.Equal(t, gnoKindList, value.Kind)
	require.Len(t, value.Elems, 2)
	assertGnoScalar(t, value.Elems[0], "string", gnoKindString, "b")
	assertGnoScalar(t, value.Elems[1], "string", gnoKindString, "c")

	// An array which is not a slice has all the elements.
	value, err = testConvertQObjectJSON(t, array, &api_gen.GnoValue{
		Type: "[3]string", Kind: gnoKindRef, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:7",
	})
	require.NoError(t, err)
	assert.Equal(t, "[3]string", value.Type)
	assert.Equal(t, gnoKindList, value.Kind)
	assert.Len(t, value.Elems, 3)

	// A slice of bytes
	value, err = testConvertQObjectJSON(t, `{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":null,"Data":"aGVsbG8="}`, &api_gen.GnoValue{
		Type: "[]uint8", Kind: gnoKindRef, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:7", Offset: 1, Length: 3,
	})
	require.NoError(t, err)
	assert.Equal(t, gnoKindBytes, value.Kind)
	assert.Equal(t, []byte("ell"), value.Data)
}

func TestConvertQObjectJSONPointer(t *testing.T) {
	const objectInfo = `"ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:7","ModTime":"0","RefCount":"1","LastObjectSize":"0"}`
	const intType = `{"@type":"/gno.PrimitiveType","value":"32"}`
	ref := &api_gen.GnoValue{Name: "count", Type: "*int", Kind: gnoKindPointer, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:7"}

	// A pointer to a heap item. The result replaces the pointer, so it has the same name.
	value, err := testConvertQObjectJSON(t, `{"@type":"/gno.HeapItemValue",`+objectInfo+`,"Value":{"T":`+intType+`,"N":"BwAAAAAAAAA="}}`, ref)
	require.NoError(t, err)
	assert.Equal(t, "count", value.Name)
	assertGnoScalar(t, value, "int", gnoKindNumber, "7")

	// A pointer to a field of a struct
	const structValue = `{"@type":"/gno.StructValue",` + objectInfo + `,"Fields":[` +
		`{"T":` + intType + `,"N":"AQAAAAAAAAA="},{"T":` + intType + `,"N":"AgAAAAAAAAA="}]}`
	ref.Index = 1
	value, err = testConvertQObjectJSON(t, structValue, ref)
	require.NoError(t, err)
	assert.Equal(t, "count", value.Name)
	assertGnoScalar(t, value, "int", gnoKindNumber, "2")

	// A pointer to an element of a byte array
	ref.Type = "*uint8"
	value, err = testConvertQObjectJSON(t, `{"@type":"/gno.ArrayValue",`+objectInfo+`,"List":null,"Data":"aGk="}`, ref)
	require.NoError(t, err)
	assertGnoScalar(t, value, "uint8", gnoKindNumber, "105")

	// The index is out of range.
	ref.Index = 2
	_, err = testConvertQObjectJSON(t, structValue, ref)
	assert.Error(t, err)
}

func TestConvertQObjectJSONPointerToBlock(t *testing.T) {
	// A package block, like the one of gno.land/r/sys/users. A pointer to a variable in the block which is not a
	// heap item has the block as its base.
	const block = `{"@type":"/gno.Block","ObjectInfo":{"ID":"0ed754d80e40da2645666583e74d4ab093970a6a:2","Hash":"b61a008de48cadda4e58a55465f03947fa4f2b52","ModTime":"541","RefCount":"8","IsEscaped":true,"LastObjectSize":"10607"},"Source":{"@type":"/gno.RefNode","Location":{"PkgPath":"gno.land/r/sys/users","File":"","Span":{"Pos":{"Line":"0","Column":"0"},"End":{"Line":"0","Column":"0"},"Num":"0"}},"BlockNode":null},"Values":[` +
		`{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"hello"}},` +
		`{"T":{"@type":"/gno.RefType","ID":"gno.land/r/sys/users.UserData"},"V":{"@type":"/gno.RefValue","ObjectID":"0ed754d80e40da2645666583e74d4ab093970a6a:535","Hash":"c71b7fdf1ea86d0d240dbee17e52c795e78ee716"}}],"Parent":null,"Blank":{}}`
	ref := &api_gen.GnoValue{Type: "*string", Kind: gnoKindPointer, ObjectId: "0ed754d80e40da2645666583e74d4ab093970a6a:2"}

	// A string variable
	value, err := testConvertQObjectJSON(t, block, ref)
	require.NoError(t, err)
	assertGnoScalar(t, value, "string", gnoKindString, "hello")

	// A struct variable, where the struct is persisted separately, so the ref is followed.
	ref.Type = "*gno.land/r/sys/users.UserData"
	ref.TypeId = "gno.land/r/sys/users.UserData"
	ref.Index = 1
	value, err = testConvertQObjectJSON(t, block, ref)
	require.NoError(t, err)
	assert.Equal(t, "gno.land/r/sys/users.UserData", value.Type)
	assert.Equal(t, gnoKindStruct, value.Kind)
	require.Len(t, value.Fields, 3)
	assert.Equal(t, "username", value.Fields[1].Name)
	assertGnoScalar(t, value.Fields[1], "string", gnoKindString, "test1")

	// The index is out of range.
	ref.Index = 2
	_, err = testConvertQObjectJSON(t, block, ref)
	assert.Error(t, err)
}

func TestConvertQObjectJSONFollowRef(t *testing.T) {
	// The heap item which a *UserData points to. The UserData is persisted separately, so the heap item only has
	// a ref to it, which is followed.
	var queried []string
	converter := newGnoValueConverter(testQuery(t, &queried), zap.NewNop())
	value, err := converter.convertQObjectJSON([]byte(`{"objectid":"0ed754d80e40da2645666583e74d4ab093970a6a:534","value":{"@type":"/gno.HeapItemValue","ObjectInfo":{"ID":"0ed754d80e40da2645666583e74d4ab093970a6a:534","Hash":"65c78d32cc4893fc17f9addab1c7e77e0badc133","ModTime":"0","RefCount":"2","IsEscaped":true,"LastObjectSize":"351"},"Value":{"T":{"@type":"/gno.RefType","ID":"gno.land/r/sys/users.UserData"},"V":{"@type":"/gno.RefValue","ObjectID":"0ed754d80e40da2645666583e74d4ab093970a6a:535","Hash":"c71b7fdf1ea86d0d240dbee17e52c795e78ee716"}}}}`),
		&api_gen.GnoValue{
			Name: "user", Type: "*gno.land/r/sys/users.UserData", Kind: gnoKindPointer,
			ObjectId: "0ed754d80e40da2645666583e74d4ab093970a6a:534", TypeId: "gno.land/r/sys/users.UserData",
		})
	require.NoError(t, err)

	// The name of the pointer is kept after following the ref.
	assert.Equal(t, "user", value.Name)
	assert.Equal(t, "gno.land/r/sys/users.UserData", value.Type)
	assert.Equal(t, gnoKindStruct, value.Kind)
	require.Len(t, value.Fields, 3)
	assert.Equal(t, "addr", value.Fields[0].Name)
	assertGnoScalar(t, value.Fields[0], "address", gnoKindString, "g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5")
	assert.Equal(t, "username", value.Fields[1].Name)
	assertGnoScalar(t, value.Fields[1], "string", gnoKindString, "test1")
	assert.Equal(t, "deleted", value.Fields[2].Name)
	assertGnoScalar(t, value.Fields[2], "bool", gnoKindBool, "false")
	assert.Equal(t, []string{"0ed754d80e40da2645666583e74d4ab093970a6a:535", "gno.land/r/sys/users.UserData"}, queried)

	// A heap item which has a slice of a persisted array. The ref is followed, and only has the elements of the slice.
	const stringType = `{"@type":"/gno.PrimitiveType","value":"16"}`
	value, err = converter.convertQObjectJSON([]byte(`{"objectid":"0dc5eff18974d66f6720bc387fc7e5c63514d961:20","value":{"@type":"/gno.HeapItemValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:20","ModTime":"0","RefCount":"1","LastObjectSize":"0"},"Value":{"T":{"@type":"/gno.SliceType","Elt":`+stringType+`,"Vrd":false},"V":{"@type":"/gno.SliceValue","Base":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:21","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"},"Offset":"1","Length":"2","Maxcap":"2"}}}}`),
		&api_gen.GnoValue{Type: "*[]string", Kind: gnoKindPointer, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:20"})
	require.NoError(t, err)
	assert.Equal(t, "[]string", value.Type)
	assert.Equal(t, gnoKindList, value.Kind)
	require.Len(t, value.Elems, 2)
	assertGnoScalar(t, value.Elems[0], "string", gnoKindString, "b")
	assertGnoScalar(t, value.Elems[1], "string", gnoKindString, "c")
}

func TestConvertQObjectJSONFollowOneRef(t *testing.T) {
	// The heap item has a ref to another heap item which has a ref. Only the first ref is followed.
	var queried []string
	converter := newGnoValueConverter(testQuery(t, &queried), zap.NewNop())
	value, err := converter.convertQObjectJSON([]byte(`{"objectid":"0dc5eff18974d66f6720bc387fc7e5c63514d961:22","value":{"@type":"/gno.HeapItemValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:22","ModTime":"0","RefCount":"1","LastObjectSize":"0"},"Value":{"T":{"@type":"/gno.RefType","ID":"gno.land/r/foo.Node"},"V":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:23","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"}}}}`),
		&api_gen.GnoValue{Type: "*gno.land/r/foo.Node", Kind: gnoKindPointer, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:22"})
	require.NoError(t, err)

	assert.Equal(t, gnoKindRef, value.Kind)
	assert.Equal(t, "0dc5eff18974d66f6720bc387fc7e5c63514d961:24", value.ObjectId)
	assert.Equal(t, []string{"0dc5eff18974d66f6720bc387fc7e5c63514d961:23"}, queried)

	// The object of the followed ref is not found. The error has the object ID, and wraps the error of the query.
	_, err = converter.convertQObjectJSON([]byte(`{"objectid":"0dc5eff18974d66f6720bc387fc7e5c63514d961:22","value":{"@type":"/gno.HeapItemValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:22","ModTime":"0","RefCount":"1","LastObjectSize":"0"},"Value":{"T":{"@type":"/gno.RefType","ID":"gno.land/r/foo.Node"},"V":{"@type":"/gno.RefValue","ObjectID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:99","Hash":"64ed6ee11c3332f9172a134ca50f4a3b03b3a829"}}}}`),
		&api_gen.GnoValue{Type: "*gno.land/r/foo.Node", Kind: gnoKindPointer, ObjectId: "0dc5eff18974d66f6720bc387fc7e5c63514d961:22"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "can't follow the ref to object 0dc5eff18974d66f6720bc387fc7e5c63514d961:99")
	assert.EqualError(t, errors.Unwrap(err), "not found: 0dc5eff18974d66f6720bc387fc7e5c63514d961:99")
}

func TestConvertQObjectJSONMalformed(t *testing.T) {
	converter := newGnoValueConverter(func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error) {
		return nil, errors.New("unexpected")
	}, zap.NewNop())
	_, err := converter.convertQObjectJSON([]byte(`not JSON`), &api_gen.GnoValue{Kind: gnoKindRef, ObjectId: "x:1"})
	assert.Error(t, err)
}
