package primitives

import (
	"encoding/json"

	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

//go:generate stringer -linecomment -type=Primitive
type Primitive int

const (
	Number   Primitive = Primitive(symbols.Number)      // number
	String   Primitive = Primitive(symbols.QuoteSingle) // string
	Boolean  Primitive = Primitive(symbols.Boolean)     // boolean
	Array    Primitive = Primitive(symbols.ArrayBegin)  // array
	Object   Primitive = Primitive(symbols.ObjectBegin) // object
	Null     Primitive = Primitive(symbols.Null)        // null
	Bareword Primitive = 0                              // bareword
	Other    Primitive = -1                             // other
)

type DataType struct {
	Primitive Primitive
	value     any
	subshell  []rune
	MxDT      string
}

func NewPrimitive(primitive Primitive, value any) *DataType {
	return &DataType{
		Primitive: primitive,
		value:     value,
	}
}

func NewFunction(primitive Primitive, block []rune) *DataType {
	return &DataType{
		Primitive: primitive,
		subshell:  block,
	}
}

func Scalar2Primitive(dt string, value any) *DataType {
	switch dt {
	case types.Number, types.Integer, types.Float:
		return &DataType{Primitive: Number, value: value}
	case types.Boolean:
		return &DataType{Primitive: Boolean, value: value}
	case types.Null:
		return &DataType{Primitive: Null, value: value}
	case types.String:
		return &DataType{Primitive: String, value: value}
	default:
		return &DataType{Primitive: Other, MxDT: dt, value: value}
	}
}

func (dt *DataType) DataType() string {
	switch dt.Primitive {
	case Number:
		return types.Number
	case String:
		return types.String
	case Boolean:
		return types.Boolean
	case Array:
		return types.Json
	case Object:
		return types.Json
	case Null:
		return types.Null
	case Bareword:
		return types.Null
	case Other:
		return dt.MxDT
	default:
		return types.Generic
	}
}

func (dt *DataType) Marshal() ([]rune, error) {
	b, err := json.Marshal(dt.Value)
	if err != nil {
		return nil, err
	}

	r := []rune(string(b))
	return r, nil
}

func (dt *DataType) NotValue() {
	dt.value = !dt.value.(bool)
}

func (dt *DataType) Value() any {
	if dt.subshell == nil {
		return dt.value
	}
	return nil //TODO
}
