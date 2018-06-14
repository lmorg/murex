package typemgmt

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["datatype"] = cmdSetDt
	proc.GoFunctions["exec"] = proc.External
	proc.GoFunctions["pty"] = proc.External
	proc.GoFunctions["die"] = cmdDie
	proc.GoFunctions["exit"] = cmdExit
	proc.GoFunctions["null"] = cmdNull
	proc.GoFunctions["true"] = cmdTrue
	proc.GoFunctions["false"] = cmdFalse
	proc.GoFunctions["!"] = cmdNot
	proc.GoFunctions["cast"] = cmdCast
}

func cmdSetDt(p *proc.Process) error {
	dt := p.Parameters.StringAll()
	//p.Scope.Stdout.SetDataType(dt)
	//p.Parent.Stdout.SetDataType(dt)
	p.Stdout.SetDataType(dt)
	return nil
}

func cmdNull(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	p.Stdin.ReadAll()
	return nil
}

func cmdTrue(p *proc.Process) error {
	if s, _ := p.Parameters.String(0); s != "-s" {
		p.Stdout.SetDataType(types.Boolean)
		p.Stdout.Writeln(types.TrueByte)
	} else {
		p.Stdout.SetDataType(types.Null)
	}

	return nil
}

func cmdFalse(p *proc.Process) error {
	if s, _ := p.Parameters.String(0); s != "-s" {
		p.Stdout.SetDataType(types.Boolean)
		p.Stdout.Writeln(types.FalseByte)
	} else {
		p.Stdout.SetDataType(types.Null)
	}

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

	os.Exit(1)
	return nil
}

func cmdExit(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	i, _ := p.Parameters.Int(0)

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
