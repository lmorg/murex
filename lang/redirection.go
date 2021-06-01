package lang

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

var rxNamedPipe = regexp.MustCompile(`^<(state_|test_|\!)?[a-zA-Z0-9]+>$`)

func parseRedirection(p *Process) {
	//p.NamedPipeOut = "out"
	//p.NamedPipeErr = "err"

	for i := range p.Parameters.Tokens {
		// nil tokens can sometimes get popped into the token array. This is really a "bug" of the parser where
		// speed is valued over correctness. However it does mean we need to ignore them here
		if p.Parameters.Tokens[i][0].Type == parameters.TokenTypeNil {
			continue
		}

		if p.Parameters.Tokens[i][0].Type != parameters.TokenTypeNamedPipe {
			break
		}

		l := len(p.Parameters.Tokens[i][0].Key)
		if l < 2 || p.Parameters.Tokens[i][0].Key[l-1] != '>' {
			p.Stderr.Writeln([]byte(fmt.Sprintf("Invalid format used in named pipe: '%s'", p.Parameters.Tokens[i][0].Key)))
			continue
		}

		name := p.Parameters.Tokens[i][0].Key[1 : l-1]

		if !rxNamedPipe.MatchString(p.Parameters.Tokens[i][0].Key) {
			err := parseRedirectionTemp(p, name)
			if err != nil {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
			}
			continue
		}

		switch {
		case len(name) > 5 && name[:5] == "test_":
			if p.NamedPipeTest == "" {
				testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
				if err == nil && testEnabled.(bool) {
					p.NamedPipeTest = name[5:]
				}
			} else {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined test multiple times"))
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
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined stderr multiple times"))
			}

		default:
			if p.NamedPipeOut == "" {
				p.NamedPipeOut = name
			} else {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: you defined stdout multiple times"))
			}
		}
	}
}

func parseRedirectionTemp(p *Process, namedPipe string) error {
	split := strings.SplitN(namedPipe, ":", 2)
	if len(split) != 2 {
		return fmt.Errorf("Invalid format used in named pipe: '%s'", namedPipe)
	}

	name := fmt.Sprintf("tmp:%d/%.3x", p.Id, md5.Sum([]byte(namedPipe)))

	err := GlobalPipes.CreatePipe(name, split[0], split[1])
	if err != nil {
		return err
	}

	p.NamedPipeOut = name

	return nil
}
