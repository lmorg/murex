package typemgmt

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"io"
	"os"
)

func init() {
	proc.GoFunctions["exec"] = proc.GoFunction{Func: proc.External, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["pty"] = proc.GoFunction{Func: proc.ExternalPty, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["die"] = proc.GoFunction{Func: cmdDie, TypeIn: types.Generic, TypeOut: types.Die}
	proc.GoFunctions["exit"] = proc.GoFunction{Func: cmdExit, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["null"] = proc.GoFunction{Func: cmdNull, TypeIn: types.Generic, TypeOut: types.Null}
	proc.GoFunctions["true"] = proc.GoFunction{Func: cmdTrue, TypeIn: types.Null, TypeOut: types.Boolean}
	proc.GoFunctions["false"] = proc.GoFunction{Func: cmdFalse, TypeIn: types.Null, TypeOut: types.Boolean}
	proc.GoFunctions["!"] = proc.GoFunction{Func: cmdNot, TypeIn: types.Generic, TypeOut: types.Boolean}
	proc.GoFunctions["untype"] = proc.GoFunction{Func: cmdUntype, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdNull(*proc.Process) error {
	return nil
}

func cmdTrue(p *proc.Process) error {
	p.Stdout.Writeln(types.TrueByte)
	return nil
}

func cmdFalse(p *proc.Process) error {
	p.Stdout.Writeln(types.FalseByte)
	p.ExitNum = 1
	return nil
}

func cmdNot(p *proc.Process) error {
	val := !types.IsTrue(p.Stdin.ReadAll(), p.Previous.ExitNum)
	if val {
		p.Stdout.Writeln(types.TrueByte)
	} else {
		p.Stdout.Writeln(types.FalseByte)
	}
	return nil
}

func cmdDie(*proc.Process) error {
	if shell.Instance != nil {
		shell.Instance.Terminal.ExitRawMode()
	}
	os.Exit(1)
	return nil
}

func cmdExit(p *proc.Process) error {
	i, _ := p.Parameters.Int(0)
	if shell.Instance != nil {
		shell.Instance.Terminal.ExitRawMode()
	}
	os.Exit(i)
	return nil
}

func cmdUntype(p *proc.Process) (err error) {
	_, err = io.Copy(p.Stdout, p.Stdin)
	return
}
