package typemgmt

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"io"
	"os"
	"strings"
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
	proc.GoFunctions["cast"] = proc.GoFunction{Func: cmdCast, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdNull(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	p.Stdin.ReadAll()
	return nil
}

func cmdTrue(p *proc.Process) error {
	p.Stdout.SetDataType(types.Boolean)
	p.Stdout.Writeln(types.TrueByte)
	return nil
}

func cmdFalse(p *proc.Process) error {
	p.Stdout.SetDataType(types.Boolean)
	p.Stdout.Writeln(types.FalseByte)
	p.ExitNum = 1
	return nil
}

func cmdNot(p *proc.Process) error {
	p.Stdout.SetDataType(types.Boolean)

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	val := !types.IsTrue(b, p.Previous.ExitNum)
	if val {
		p.Stdout.Writeln(types.TrueByte)
	} else {
		p.Stdout.Writeln(types.FalseByte)
	}
	return nil
}

func cmdDie(p *proc.Process) error {
	p.Stdout.SetDataType(types.Die)

	if shell.Instance != nil {
		shell.Instance.Terminal.ExitRawMode()
	}
	os.Exit(1)
	return nil
}

func cmdExit(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	i, _ := p.Parameters.Int(0)
	if shell.Instance != nil {
		shell.Instance.Terminal.ExitRawMode()
	}
	os.Exit(i)
	return nil
}

func cmdCast(p *proc.Process) error {
	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	// Data types are lower case. So lets help people out a little here.
	dt := strings.ToLower(s)

	// Technically you could use the following values as data types, but it's unlikely anyone would intend to do so,
	// so lets just disable them with a helpful error to ease debugging.
	switch dt {
	case "string":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.String + "`?")
	case "number":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Number + "`?")
	case "integer":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Integer + "`?")
	case "boolean":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Boolean + "`?")
	case "code", "codeblock":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.CodeBlock + "`?")
	case "generic":
		return errors.New("`" + s + "` is an invalid data type. Presumably you meant `" + types.Generic + "`?")
	}

	p.Stdout.SetDataType(dt)
	_, err = io.Copy(p.Stdout, p.Stdin)
	return err
}
