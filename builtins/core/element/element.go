package element

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("[[", element, types.Unmarshal, types.Marshal)
}

func element(p *lang.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic caught, please report this to https://github.com/lmorg/murex/issues : %s", r)
		}
	}()

	dt := p.Stdin.GetDataType()
	//p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	params := p.Parameters.StringArray()

	switch len(params) {
	case 0:
		return errors.New("missing parameter. Requires nested index")

	case 1:
		if strings.HasSuffix(params[0], "]]") {
			params[0] = params[0][0 : len(params[0])-2]
		} else {
			return fmt.Errorf("missing closing brackets, ` ]]`\nExpression: %s", p.Parameters.StringAll())
		}

	case 2:
		last := len(params) - 1
		if strings.HasSuffix(params[last], "]]") {
			if len(params[last]) != 2 {
				return errors.New("too many parameters")
			}
		} else {
			return fmt.Errorf("missing closing brackets, ` ]]`\nExpression: %s", p.Parameters.StringAll())
		}

	default:
		return errors.New("too many parameters")
	}

	obj, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	obj, err = lang.ElementLookup(obj, params[0])
	if err != nil {
		return err
	}

	switch v := obj.(type) {
	case string:
		p.Stdout.SetDataType(types.String)
		_, err = p.Stdout.Write([]byte(v))
	case []byte:
		p.Stdout.SetDataType(types.String)
		_, err = p.Stdout.Write(v)
	case int:
		p.Stdout.SetDataType(types.Integer)
		_, err = p.Stdout.Write([]byte(strconv.Itoa(v)))
	case int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		p.Stdout.SetDataType(types.Integer)
		_, err = fmt.Fprint(p.Stdout, v)
	case float32:
		p.Stdout.SetDataType(types.Float)
		_, err = p.Stdout.Write([]byte(types.FloatToString(float64(v))))
	case float64:
		p.Stdout.SetDataType(types.Float)
		_, err = p.Stdout.Write([]byte(types.FloatToString(v)))
	case bool:
		p.Stdout.SetDataType(types.Boolean)
		if v {
			_, err = p.Stdout.Write(types.TrueByte)
		} else {
			_, err = p.Stdout.Write(types.FalseByte)
		}
	default:
		p.Stdout.SetDataType(dt)
		b, err := lang.MarshalData(p, dt, obj)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		if err != nil {
			return err
		}
	}

	return err
}
