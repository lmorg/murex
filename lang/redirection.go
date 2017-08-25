package lang

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"regexp"
)

var rxNamedPipe *regexp.Regexp = regexp.MustCompile(`^<[\!]?[a-zA-Z0-9]+>$`)

func parseRedirection(p *proc.Process) {
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
				if name[0] == '!' {
					if p.NamedPipeErr == "" {
						p.NamedPipeErr = name[1:]
					} else {
						p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined stderr multiple times."))
					}

				} else {
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
