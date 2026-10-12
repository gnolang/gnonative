package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// The response of the "vm/qeval_json" query for "gno.land/r/gnoland/boards2/v0.GetBoard(1)".
const testQEvalJSONBoard = `{"results":[{"T":{"@type":"/gno.RefType","ID":"gno.land/p/gnoland/boards/exts/hub/v0.Board"},"V":{"@type":"/gno.StructValue","ObjectInfo":{"ID":"0dc5eff18974d66f6720bc387fc7e5c63514d961:0","ModTime":"0","RefCount":"0","LastObjectSize":"0"},"Fields":[{"T":{"@type":"/gno.PrimitiveType","value":"65536"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"16"},"V":{"@type":"/gno.StringValue","value":"openboard1"}},{"T":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"},"Vrd":false}},{"T":{"@type":"/gno.PrimitiveType","value":"4"}},{"T":{"@type":"/gno.PrimitiveType","value":"32"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"32"},"N":"AQAAAAAAAAA="},{"T":{"@type":"/gno.RefType","ID":".uverse.address"},"V":{"@type":"/gno.StringValue","value":"g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5"}},{"T":{"@type":"/gno.PrimitiveType","value":"1024"},"N":"5v2zagAAAAA="},{"T":{"@type":"/gno.PrimitiveType","value":"1024"}}]}},{"T":{"@type":"/gno.PrimitiveType","value":"4"},"N":"AQAAAAAAAAA="}]}`

// The responses of the "vm/qtype_json" query by type ID.
var testQTypeJSON = map[string]string{
	"gno.land/p/gnoland/boards/exts/hub/v0.Board": `{"typeid":"gno.land/p/gnoland/boards/exts/hub/v0.Board","type":{"@type":"/gno.DeclaredType","PkgPath":"gno.land/p/gnoland/boards/exts/hub/v0","Name":"Board","Base":{"@type":"/gno.StructType","Fields":[{"Name":"id","Type":{"@type":"/gno.PrimitiveType","value":"65536"}},{"Name":"name","Type":{"@type":"/gno.PrimitiveType","value":"16"}},{"Name":"aliases","Type":{"@type":"/gno.SliceType","Elt":{"@type":"/gno.PrimitiveType","value":"16"}}},{"Name":"readonly","Type":{"@type":"/gno.PrimitiveType","value":"4"}},{"Name":"threadCount","Type":{"@type":"/gno.PrimitiveType","value":"32"}},{"Name":"memberCount","Type":{"@type":"/gno.PrimitiveType","value":"32"}},{"Name":"creator","Type":{"@type":"/gno.DeclaredType","PkgPath":".uverse","Name":"address","Base":{"@type":"/gno.PrimitiveType","value":"16"}}},{"Name":"createdAt","Type":{"@type":"/gno.PrimitiveType","value":"1024"}},{"Name":"updatedAt","Type":{"@type":"/gno.PrimitiveType","value":"1024"}}]}}}`,
	".uverse.address":                `{"typeid":".uverse.address","type":{"@type":"/gno.DeclaredType","PkgPath":".uverse","Name":"address","Base":{"@type":"/gno.PrimitiveType","value":"16"}}}`,
	"gno.land/r/demo/boards.BoardID": `{"typeid":"gno.land/r/demo/boards.BoardID","type":{"@type":"/gno.DeclaredType","PkgPath":"gno.land/r/demo/boards","Name":"BoardID","Base":{"@type":"/gno.PrimitiveType","value":"65536"}}}`,
}

// testConvertQEvalJSON converts data, using testQTypeJSON for the type queries. Add each queried type ID to queried.
func testConvertQEvalJSON(t *testing.T, data string, queried *[]string) *api_gen.QEvalJSONValuesResponse {
	t.Helper()
	converter := newGnoValueConverter(func(typeID string) ([]byte, error) {
		if queried != nil {
			*queried = append(*queried, typeID)
		}
		res, ok := testQTypeJSON[typeID]
		if !ok {
			return nil, errors.New("type not found: " + typeID)
		}
		return []byte(res), nil
	})
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

func TestConvertQEvalJSONMalformed(t *testing.T) {
	converter := newGnoValueConverter(func(typeID string) ([]byte, error) {
		return nil, errors.New("unexpected")
	})

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
