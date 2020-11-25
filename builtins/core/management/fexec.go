package management

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
)

func init() {
	lang.GoFunctions["fexec"] = cmdFexec
}

var feFunction = map[string]func(*lang.Process, string, []string) error{
	"function": func(p *lang.Process, cmd string, params []string) error {
		block, err := lang.MxFunctions.Block(cmd)
		if err != nil {
			return err
		}

		return feBlock(p, block, cmd, params)
	},

	"private": func(p *lang.Process, cmd string, params []string) error {
		mod := strings.Split(cmd, "/")
		if len(mod) > 1 && mod[0] == "" {
			mod = mod[1:]
		}

		switch len(mod) {
		case 0, 1:
			mod = []string{"", cmd}

		case 2:
			// do nothing

		default:
			mod = []string{strings.Join(mod[0:len(mod)-1], "/"), mod[2]}
		}

		block, err := lang.PrivateFunctions.Block(mod[1], mod[0])
		if err != nil {
			return err
		}

		return feBlock(p, block, mod[1], params)
	},

	"event": func(p *lang.Process, cmd string, params []string) error {
		return fmt.Errorf("TODO: function not written yet")
	},

	"builtin": func(p *lang.Process, cmd string, params []string) error {
		if lang.GoFunctions[cmd] == nil {
			return fmt.Errorf("No builtin exists with the name `%s`", cmd)
		}

		fork := p.Fork(lang.F_DEFAULTS)
		fork.Name = cmd
		fork.Parameters = parameters.Parameters{Params: params}
		fork.FileRef = p.FileRef
		return lang.GoFunctions[cmd](fork.Process)
	},

	"help": func(p *lang.Process, cmd string, params []string) error {
		return fmt.Errorf("TODO: function not written yet")
	},
}

func feBlock(p *lang.Process, block []rune, cmd string, params []string) (err error) {
	fork := p.Fork(lang.F_FUNCTION)
	fork.Name = cmd
	fork.Parameters = parameters.Parameters{Params: params}
	fork.FileRef = p.FileRef
	p.ExitNum, err = fork.Execute(block)
	return
}

func cmdFexec(p *lang.Process) error {
	flag, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if feFunction[flag] == nil {
		return fmt.Errorf("Invalid flag '%s'. Use 'help' for more help", flag)
	}

	cmd, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	var params []string
	if p.Parameters.Len() > 2 {
		params = p.Parameters.Params[2:]
	}

	return feFunction[flag](p, cmd, params)
}
