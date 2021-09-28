package arraytools

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("len", cmdLen, types.Unmarshal, types.Integer)
}

func cmdLen(p *lang.Process) error {
	p.Stdout.SetDataType(types.Integer)

	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return err
	}

	var i int

	switch v.(type) {
	case []int:
		i = len(v.([]int))

	case []float64:
		i = len(v.([]float64))

	case []string:
		i = len(v.([]string))

	case [][]string:
		i = len(v.([][]string))

	case []interface{}:
		i = len(v.([]interface{}))

	case map[string]string:
		i = len(v.(map[string]string))

	case map[string]interface{}:
		i = len(v.(map[string]interface{}))

	case map[interface{}]string:
		i = len(v.(map[interface{}]string))

	case map[interface{}]interface{}:
		i = len(v.(map[interface{}]interface{}))

	default:
		return fmt.Errorf("I don't know how to get `len` for `%T`", v)
	}

	_, err = p.Stdout.Write([]byte(strconv.Itoa(i)))
	return err
}
