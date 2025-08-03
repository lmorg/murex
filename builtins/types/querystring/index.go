package string

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func index(p *lang.Process, params []string) error {
	match := make(map[string]bool)
	for i := range params {
		match[params[i]] = true
	}

	var (
		v       any
		typeErr error
	)

	err := p.Stdin.ReadMap(p.Config, func(m *stdio.Map) {
		if match[m.Key] {
			v, typeErr = types.ConvertGoType(m.Value, types.String)
			if typeErr != nil {
				p.Done()
				return
			}

			p.Stdout.Writeln([]byte(v.(string)))
		}
	})

	if typeErr != nil {
		return typeErr
	}

	return err
}
