package numeric

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshalInt(p *lang.Process, v any) ([]byte, error)    { return marshal(v, types.Integer) }
func marshalFloat(p *lang.Process, v any) ([]byte, error)  { return marshal(v, types.Float) }
func marshalNumber(p *lang.Process, v any) ([]byte, error) { return marshal(v, types.Number) }

func marshal(v any, dataType string) ([]byte, error) {
	i, err := types.ConvertGoType(v, dataType)
	if err != nil {
		return []byte{'0'}, err
	}

	s, err := types.ConvertGoType(i, types.String)
	return []byte(s.(string)), err
}

func unmarshalInt(p *lang.Process) (any, error)    { return unmarshal(p, types.Integer) }
func unmarshalFloat(p *lang.Process) (any, error)  { return unmarshal(p, types.Float) }
func unmarshalNumber(p *lang.Process) (any, error) { return unmarshal(p, types.Number) }

func unmarshal(p *lang.Process, dataType string) (any, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return 0, err
	}

	return types.ConvertGoType(b, dataType)
}
