package yaml

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	yaml "gopkg.in/yaml.v3"
)

func marshal(p *lang.Process, v any) ([]byte, error) {
	switch t := v.(type) {
	case [][]string:
		if lang.GetDataTypeLayout(p.Stdin.GetDataType()) != lang.DataTypeIsTable {
			break
		}

		var i int
		table := make([]map[string]any, len(t)-1)
		err := types.Table2Map(t, func(m map[string]any) error {
			table[i] = m
			i++
			return nil
		})
		if err != nil {
			return nil, err
		}
		return yaml.Marshal(table)

	}
	return yaml.Marshal(v)
}

func unmarshal(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &v)
	return
}
