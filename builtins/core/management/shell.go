package management

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/hintsummary"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/parser"
)

func init() {
	lang.DefineFunction("args", cmdArgs, types.Null)
	lang.DefineMethod("murex-parser", cmdParser, types.String, types.Json)
	lang.DefineFunction("summary", cmdSummary, types.Null)
	lang.DefineFunction("!summary", cmdBangSummary, types.Null)
	lang.DefineMethod("man-get-flags", cmdManParser, types.Text, types.Json)
}

func cmdArgs(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() != 2 {
		return errors.New("invalid parameters. Usage: args var_name { json }")
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

	hintsummary.Summary.Set(exe, summary)

	return nil
}

func cmdBangSummary(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return hintsummary.Summary.Delete(exe)
}

const manArgDescriptions = "--descriptions"

var manArgs = &parameters.Arguments{
	AllowAdditional: true,
	Flags: map[string]string{
		manArgDescriptions: types.Boolean,
		"-d":               manArgDescriptions,
	},
}

func cmdManParser(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	var (
		flags        []string
		descriptions = map[string]string{}
	)

	cmdFlags, additional, err := p.Parameters.ParseFlags(manArgs)
	if err != nil {
		return err
	}

	if p.IsMethod {
		flags = man.ParseByStdio(p.Stdin)

	} else {
		if len(additional) != 1 {
			return fmt.Errorf("invalid parameters")
		}
		exe := additional[0]

		paths := man.GetManPages(exe)
		flags, descriptions = man.ParseByPaths(exe, paths)
	}

	var b []byte
	if cmdFlags[manArgDescriptions] == types.TrueString {
		b, err = json.Marshal(descriptions, p.Stdout.IsTTY())
	} else {
		b, err = json.Marshal(flags, p.Stdout.IsTTY())
	}
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
