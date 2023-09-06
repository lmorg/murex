package io

import (
	"errors"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	lang.DefineFunction("out", cmdOut, types.String)
	lang.DefineFunction("(", cmdOutNoCR, types.String)
	lang.DefineFunction("tout", cmdTout, types.Any)
	lang.DefineFunction("err", cmdErr, types.Null)
}

func cmdOut(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err = p.Stdout.Writeln([]byte(s))
	return
}

func cmdOutNoCR(p *lang.Process) error {
	s := p.Parameters.StringAll()

	if len(s) == 0 || !strings.HasSuffix(s, ")") {
		return errors.New("missing closing ')'")
	}

	p.Stdout.SetDataType(types.String)

	s = ansi.ExpandConsts(s[:len(s)-1])

	_, err := p.Stdout.Write([]byte(s))
	return err
}

func cmdTout(p *lang.Process) error {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	s := p.Parameters.StringAllRange(1, -1)
	s = ansi.ExpandConsts(s)

	p.Stdout.SetDataType(dt)

	_, err = p.Stdout.Write([]byte(s))
	return err
}

func cmdErr(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	p.ExitNum = 1

	s := p.Parameters.StringAll()
	s = ansi.ExpandConsts(s)

	_, err := p.Stderr.Writeln([]byte(s))
	return err
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/
