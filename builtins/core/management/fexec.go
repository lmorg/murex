package management

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

type feType struct {
	fn   func(*lang.Process, string, []string) error
	desc string
}

var fe map[string]feType

func init() {
	lang.DefineFunction("fexec", cmdFexec, types.Any)

	defaults.AppendProfile(`
        autocomplete set fexec { [{
            "DynamicDesc": ({ fexec help }),
            "FlagValues": {
				"function": [{
					"DynamicDesc": ({ autocomplete.functions }),
					"ListView": true
				}],
				"private": [{
					"DynamicDesc": ({ autocomplete.privates }),
					"ListView": true
				}],
				"builtin": [{
					"DynamicDesc": ({ autocomplete.builtins }),
					"ListView": true
				}]
			}
        }] }
	`)

	fe = map[string]feType{
		"function": {
			desc: "Execute a murex public function",
			fn:   feFunction,
		},

		"private": {
			desc: "Execute a murex private function",
			fn:   fePrivate,
		},

		/*"event": {
			desc: "event",
			fn: func(p *lang.Process, cmd string, params []string) error {
				return fmt.Errorf("TODO: function not written yet")
			},
		},*/

		"builtin": {
			desc: "Execute a murex builtin",
			fn:   feBuiltin,
		},

		"help": {
			desc: "Display help message",
			fn:   feHelp,
		},
	}

}

func cmdFexec(p *lang.Process) error {
	flag, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if fe[flag].fn == nil {
		return fmt.Errorf("invalid flag '%s'. Use 'help' for more help", flag)
	}

	cmd, err := p.Parameters.String(1)
	if err != nil && flag != "help" {
		return err
	}

	var params []string
	if p.Parameters.Len() > 2 {
		params = p.Parameters.StringArray()[2:]
	}

	return fe[flag].fn(p, cmd, params)
}

func feBlock(p *lang.Process, block []rune, cmd string, params []string, fileRef *ref.File) (err error) {
	fork := p.Fork(lang.F_FUNCTION)
	fork.Name.Set(cmd)
	fork.Parameters.DefineParsed(params)
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return
}

func feFunction(p *lang.Process, cmd string, params []string) error {
	block, err := lang.MxFunctions.Block(cmd)
	if err != nil {
		return err
	}

	return feBlock(p, block, cmd, params, p.FileRef)
}

func fePrivate(p *lang.Process, cmd string, params []string) error {
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

	pvt := lang.PrivateFunctions.GetString(mod[1], mod[0])
	if pvt == nil {
		return lang.ErrPrivateNotFound(mod[0])
	}

	return feBlock(p, pvt.Block, mod[1], params, pvt.FileRef)
}

func feBuiltin(p *lang.Process, cmd string, params []string) error {
	if lang.GoFunctions[cmd] == nil {
		return fmt.Errorf("no builtin exists with the name `%s`", cmd)
	}

	fork := p.Fork(lang.F_DEFAULTS)
	fork.Name.Set(cmd)
	fork.Parameters.DefineParsed(params)
	fork.FileRef = p.FileRef
	return lang.GoFunctions[cmd](fork.Process)
}

func feHelp(p *lang.Process, _ string, _ []string) error {
	p.Stdout.SetDataType(types.Json)

	v := make(map[string]string)

	for name := range fe {
		v[name] = fe[name].desc
	}

	b, err := json.Marshal(v, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
