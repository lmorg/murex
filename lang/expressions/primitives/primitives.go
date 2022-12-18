package primitives

import (
	"encoding/json"

	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

//go:generate stringer -type=Primitive
type Primitive int

const (
	Number   Primitive = Primitive(symbols.Number)
	String   Primitive = Primitive(symbols.QuoteSingle)
	Boolean  Primitive = Primitive(symbols.Boolean)
	Array    Primitive = Primitive(symbols.ArrayBegin)
	Object   Primitive = Primitive(symbols.ObjectBegin)
	Null     Primitive = Primitive(symbols.Null)
	Bareword Primitive = 0
	Other    Primitive = -1
)

type DataType struct {
	Primitive Primitive
	Value     interface{}
	MxDT      string
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
