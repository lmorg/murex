package io

import (
	"fmt"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	proc.GoFunctions["echo"] = cmdOut
	proc.GoFunctions["out"] = cmdOut
	proc.GoFunctions["("] = cmdOutNoCR
	proc.GoFunctions["tout"] = cmdTout
	proc.GoFunctions["err"] = cmdErr
	proc.GoFunctions["print"] = cmdPrint
}

func cmdOut(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stdout.Writeln([]byte(s))
	return
}

func cmdOutNoCR(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stdout.Write([]byte(s))
	return
}

func cmdTout(p *proc.Process) (err error) {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	s := p.Parameters.StringAllRange(1, -1)
	s = ansi.ExpandConsts(s)

	p.Stdout.SetDataType(dt)

	_, err = p.Stdout.Write([]byte(s))
	return
}

func cmdErr(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)
	p.ExitNum = 1

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stderr.Writeln([]byte(s))
	return
}

func cmdPrint(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = fmt.Println([]byte(s))
	return
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/
