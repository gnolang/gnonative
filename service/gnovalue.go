// This file converts the Amino JSON result of the "vm/qeval_json" and "vm/qobject_json" queries to the
// GnoValue of the gRPC API.

package service

import (
	"fmt"
	"math"
	"strings"

	"github.com/gnolang/gno/gno.land/pkg/gnoclient"
	gno "github.com/gnolang/gno/gnovm/pkg/gnolang"
	"github.com/gnolang/gno/tm2/pkg/amino"
	ctypes "github.com/gnolang/gno/tm2/pkg/bft/rpc/core/types"
	"github.com/gnolang/gnonative/v4/service/utils"
	"go.uber.org/zap"

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
// declared type, so use query to get the type definition, which has the struct field names.
type gnoValueConverter struct {
	// query is the client's Query, which is used for the "vm/qtype_json" query of a type ID, and for the
	// "vm/qobject_json" query to follow a ref in convertQObjectJSON.
	query func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error)
	// types caches the result of the "vm/qtype_json" query by type ID. The value is nil if the type can't be resolved.
	types map[string]gno.Type
	// logger is used for a debug message when a type can't be resolved.
	logger *zap.Logger
}

func newGnoValueConverter(query func(cfg gnoclient.QueryCfg) (*ctypes.ResultABCIQuery, error), logger *zap.Logger) *gnoValueConverter {
	return &gnoValueConverter{query: query, types: map[string]gno.Type{}, logger: logger}
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

// convertQObjectJSON decodes the response data of the "vm/qobject_json" query for value.ObjectId and converts
// the object. value is the GnoValue of kind "ref" or "pointer" which has the object ID. If value is a slice, only
// convert the elements of the slice. If value is a pointer, convert the value which is pointed to. If the result
// is a ref to another persisted object (like a pointer to a struct which is persisted separately), query and
// convert that object. Only follow one ref like this, so that there is no loop.
func (c *gnoValueConverter) convertQObjectJSON(data []byte, value *api_gen.GnoValue) (res *api_gen.GnoValue, err error) {
	// The data is from the remote node. Don't let a malformed value panic the service.
	defer func() {
		if r := recover(); r != nil {
			res = nil
			err = fmt.Errorf("can't convert the qobject_json result: %v", r)
		}
	}()

	res, err = c.convertObject(data, value)
	if err != nil || res.Kind != gnoKindRef {
		return res, err
	}

	// TODO: Use gnoclient QObjectJSON
	qres, err := c.query(gnoclient.QueryCfg{Path: "vm/qobject_json", Data: []byte(res.ObjectId)})
	if err != nil {
		return nil, err
	}
	return c.convertObject(qres.Response.Data, res)
}

// convertObject decodes the response data of the "vm/qobject_json" query and converts the object for value
// as described in convertQObjectJSON, but doesn't follow a ref in the result.
func (c *gnoValueConverter) convertObject(data []byte, value *api_gen.GnoValue) (res *api_gen.GnoValue, err error) {
	var qres utils.ObjectResponse
	if err := amino.UnmarshalJSON(data, &qres); err != nil {
		return nil, fmt.Errorf("can't decode the qobject_json result: %w", err)
	}

	if value.Kind == gnoKindPointer {
		return c.convertPointedTo(qres.Value, value.Index)
	}

	if av, ok := qres.Value.(*gno.ArrayValue); ok && value.Length > 0 {
		// A slice of the persisted array
		res = &api_gen.GnoValue{}
		c.convertArray(res, av, int(min(value.Offset, math.MaxInt)), int(min(value.Length, math.MaxInt)))
	} else {
		// The object doesn't have its type, so use the type ID from the value (if any) to get the struct field names.
		var t gno.Type
		if value.TypeId != "" {
			t = gno.RefType{ID: gno.TypeID(value.TypeId)}
		}
		res = c.convert(gno.TypedValue{T: t, V: qres.Value})
	}
	// The result replaces the value, so it has the same name and type.
	res.Name = value.Name
	res.Type = value.Type
	return res, nil
}

// convertPointedTo converts the value in the persisted object which a pointer points to. If the object is a
// struct or array, index is the field or element.
func (c *gnoValueConverter) convertPointedTo(object gno.Value, index uint64) (*api_gen.GnoValue, error) {
	switch cv := object.(type) {
	case *gno.HeapItemValue:
		return c.convert(cv.Value), nil
	case *gno.StructValue:
		if index < uint64(len(cv.Fields)) {
			return c.convert(cv.Fields[index]), nil
		}
	case *gno.ArrayValue:
		if cv.Data != nil {
			if index < uint64(len(cv.Data)) {
				tv := gno.TypedValue{T: gno.Uint8Type}
				tv.SetUint8(cv.Data[index])
				return c.convert(tv), nil
			}
		} else if index < uint64(len(cv.List)) {
			return c.convert(cv.List[index]), nil
		}
	default:
		return &api_gen.GnoValue{Kind: gnoKindUnknown}, nil
	}
	return nil, fmt.Errorf("the pointer index %d is out of range of the qobject_json result", index)
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
			c.logger.Debug("resolveType: can't convert the qtype_json result", zap.String("typeId", typeID), zap.Any("panic", r))
			t = nil
		}
		c.types[typeID] = t
	}()

	// TODO: Use gnoclient QTypeJSON
	qres, err := c.query(gnoclient.QueryCfg{Path: "vm/qtype_json", Data: []byte(typeID)})
	if err != nil {
		// This may be a network error, which is cached the same as a type which doesn't exist, so log it.
		c.logger.Debug("resolveType: qtype_json query error", zap.String("typeId", typeID), zap.Error(err))
		return nil
	}
	var tres utils.TypeResponse
	if err := amino.UnmarshalJSON(qres.Response.Data, &tres); err != nil {
		c.logger.Debug("resolveType: can't decode the qtype_json result", zap.String("typeId", typeID), zap.Error(err))
		return nil
	}
	return tres.Type
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
			c.convertSliceRef(res, tv.T, base.ObjectID.String(), cv.Offset, cv.Length)
		case gno.ExportRefValue:
			c.convertSliceRef(res, tv.T, base.ObjectID, cv.Offset, cv.Length)
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
		if cv.Index < 0 {
			// A special index like PointerIndexMap, where the value which is pointed to is not a field or
			// element of the base, so we can't get it.
			res.Kind = gnoKindUnknown
			break
		}
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
			res.Index = uint64(cv.Index)
		case gno.ExportRefValue:
			res.ObjectId = base.ObjectID
			res.TypeId = utils.GetTypeIDQueryString(tv.T)
			res.Index = uint64(cv.Index)
		}
	case *gno.FuncValue:
		res.Kind = gnoKindFunc
		res.Scalar = string(cv.Name)
	case *gno.BoundMethodValue:
		res.Kind = gnoKindFunc
		if cv.Func != nil {
			res.Scalar = string(cv.Func.Name)
		}
	case *gno.HeapItemValue:
		// A persisted object which holds one value, like the target of a pointer. The value has its own type.
		return c.convert(cv.Value)
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

// convertSliceRef sets res to the slice with type t of the persisted array with the object ID. An empty slice
// is an empty list instead of a ref, so that a ref of a slice always has a length which is not 0.
func (c *gnoValueConverter) convertSliceRef(res *api_gen.GnoValue, t gno.Type, objectID string, offset int, length int) {
	if length <= 0 {
		res.Kind = gnoKindList
		if st, ok := c.baseType(t).(*gno.SliceType); ok {
			res.Kind = gnoListKind(st.Elt)
		}
		return
	}

	res.Kind = gnoKindRef
	res.ObjectId = objectID
	res.Offset = uint64(max(offset, 0))
	res.Length = uint64(length)
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
