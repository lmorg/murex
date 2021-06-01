package management

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/runmode"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/parser"
)

func init() {
	lang.GoFunctions["args"] = cmdArgs
	lang.GoFunctions["source"] = cmdSource
	lang.GoFunctions["."] = cmdSource
	lang.GoFunctions["version"] = cmdVersion
	lang.GoFunctions["murex-parser"] = cmdParser
	lang.GoFunctions["summary"] = cmdSummary
	lang.GoFunctions["!summary"] = cmdBangSummary

	defaults.AppendProfile(`
		autocomplete set version { [{
			"Flags": [ "--short", "--no-app-name", "--license", "--copyright" ]
		}] }
	`)
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

	params := p.Scope.Parameters.StringArray()
	if p.Scope.Id == 0 && len(params) > 0 {
		jObj.Self = params[0]
		if len(params) == 1 {
			params = []string{}
		} else {
			params = params[1:]
		}
	} else {
		jObj.Self = p.Scope.Name.String()
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

	//return p.Scope.Variables.Set(varName, b, types.Json)
	return p.Variables.Set(p, varName, b, types.Json)
}

func quickHash(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func cmdSource(p *lang.Process) error {
	var (
		block []rune
		name  string
		err   error
		b     []byte
	)

	if p.IsMethod {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))
		name = "<stdin>"

	} else {
		block, err = p.Parameters.Block(0)
		if err == nil {
			b = []byte(string(block))
			name = "N/A"

		} else {
			// get block from file
			name, err = p.Parameters.String(0)
			if err != nil {
				return err
			}

			file, err := os.Open(name)
			if err != nil {
				return err
			}

			b, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			block = []rune(string(b))
		}
	}

	module := quickHash(name + time.Now().String())
	fileRef := &ref.File{Source: ref.History.AddSource(name, "source/"+module, b)}

	p.RunMode = runmode.Normal
	fork := p.Fork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_NO_STDIN)

	fork.Name = p.Name
	fork.FileRef = fileRef
	p.ExitNum, err = fork.Execute(block)
	return err
}

var rxVersionNum = regexp.MustCompile(`^[0-9]+\.[0-9]+`)

func cmdVersion(p *lang.Process) error {
	s, _ := p.Parameters.String(0)

	switch s {

	case "--short":
		p.Stdout.SetDataType(types.Number)
		num := rxVersionNum.FindStringSubmatch(app.Version)
		if len(num) != 1 {
			return errors.New("Unable to extract version number from string")
		}
		_, err := p.Stdout.Write([]byte(num[0]))
		return err

	case "--no-app-name":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Version))
		return err

	case "--license":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.License))
		return err

	case "--copyright":
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Writeln([]byte(app.Copyright))
		return err

	case "":
		p.Stdout.SetDataType(types.String)
		v := fmt.Sprintf("%s: %s\n%s\n%s", app.Name, app.Version, app.License, app.Copyright)
		_, err := p.Stdout.Writeln([]byte(v))
		return err

	default:
		return fmt.Errorf("Not a valid parameter: %s", s)
	}

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

func cmdBangSummary(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return shell.Summary.Delete(exe)
}
