package unicode

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func marshal(_ *lang.Process, iface interface{}) (b []byte, err error) {
	switch v := iface.(type) {
	case string:
		return []byte(v), nil

	case int, float64, bool, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32:
		return []byte(fmt.Sprint(v)), nil

	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case []interface{}:
		for i := range v {
			b = append(b, []byte(fmt.Sprint(v[i])+utils.NewLineString)...)
		}
		return

	case interface{}:
		return []byte(fmt.Sprint(v)), nil

	default:
		err = fmt.Errorf("I don't know how to marshal %T into a `%s`. Data types are possibly incompatible?", v, types.Generic)
		return
	}
}

func unmarshal(p *lang.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	return string(b), err
}
