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
		for len(p.Parameters.Tokens[0]) > 0 && i < 2 {

			if p.Parameters.Tokens[0][0].Type == parameters.TokenTypeValue && rxNamedPipe.MatchString(p.Parameters.Tokens[0][0].Key) {
				name := p.Parameters.Tokens[0][0].Key[1 : len(p.Parameters.Tokens[0][0].Key)-1]
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
