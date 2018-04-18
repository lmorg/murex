package numeric

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func marshalInt(p *proc.Process, v interface{}) ([]byte, error)    { return marshal(v, types.Integer) }
func marshalFloat(p *proc.Process, v interface{}) ([]byte, error)  { return marshal(v, types.Float) }
func marshalNumber(p *proc.Process, v interface{}) ([]byte, error) { return marshal(v, types.Number) }

func marshal(v interface{}, dataType string) ([]byte, error) {
	i, err := types.ConvertGoType(v, dataType)
	if err != nil {
		return []byte{'0'}, err
	}

	s, err := types.ConvertGoType(i, types.String)
	return []byte(s.(string)), err
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
