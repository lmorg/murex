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
	Function Primitive = -2                             // functions and sub-shells
)

func (primitive Primitive) IsComparable() bool {
	return primitive == Number || primitive == String || primitive == Boolean || primitive == Null
}

type DataType struct {
	v  *Value
	fn FunctionT
}

type FunctionT func() (*Value, error)

func NewPrimitive(primitive Primitive, value any) *DataType {
	return &DataType{
		v: &Value{
			Primitive: primitive,
			Value:     value,
		},
	}
}

func NewFunction(fn FunctionT) *DataType {
	return &DataType{
		v:  &Value{Primitive: Function},
		fn: fn,
	}
}

func NewScalar(mxdt string, value any) *DataType {
	return &DataType{
		v: &Value{
			Primitive: DataType2Primitive(mxdt),
			DataType:  mxdt,
			Value:     value,
		},
	}
}

func DataType2Primitive(dt string) Primitive {
	switch dt {
	case types.Number, types.Integer, types.Float:
		return Number
	case types.Boolean:
		return Boolean
	case types.Null:
		return Null
	case types.String, types.Generic:
		return String
	default:
		return Other
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
	dt.v.Value = !dt.v.Value.(bool)
}

type Value struct {
	Primitive Primitive
	Value     any
	DataType  string
	ExitNum   int
}

func (dt *DataType) GetValue() (*Value, error) {
	if dt.v.Primitive == Function {
		if dt.fn == nil {
			panic("undefined function")
		}

		var err error
		dt.v, err = dt.fn()
		dt.v.Primitive = DataType2Primitive(dt.v.DataType)
		return dt.v, err
	}

	if dt.v.DataType == "" {
		dt.v.DataType = dt.dataType()
	}

	return dt.v, nil
}

func (dt *DataType) dataType() string {
	switch dt.v.Primitive {
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
		return dt.v.DataType
	default:
		return types.Generic
	}
}
