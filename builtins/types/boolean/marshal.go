package boolean

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(_ *lang.Process, v interface{}) ([]byte, error) {
	i, err := types.ConvertGoType(v, types.Boolean)
	if err != nil {
		return types.FalseByte, err
	}

	s, err := types.ConvertGoType(i, types.String)
	return []byte(s.(string)), err
}

func unmarshal(p *lang.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return false, err
	}

	return types.ConvertGoType(b, types.Boolean)
}
