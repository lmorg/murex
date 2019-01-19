package lang

import (
	"regexp"

	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
)

var rxNamedPipe *regexp.Regexp = regexp.MustCompile(`^<(test_|\!)?[a-zA-Z0-9]+>$`)

func parseRedirection(p *Process) {
	//p.NamedPipeOut = "out"
	//p.NamedPipeErr = "err"

	if len(p.Parameters.Tokens) > 0 {
		var i int
		end := 2

		for i < end {

			if i == len(p.Parameters.Tokens) {
				break
			}

			if p.Parameters.Tokens[i][0].Type == parameters.TokenTypeNil {
				i++
				end++
				continue
			}

			if p.Parameters.Tokens[i][0].Type == parameters.TokenTypeValue && rxNamedPipe.MatchString(p.Parameters.Tokens[i][0].Key) {
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

				// Instead of deleting it, lets mark it as nil.
				p.Parameters.Tokens[i][0].Type = parameters.TokenTypeNil
				i++
			} else {
				break
			}
		}
	}
}
