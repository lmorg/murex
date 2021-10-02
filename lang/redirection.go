package lang

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
)

//var rxNamedPipe = regexp.MustCompile(`^<(state_|test_|\!)?[a-zA-Z0-9]+>$`)

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
			pipeErr(p, fmt.Sprintf("invalid format used: '%s'", p.Parameters.Tokens[i][0].Key))
			continue
		}

		name := p.Parameters.Tokens[i][0].Key[1 : l-1]

		/*if !rxNamedPipe.MatchString(p.Parameters.Tokens[i][0].Key) ||
			len(name) < 5 || name[:4] != "env:" {

			err := parseRedirectionTemp(p, name)
			if err != nil {
				p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + err.Error()))
			}
			continue
		}*/

		switch {
		case len(name) > 5 && name[:5] == "test_":
			if p.NamedPipeTest == "" {
				testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
				if err == nil && testEnabled.(bool) {
					p.NamedPipeTest = name[5:]
				}
			} else {
				pipeErr(p, "you specified test multiple times")
			}

		case len(name) > 6 && name[:6] == "state_":
			if p.NamedPipeTest == "" {
				testEnabled, err := p.Config.Get("test", "enabled", types.Boolean)
				if err == nil && testEnabled.(bool) {
					p.testState = append(p.testState, name[6:])
				}
			}

		case len(name) > 4 && name[:4] == "env:":
			p.Exec.Env = append(p.Exec.Env, name[4:])

		case name[0] == '!':
			if p.NamedPipeErr == "" {
				p.NamedPipeErr = name[1:]
			} else {
				pipeErr(p, "you specified stderr multiple times")
			}

		case strings.Contains(name, ":"):
			err := parseRedirectionTemp(p, name)
			if err != nil {
				pipeErr(p, err.Error())
			}

		default:
			if p.NamedPipeOut == "" {
				p.NamedPipeOut = name
			} else {
				pipeErr(p, "you specified stdout multiple times")
			}
		}
	}
}

func parseRedirectionTemp(p *Process, namedPipe string) error {
	split := strings.SplitN(namedPipe, ":", 2)
	if len(split) != 2 {
		return fmt.Errorf("invalid format used: '%s'", namedPipe)
	}

	name := fmt.Sprintf("tmp:%d/%.3x", p.Id, md5.Sum([]byte(namedPipe)))

	err := GlobalPipes.CreatePipe(name, split[0], split[1])
	if err != nil {
		return err
	}

	p.NamedPipeOut = name

	return nil
}

func pipeErr(p *Process, msg string) {
	p.Stderr.Writeln([]byte("Invalid usage of named pipes: " + msg))
}
