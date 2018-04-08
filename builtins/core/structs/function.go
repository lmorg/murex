package structs

import (
	"errors"
	"regexp"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["alias"] = cmdAlias
	proc.GoFunctions["!alias"] = cmdUnalias
	proc.GoFunctions["func"] = cmdFunc
	proc.GoFunctions["!func"] = cmdUnfunc
}

var rxAlias *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)=(.*?)[\s]*$`)

func cmdAlias(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)
		b, err := utils.JsonMarshal(proc.GlobalAliases.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		return err

	}

	p.Stdout.SetDataType(types.Null)

	s := p.Parameters.StringAll()

	if !rxAlias.MatchString(s) {
		return errors.New("Invalid syntax. Expecting `alias new_name=original_name parameter1 parameter2 ...`")
	}

	split := rxAlias.FindStringSubmatch(s)
	name := split[1]
	params := append([]string{split[2]}, p.Parameters.StringArray()[1:]...)

	proc.GlobalAliases.Add(name, params)
	return nil
}

func cmdUnalias(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	for _, name := range p.Parameters.StringArray() {
		err := proc.GlobalAliases.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmdFunc(p *proc.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	proc.MxFunctions.Define(name, block)
	return nil
}

func cmdUnfunc(p *proc.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return proc.MxFunctions.Undefine(name)
}
