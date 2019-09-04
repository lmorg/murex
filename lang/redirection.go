package lang

import (
	"regexp"

	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
)

var rxNamedPipe = regexp.MustCompile(`^<(state_|test_|\!)?[a-zA-Z0-9]+>$`)

func parseRedirection(p *Process) {
	//p.NamedPipeOut = "out"
	//p.NamedPipeErr = "err"

	for i := range p.Parameters.Tokens {
		if p.Parameters.Tokens[i][0].Type != parameters.TokenTypeNamedPipe || !rxNamedPipe.MatchString(p.Parameters.Tokens[i][0].Key) {
			break
		}

		name := p.Parameters.Tokens[i][0].Key[1 : len(p.Parameters.Tokens[i][0].Key)-1]

		switch {
		case len(name) > 5 && name[:5] == "test_":
			if p.NamedPipeTest == "" {
				testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
				if err == nil && testEnabled.(bool) {
					p.NamedPipeTest = name[5:]
				}
			} else {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined test multiple times."))
			}

		case len(name) > 6 && name[:6] == "state_":
			if p.NamedPipeTest == "" {
				testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
				if err == nil && testEnabled.(bool) {
					p.testState = append(p.testState, name[6:])
				}
			}

		case name[0] == '!':
			if p.NamedPipeErr == "" {
				p.NamedPipeErr = name[1:]
			} else {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined stderr multiple times."))
			}

		default:
			if p.NamedPipeOut == "" {
				p.NamedPipeOut = name
			} else {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined stdout multiple times."))
			}
		}
	}
}
