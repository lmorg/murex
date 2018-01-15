package typemgmt

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"regexp"
)

func init() {
	proc.GoFunctions["globals"] = cmdGlobals
	proc.GoFunctions["set"] = cmdSet
	proc.GoFunctions["!set"] = cmdUnset
	proc.GoFunctions["export"] = cmdExport
	proc.GoFunctions["!export"] = cmdUnexport
	proc.GoFunctions["unset"] = cmdUnexport
}

var (
	rxSet     *regexp.Regexp = regexp.MustCompile(`(?sm)^([_a-zA-Z0-9]+)\s*=(.*$)`)
	rxVarName *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)
)

func cmdGlobals(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := utils.JsonMarshal(proc.GlobalVars.Dump(), p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	p.Stdout.Writeln(b)

	return nil
}

func cmdSet(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	params := p.Parameters.StringAll()

	// Set variable as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `set` as method.")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		dt := p.Stdin.GetDataType()
		return proc.GlobalVars.Set(params, string(b), dt)
	}

	// Only one parameter, so unset variable:
	if rxVarName.MatchString(params) {
		proc.GlobalVars.Unset(params)
		return nil
	}

	// Set variable as parameters:
	match := rxSet.FindAllStringSubmatch(params, -1)

	return proc.GlobalVars.Set(match[0][1], match[0][2], types.String)
}

func cmdUnset(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	proc.GlobalVars.Unset(varName)
	return nil
}

func cmdExport(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	params := p.Parameters.StringAll()

	// Set env as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `export` as method.")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		return os.Setenv(params, string(b))
	}

	// Only one parameter, so unset env:
	if rxVarName.MatchString(params) {
		return os.Unsetenv(params)
	}

	// Set env as parameters:
	match := rxSet.FindAllStringSubmatch(params, -1)
	return os.Setenv(match[0][1], match[0][2])
}

func cmdUnexport(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = os.Unsetenv(varName)
	return err
}
