package management

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/lmorg/murex/lang/ref"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/parser"
)

func init() {
	lang.GoFunctions["args"] = cmdArgs
	lang.GoFunctions["params"] = cmdParams
	lang.GoFunctions["source"] = cmdSource
	lang.GoFunctions["."] = cmdSource
	lang.GoFunctions["version"] = cmdVersion
	lang.GoFunctions["murex-parser"] = cmdParser
	lang.GoFunctions["summary"] = cmdSummary
}

func cmdArgs(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Boolean)

	if p.Parameters.Len() != 2 {
		return errors.New("Invalid parameters. Usage: args var_name { json }")
	}

	varName, _ := p.Parameters.String(0)

	var args parameters.Arguments
	b, _ := p.Parameters.Byte(1)
	err = json.UnmarshalMurex(b, &args)
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

	b, err = json.Marshal(jObj, false)
	if err != nil {
		return err
	}

	return p.Scope.Variables.Set(varName, b, types.Json)
}

func cmdParams(p *lang.Process) error {
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

func quickHash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func cmdSource(p *lang.Process) error {
	var block []rune

	fileRef := p.FileRef

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

			module := quickHash(name + time.Now().String())

			fileRef = &ref.File{Source: ref.History.AddSource(name, "source/"+module, b)}
		}

	}

	var err error
	p.RunMode = runmode.Shell
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
	fork.Stdout = p.Stdout
	fork.Stderr = p.Stderr
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return err
}

func cmdVersion(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)
	_, err := p.Stdout.Writeln([]byte(config.AppName + ": " + config.Version))
	return err
}

func cmdParser(p *lang.Process) error {
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

	//_, err = p.Stdout.Write(b)

	/*// start new parser
	nodes, _ := lang.ParseBlock(block)
	b, err = json.Marshal(nodes, p.Stdout.IsTTY())
	if err != nil {
		return err
	}
	// end new parser*/

	_, err = p.Stdout.Write(b)

	return err
}

func cmdSummary(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	summary, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	shell.Summary.Set(exe, summary)

	return nil
}
