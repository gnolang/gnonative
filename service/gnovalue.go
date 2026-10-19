// This file converts the Amino JSON result of the "vm/qeval_json" query to the GnoValue of the gRPC API.

package service

import (
	"fmt"
	"strings"

	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	"github.com/gnolang/gno/tm2/pkg/amino"
	"github.com/gnolang/gnonative/v4/service/utils"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// The kinds of api_gen.GnoValue
const (
	gnoKindBool    = "bool"
	gnoKindNumber  = "number"
	gnoKindString  = "string"
	gnoKindBytes   = "bytes"
	gnoKindStruct  = "struct"
	gnoKindList    = "list"
	gnoKindMap     = "map"
	gnoKindPointer = "pointer"
	gnoKindFunc    = "func"
	gnoKindType    = "type"
	gnoKindRef     = "ref"
	gnoKindNil     = "nil"
	gnoKindUnknown = "unknown"
)

// gnoValueConverter converts a gno.TypedValue to an api_gen.GnoValue. A value only has the ID of its
// declared type, so use queryType to get the type definition, which has the struct field names.
type gnoValueConverter struct {
	// queryType does the "vm/qtype_json" query for the type ID and returns the response data.
	queryType func(typeID string) ([]byte, error)
	// types caches the result of queryType by type ID. The value is nil if the type can't be resolved.
	types map[string]gno.Type
}

func newGnoValueConverter(queryType func(typeID string) ([]byte, error)) *gnoValueConverter {
	return &gnoValueConverter{queryType: queryType, types: map[string]gno.Type{}}
}

// convertQEvalJSON decodes the response data of the "vm/qeval_json" query and converts each result.
func (c *gnoValueConverter) convertQEvalJSON(data []byte) (res *api_gen.QEvalJSONValuesResponse, err error) {
	// The data is from the remote node. Don't let a malformed value panic the service.
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = fmt.Errorf("can't convert the qeval_json result: %v", r)
		}
	}()

	var qres utils.EvalResponse
	if err := amino.UnmarshalJSON(data, &qres); err != nil {
		return nil, fmt.Errorf("can't decode the qeval_json result: %w", err)
	}

	res = &api_gen.QEvalJSONValuesResponse{Results: make([]*api_gen.GnoValue, len(qres.Results))}
	for i, tv := range qres.Results {
		res.Results[i] = c.convert(tv)
	}
	if qres.Error != nil {
		res.Error = *qres.Error
	}
	return res, nil
}

// resolveType returns the type definition for the type ID, or nil if it can't be resolved.
// The result (including nil) is cached in c.types, so each type ID is queried at most once.
// The cache only lasts as long as the converter, which is made for each request, so it doesn't
// need to be invalidated. (Between requests, a realm can be redeployed, or SetRemote can change
// the node, so the definition for a type ID may change. And a failure may be a network error
// which should be retried in the next request.)
func (c *gnoValueConverter) resolveType(typeID string) (t gno.Type) {
	if t, ok := c.types[typeID]; ok {
		return t
	}
	defer func() {
		if r := recover(); r != nil {
			t = nil
		}
		c.types[typeID] = t
	}()

	data, err := c.queryType(typeID)
	if err != nil {
		return nil
	}
	var qres utils.TypeResponse
	if err := amino.UnmarshalJSON(data, &qres); err != nil {
		return nil
	}
	return qres.Type
}

// baseType resolves a type ID and unwraps a declared type to return its underlying type, or nil if
// it can't be resolved.
func (c *gnoValueConverter) baseType(t gno.Type) gno.Type {
	// A declared type whose base is another reference is not expected, but don't loop forever.
	for range 8 {
		switch ct := t.(type) {
		case gno.RefType:
			t = c.resolveType(string(ct.ID))
		case *gno.DeclaredType:
			t = gno.BaseOf(ct)
		default:
			return t
		}
	}
	return nil
}

// gnoTypeString returns the type as in Gno code, where a declared type has its package path.
func gnoTypeString(t gno.Type) string {
	switch ct := t.(type) {
	case nil:
		return ""
	case gno.PrimitiveType:
		return utils.PrimitiveTypeName(ct)
	case gno.RefType:
		// Example: ".uverse.address"
		return strings.TrimPrefix(string(ct.ID), ".uverse.")
	case *gno.DeclaredType:
		if ct.PkgPath == ".uverse" {
			return string(ct.Name)
		}
		return ct.PkgPath + "." + string(ct.Name)
	case *gno.PointerType:
		return "*" + gnoTypeString(ct.Elt)
	case *gno.SliceType:
		return "[]" + gnoTypeString(ct.Elt)
	case *gno.ArrayType:
		return fmt.Sprintf("[%d]%s", ct.Len, gnoTypeString(ct.Elt))
	case *gno.MapType:
		return "map[" + gnoTypeString(ct.Key) + "]" + gnoTypeString(ct.Value)
	default:
		return t.String()
	}
}

func (c *gnoValueConverter) convert(tv gno.TypedValue) *api_gen.GnoValue {
	res := &api_gen.GnoValue{Type: gnoTypeString(tv.T)}

	switch cv := tv.V.(type) {
	case nil:
		c.convertNoValue(res, tv)
	case gno.StringValue:
		res.Kind = gnoKindString
		res.Scalar = string(cv)
	case gno.BigintValue:
		res.Kind = gnoKindNumber
		res.Scalar, _ = cv.MarshalAmino()
	case gno.BigdecValue:
		res.Kind = gnoKindNumber
		res.Scalar, _ = cv.MarshalAmino()
	case *gno.StructValue:
		res.Kind = gnoKindStruct
		st, _ := c.baseType(tv.T).(*gno.StructType)
		res.Fields = make([]*api_gen.GnoValue, len(cv.Fields))
		for i, field := range cv.Fields {
			res.Fields[i] = c.convert(field)
			if st != nil && i < len(st.Fields) {
				res.Fields[i].Name = string(st.Fields[i].Name)
			}
		}
	case *gno.ArrayValue:
		c.convertArray(res, cv, 0, cv.GetLength())
	case *gno.SliceValue:
		switch base := cv.Base.(type) {
		case *gno.ArrayValue:
			c.convertArray(res, base, cv.Offset, cv.Length)
		case gno.RefValue:
			res.Kind = gnoKindRef
			res.ObjectId = base.ObjectID.String()
			res.Offset = uint64(max(cv.Offset, 0))
			res.Length = uint64(max(cv.Length, 0))
		case gno.ExportRefValue:
			res.Kind = gnoKindRef
			res.ObjectId = base.ObjectID
			res.Offset = uint64(max(cv.Offset, 0))
			res.Length = uint64(max(cv.Length, 0))
		default:
			res.Kind = gnoKindList
		}
	case *gno.MapValue:
		res.Kind = gnoKindMap
		if cv.List != nil {
			for item := cv.List.Head; item != nil; item = item.Next {
				res.Keys = append(res.Keys, c.convert(item.Key))
				res.Elems = append(res.Elems, c.convert(item.Value))
			}
		}
	case gno.PointerValue:
		res.Kind = gnoKindPointer
		switch base := cv.Base.(type) {
		case *gno.HeapItemValue:
			res.Elems = []*api_gen.GnoValue{c.convert(base.Value)}
		case *gno.ArrayValue:
			if cv.Index >= 0 && cv.Index < len(base.List) {
				res.Elems = []*api_gen.GnoValue{c.convert(base.List[cv.Index])}
			}
		case *gno.StructValue:
			if cv.Index >= 0 && cv.Index < len(base.Fields) {
				res.Elems = []*api_gen.GnoValue{c.convert(base.Fields[cv.Index])}
			}
		case gno.RefValue:
			res.ObjectId = base.ObjectID.String()
			res.TypeId = utils.GetTypeIDQueryString(tv.T)
		case gno.ExportRefValue:
			res.ObjectId = base.ObjectID
			res.TypeId = utils.GetTypeIDQueryString(tv.T)
		}
	case *gno.FuncValue:
		res.Kind = gnoKindFunc
		res.Scalar = string(cv.Name)
	case *gno.BoundMethodValue:
		res.Kind = gnoKindFunc
		if cv.Func != nil {
			res.Scalar = string(cv.Func.Name)
		}
	case gno.TypeValue:
		res.Kind = gnoKindType
		res.Scalar = gnoTypeString(cv.Type)
	case gno.RefValue:
		res.Kind = gnoKindRef
		res.ObjectId = cv.ObjectID.String()
		res.TypeId = utils.GetTypeIDQueryString(tv.T)
	case gno.ExportRefValue:
		res.Kind = gnoKindRef
		res.ObjectId = cv.ObjectID
		res.TypeId = utils.GetTypeIDQueryString(tv.T)
	default:
		res.Kind = gnoKindUnknown
	}

	return res
}

// convertArray sets res to the length elements of the array starting at offset.
func (c *gnoValueConverter) convertArray(res *api_gen.GnoValue, av *gno.ArrayValue, offset int, length int) {
	if av.Data != nil {
		res.Kind = gnoKindBytes
		offset, end := utils.ClampSliceWindow(offset, length, len(av.Data))
		res.Data = av.Data[offset:end]
		return
	}

	res.Kind = gnoKindList
	offset, end := utils.ClampSliceWindow(offset, length, len(av.List))
	res.Elems = make([]*api_gen.GnoValue, end-offset)
	for i, elem := range av.List[offset:end] {
		res.Elems[i] = c.convert(elem)
	}
}

// gnoListKind returns the kind of a slice or array with the element type, which is "bytes" for a byte.
func gnoListKind(elt gno.Type) string {
	if pt, ok := elt.(gno.PrimitiveType); ok && pt.Kind() == gno.Uint8Kind {
		return gnoKindBytes
	}
	return gnoKindList
}

// convertNoValue sets res for tv.V nil. This is a primitive value in tv.N, or else the nil value of the type.
func (c *gnoValueConverter) convertNoValue(res *api_gen.GnoValue, tv gno.TypedValue) {
	if tv.T == nil {
		res.Kind = gnoKindNil
		return
	}

	switch bt := c.baseType(tv.T).(type) {
	case nil:
		// The type can't be resolved, so we don't know how to interpret tv.N .
		res.Kind = gnoKindUnknown
		res.Data = tv.N[:]
	case gno.PrimitiveType:
		switch bt.Kind() {
		case gno.BoolKind:
			res.Kind = gnoKindBool
			res.Scalar = utils.PrimitiveString(bt, tv)
		case gno.StringKind:
			res.Kind = gnoKindString
		case gno.IntKind, gno.Int8Kind, gno.Int16Kind, gno.Int32Kind, gno.Int64Kind,
			gno.UintKind, gno.Uint8Kind, gno.Uint16Kind, gno.Uint32Kind, gno.Uint64Kind,
			gno.Float32Kind, gno.Float64Kind:
			res.Kind = gnoKindNumber
			res.Scalar = utils.PrimitiveString(bt, tv)
		case gno.BigintKind, gno.BigdecKind:
			res.Kind = gnoKindNumber
			res.Scalar = "0"
		default:
			res.Kind = gnoKindUnknown
			res.Data = tv.N[:]
		}
	case *gno.SliceType:
		res.Kind = gnoListKind(bt.Elt)
	case *gno.ArrayType:
		res.Kind = gnoListKind(bt.Elt)
	case *gno.MapType:
		res.Kind = gnoKindMap
	case *gno.StructType:
		res.Kind = gnoKindStruct
	default:
		// A nil pointer, func or interface.
		res.Kind = gnoKindNil
	}
}
