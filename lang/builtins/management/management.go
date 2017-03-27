package management

import (
	"encoding/json"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"os"
	"strconv"
)

func init() {
	proc.GoFunctions["debugmode"] = proc.GoFunction{Func: cmdDebugMode, TypeIn: types.Null, TypeOut: types.Boolean}
	proc.GoFunctions["exitnum"] = proc.GoFunction{Func: cmdExitNum, TypeIn: types.Generic, TypeOut: types.Integer}
	proc.GoFunctions["config"] = proc.GoFunction{Func: cmdConfig, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["builtins"] = proc.GoFunction{Func: cmdListBuiltins, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["cd"] = proc.GoFunction{Func: cmdCd, TypeIn: types.Null, TypeOut: types.Null}
}

func cmdDebugMode(p *proc.Process) error {
	v, err := p.Parameters.Bool(0)
	if err != nil {
		p.Stdout.Writeln(types.FalseByte)
		p.ExitNum = 1
		return nil
	}
	debug.Enable = v
	if v == false {
		p.Stdout.Writeln(types.FalseByte)
		p.ExitNum = 1
		return nil
	}
	p.Stdout.Writeln(types.TrueByte)
	return nil
}

func cmdExitNum(p *proc.Process) error {
	p.Stdout.Writeln([]byte(strconv.Itoa(p.Previous.ExitNum)))
	return nil
}

func cmdListBuiltins(p *proc.Process) error {
	for name := range proc.GoFunctions {
		p.Stdout.Writeln([]byte(name))
	}

	return nil
}

func cmdConfig(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		b, err := json.MarshalIndent(proc.GlobalConf.Dump(), "", "\t")
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	return nil
}

func cmdCd(p *proc.Process) error {
	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = os.Chdir(s)
	return err
}
