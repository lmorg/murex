package lang

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

func parseRedirection(p *Process) {
	//p.NamedPipeOut = "out"
	//p.NamedPipeErr = "err"

	for _, name := range p.namedPipes {
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
			p.Envs = append(p.Envs, name[4:])

		case len(name) > 4 && name[:4] == "fid:":
			varName := name[4:]
			err := p.Variables.Set(p, varName, p.Id, types.Integer)
			if err != nil {
				ShellProcess.Stderr.Writeln([]byte(
					fmt.Sprintf("Cannot write variable '%s': %s", varName, err.Error()),
				))
			}

		case len(name) > 4 && name[:4] == "pid:":
			panic("TODO")
			/*varName := name[4:]
			p.Exec.Callback = func(pid int) {
				err := p.Variables.Set(p, varName, pid, types.Integer)
				if err != nil {
					ShellProcess.Stderr.Writeln([]byte(
						fmt.Sprintf("Cannot write variable '%s': %s", varName, err.Error()),
					))
				}
			}*/

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
