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
	Value     any
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
