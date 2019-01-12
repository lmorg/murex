package management

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/parser"
)

func init() {
	proc.GoFunctions["args"] = cmdArgs
	proc.GoFunctions["params"] = cmdParams
	proc.GoFunctions["source"] = cmdSource
	proc.GoFunctions["."] = cmdSource
	proc.GoFunctions["version"] = cmdVersion
	proc.GoFunctions["murex-parser"] = cmdParser
	proc.GoFunctions["summary"] = cmdSummary
}

func cmdArgs(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Boolean)

	if p.Parameters.Len() != 1 {
		return errors.New("Invalid parameters. Expecting JSON input")
	}

	var args parameters.Arguments
	err = json.Unmarshal(p.Parameters.ByteAll(), &args)
	if err != nil {
		return err
	}

	type flags struct {
		Self       string
		Flags      map[string]string
		Additional []string
		Error      string
	}
	var jObj flags

	params := p.Scope.Parameters.Params
	if p.Scope.Id == 0 && len(params) > 0 {
		jObj.Self = params[0]
		if len(params) == 1 {
			params = []string{}
		} else {
			params = params[1:]
		}
	} else {
		jObj.Self = p.Scope.Name
	}

	jObj.Flags, jObj.Additional, err = parameters.ParseFlags(params, &args)
	if err != nil {
		jObj.Error = err.Error()
		p.ExitNum = 1
	}

	b, err := json.Marshal(jObj, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	err = p.Scope.Variables.Set("ARGS", string(b), types.Json)
	return err
}

func cmdParams(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	params := append([]string{p.Scope.Name}, p.Scope.Parameters.Params...)

	//debug.Json("builtin.params:", params)

	b, err := json.Marshal(&params, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdSource(p *proc.Process) error {
	var block []rune

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

	} else {
		var err error
		block, err = p.Parameters.Block(0)

		if err != nil {
			// get block from file
			name, err := p.Parameters.String(0)
			if err != nil {
				return err
			}

			file, err := os.Open(name)
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			block = []rune(string(b))
		}

	}

	var err error
	p.RunMode = runmode.Shell
	p.ExitNum, err = lang.RunBlockShellConfigSpace(block, nil, p.Stdout, p.Stderr)
	return err
}

func cmdVersion(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	_, err := p.Stdout.Writeln([]byte(config.AppName + ": " + config.Version))
	return err
}

func cmdParser(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	var (
		block []rune
		pos   int
		//syntaxHighlight bool
	)

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

		//syntaxHighlight, _ = p.Parameters.Bool(0)

	} else {
		r, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}
		block = r
		pos, _ = p.Parameters.Int(1)
	}

	pt, _ /*ansiHighlighted*/ := parser.Parse(block, pos)

	/*if syntaxHighlight {
		_, err := p.Stdout.Write([]byte(ansiHighlighted))
		return err
	}*/

	b, err := json.Marshal(pt, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdSummary(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	summary, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	shell.Summary[exe] = summary

	return nil
}
