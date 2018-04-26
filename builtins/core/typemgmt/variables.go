package typemgmt

import (
	"errors"
	"os"
	"regexp"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["set"] = cmdSet
	proc.GoFunctions["!set"] = cmdUnset
	proc.GoFunctions["global"] = cmdGlobal
	proc.GoFunctions["!global"] = cmdUnglobal
	proc.GoFunctions["export"] = cmdExport
	proc.GoFunctions["!export"] = cmdUnexport
	proc.GoFunctions["unset"] = cmdUnexport
}

var (
	rxSet     *regexp.Regexp = regexp.MustCompile(`(?sm)^([_a-zA-Z0-9]+)\s*=(.*$)`)
	rxVarName *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)
)

func cmdSet(p *proc.Process) error    { return set(p, p) }
func cmdGlobal(p *proc.Process) error { return set(p, proc.ShellProcess) }

func cmdUnset(p *proc.Process) error    { return unset(p, p) }
func cmdUnglobal(p *proc.Process) error { return unset(p, proc.ShellProcess) }

func set(p *proc.Process, scope *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	params := p.Parameters.StringAll()

	// Set variable as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `set` / `global` as method.")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		dt := p.Stdin.GetDataType()
		return p.Variables.Set(params, string(b), dt)
	}

	// Only one parameter, so unset variable:
	if rxVarName.MatchString(params) {
		err := p.Variables.Unset(params)
		return err
	}

	// Set variable as parameters:
	match := rxSet.FindAllStringSubmatch(params, -1)

	return scope.Variables.Set(match[0][1], match[0][2], types.String)
}

func unset(p *proc.Process, scope *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = scope.Variables.Unset(varName)
	return err
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

	/*// Only one parameter, so unset env:
	if rxVarName.MatchString(params) {
		return os.Unsetenv(params)
	}*/

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
