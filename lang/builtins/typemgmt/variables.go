package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
	"regexp"
)

func init() {
	proc.GoFunctions["globals"] = proc.GoFunction{Func: cmdGlobals, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["set"] = proc.GoFunction{Func: cmdSet, TypeIn: types.Generic, TypeOut: types.Null}
	proc.GoFunctions["export"] = proc.GoFunction{Func: cmdExport, TypeIn: types.Generic, TypeOut: types.Null}
}

func cmdGlobals(p *proc.Process) error {
	b, err := json.MarshalIndent(proc.GlobalVars.Dump(), "", "\t")
	if err != nil {
		return err
	}

	p.Stdout.Writeln(b)

	return nil
}

var (
	rxSet     *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*=(.*)`)
	rxVarName *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)
)

func cmdSet(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	params := p.Parameters.StringAll()

	// Set variable as method:
	if p.Method {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `set` as method.")
		}
		return proc.GlobalVars.Set(params, string(p.Stdin.ReadAll()), p.Previous.ReturnType)
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

func cmdExport(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name.")
	}

	params := p.Parameters.StringAll()

	// Set env as method:
	if p.Method {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `export` as method.")
		}
		return os.Setenv(params, string(p.Stdin.ReadAll()))
	}

	// Only one parameter, so unset env:
	if rxVarName.MatchString(params) {
		return os.Unsetenv(params)
	}

	// Set env as parameters:
	match := rxSet.FindAllStringSubmatch(params, -1)
	return os.Setenv(match[0][1], match[0][2])
}
