package lang

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"regexp"
)

var rxNamedPipe *regexp.Regexp = regexp.MustCompile(`^<[\!]?[a-zA-Z0-9]+>$`)

func parseRedirection(p *proc.Process) {
	if len(p.Parameters.Tokens) > 0 {
		var i int
		for len(p.Parameters.Tokens[0]) > 0 && i < 2 {

			if p.Parameters.Tokens[0][0].Type == parameters.TokenTypeValue && rxNamedPipe.MatchString(p.Parameters.Tokens[0][0].Key) {
				name := p.Parameters.Tokens[0][0].Key[1 : len(p.Parameters.Tokens[0][0].Key)-1]
				if name[0] == '!' {
					switch name[1:] {
					case "out":
						p.Stderr = p.Stdout
					case "err":
						p.Stderr.Writeln([]byte("Invalid usage of named pipes: stderr defaults to <err>."))
					default:
						pipe, err := proc.GlobalPipes.Get(name[1:])
						if err == nil {
							p.Stderr = pipe
						} else {
							p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
						}
					}

				} else {
					switch name {
					case "out":
						p.Stderr.Writeln([]byte("Invalid usage of named pipes: stdout defaults to <out>."))
					case "err":
						p.Stdout = p.Stderr
					default:
						pipe, err := proc.GlobalPipes.Get(name)
						if err == nil {
							p.Stdout = pipe
						} else {
							p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
						}
					}
				}

				if len(p.Parameters.Tokens) > 1 {
					p.Parameters.Tokens = p.Parameters.Tokens[1:]
				} else {
					p.Parameters.Tokens[0][0].Type = parameters.TokenTypeNil
				}
				i++
			} else {
				break
			}
		}
	}
}
