package primitives

import (
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

//go:generate stringer -type=Primitive
type Primitive int

const (
	Number  Primitive = Primitive(symbols.Number)
	String  Primitive = Primitive(symbols.QuoteSingle)
	Boolean Primitive = Primitive(symbols.Boolean)
	Array   Primitive = Primitive(symbols.ArrayBegin)
	Object  Primitive = Primitive(symbols.ObjectBegin)
	Null    Primitive = Primitive(symbols.Null)
)

type DataType struct {
	Primitive Primitive
	Value     interface{}
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
	default:
		return types.Generic
	}
}
