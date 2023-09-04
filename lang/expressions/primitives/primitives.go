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
	Function Primitive = -2                             // functions and subshells
)

type DataType struct {
	primitive Primitive
	value     any
	mxDT      string
	fn        FunctionT
}

type FunctionT func() (any, string, error)

func NewPrimitive(primitive Primitive, value any) *DataType {
	return &DataType{
		primitive: primitive,
		value:     value,
	}
}

func NewFunction(fn FunctionT) *DataType {
	return &DataType{
		primitive: Function,
		fn:        fn,
	}
}

func Scalar2Primitive(mxdt string, value any) *DataType {
	return &DataType{
		primitive: scalar2Primitive(mxdt),
		mxDT:      mxdt,
		value:     value,
	}
}

func scalar2Primitive(dt string) Primitive {
	switch dt {
	case types.Number, types.Integer, types.Float:
		return Number
	case types.Boolean:
		return Boolean
	case types.Null:
		return Null
	case types.String:
		return String
	default:
		return Other
	}
}

func (dt *DataType) dataType() string {
	switch dt.primitive {
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
		return dt.mxDT
	default:
		return types.Generic
	}
}

func (v *Value) Marshal() ([]rune, error) {
	b, err := json.Marshal(v.Value)
	if err != nil {
		return nil, err
	}

	r := []rune(string(b))
	return r, nil
}

func (dt *DataType) NotValue() {
	dt.value = !dt.value.(bool)
}

type Value struct {
	Primitive Primitive
	Value     any
	DataType  string
}

func (dt *DataType) GetValue() (*Value, error) {
	if dt.primitive == Function {
		if dt.fn == nil {
			panic("undefined function")
		}
		var err error
		dt.value, dt.mxDT, err = dt.fn()
		dt.primitive = scalar2Primitive(dt.mxDT)
		return &Value{dt.primitive, dt.value, dt.mxDT}, err
	}

	if dt.mxDT == "" {
		dt.mxDT = dt.dataType()
	}

	return &Value{dt.primitive, dt.value, dt.mxDT}, nil
}
