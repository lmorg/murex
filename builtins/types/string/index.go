package string

import "github.com/lmorg/murex/lang/proc"

func index(p *proc.Process, params []string) error {
	match := make(map[string]bool)
	for i := range params {
		match[params[i]] = true
	}

	err := p.Stdin.ReadMap(p.Config, func(key, value string, last bool) {
		if match[key] {
			p.Stdout.Writeln([]byte(value))
		}
	})

	return err
}
