package io

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	lang.GoFunctions["echo"] = cmdOut
	lang.GoFunctions["out"] = cmdOut
	lang.GoFunctions["("] = cmdOutNoCR
	lang.GoFunctions["tout"] = cmdTout
	lang.GoFunctions["err"] = cmdErr
}

func cmdOut(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stdout.Writeln([]byte(s))
	return
}

func cmdOutNoCR(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stdout.Write([]byte(s))
	return
}

func cmdTout(p *lang.Process) (err error) {
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

func cmdErr(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Null)
	p.ExitNum = 1

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stderr.Writeln([]byte(s))
	return
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/
