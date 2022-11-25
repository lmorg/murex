package primitives

import "github.com/lmorg/murex/builtins/core/expressions/symbols"

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
	DataType  string
	Value     interface{}
}
