package io

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["echo"] = proc.GoFunction{Func: cmdOut, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["out"] = proc.GoFunction{Func: cmdOut, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["err"] = proc.GoFunction{Func: cmdErr, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["print"] = proc.GoFunction{Func: cmdPrint, TypeIn: types.Null, TypeOut: types.Null}
}

func cmdOut(p *proc.Process) (err error) {
	if f, _ := p.Parameters.String(0); f == "-n" {
		_, err = p.Stdout.Write(p.Parameters[1:].AllByte())
		return
	}
	_, err = p.Stdout.Writeln(p.Parameters.AllByte())
	return
}

func cmdErr(p *proc.Process) (err error) {
	if f, _ := p.Parameters.String(0); f == "-n" {
		_, err = p.Stdout.Write(p.Parameters[1:].AllByte())
		return
	}
	_, err = p.Stderr.Writeln(p.Parameters.AllByte())
	return
}

func cmdPrint(p *proc.Process) (err error) {
	_, err = fmt.Println(p.Parameters.AllByte())
	return
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/
