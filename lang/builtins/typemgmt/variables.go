package typemgmt

import (
	"encoding/json"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["globals"] = proc.GoFunction{Func: cmdGlobals, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["set"] = proc.GoFunction{Func: cmdSet, TypeIn: types.Generic, TypeOut: types.Null}
}

func cmdSet(p *proc.Process) error {
	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = proc.GlobalVars.Set(s, string(p.Stdin.ReadAll()), p.Previous.ReturnType)

	return err
}

func cmdGlobals(p *proc.Process) error {
	b, err := json.MarshalIndent(proc.GlobalVars.Dump(), "", "\t")
	if err != nil {
		return err
	}

	p.Stdout.Writeln(b)

	return nil
}
