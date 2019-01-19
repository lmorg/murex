package string

import "github.com/lmorg/murex/lang"

func index(p *lang.Process, params []string) error {
	match := make(map[string]bool)
	for i := range params {
		match[params[i]] = true
	}

	err := p.Stdin.ReadMap(p.Config, func(key, value string, last bool) {
		if p.IsNot {
			if !match[key] {
				p.Stdout.Writeln([]byte(value))
			}
		} else {
			if match[key] {
				p.Stdout.Writeln([]byte(value))
			}
		}
	})

	return err
}
