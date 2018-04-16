package numeric

import (
	"strconv"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func marshalInt(_ *proc.Process, i interface{}) (b []byte, err error) {
	s := strconv.Itoa(i.(int))
	return []byte(s), nil
}

func marshalFloat(_ *proc.Process, f interface{}) (b []byte, err error) {
	s := types.FloatToString(f.(float64))
	return []byte(s), nil
}

func unmarshalInt(p *proc.Process) (interface{}, error)    { return unmarshal(p, types.Integer) }
func unmarshalFloat(p *proc.Process) (interface{}, error)  { return unmarshal(p, types.Float) }
func unmarshalNumber(p *proc.Process) (interface{}, error) { return unmarshal(p, types.Number) }

func unmarshal(p *proc.Process, dataType string) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return 0, err
	}

	return types.ConvertGoType(b, dataType)
}
