package io

import (
	"compress/gzip"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"io"
	"os"
	"time"
)

func init() {
	proc.GoFunctions["echo"] = proc.GoFunction{Func: cmdOut, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["out"] = proc.GoFunction{Func: cmdOut, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["err"] = proc.GoFunction{Func: cmdErr, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["print"] = proc.GoFunction{Func: cmdPrint, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["text"] = proc.GoFunction{Func: cmdText, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["open"] = proc.GoFunction{Func: cmdOpen, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["pt"] = proc.GoFunction{Func: cmdPipeTelemetry, TypeIn: types.Generic, TypeOut: types.Generic}
}

func cmdOut(p *proc.Process) (err error) {
	_, err = p.Stdout.Writeln(p.Parameters.AllByte())
	return
}

func cmdErr(p *proc.Process) (err error) {
	_, err = p.Stderr.Writeln(p.Parameters.AllByte())
	return
}

func cmdPrint(p *proc.Process) (err error) {
	_, err = fmt.Println(p.Parameters.AllByte())
	return
}

func cmdText(p *proc.Process) error {
	for _, filename := range p.Parameters {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
			gz, err := gzip.NewReader(file)
			if err != nil {
				file.Close()
				return err
			}
			_, err = io.Copy(p.Stdout, gz)
			file.Close()
			if err != nil {
				return err
			}

		} else {
			_, err = io.Copy(p.Stdout, file)
			file.Close()
			if err != nil {
				return err
			}

		}

	}

	return nil
}

func cmdOpen(p *proc.Process) error {
	for _, filename := range p.Parameters {
		file, err := os.Open(filename)
		if err != nil {
			file.Close()
			return err
		}
		_, err = io.Copy(p.Stdout, file)
		if err != nil {
			file.Close()
			return err
		}

		file.Close()
	}

	return nil
}

func cmdPipeTelemetry(p *proc.Process) error {
	quit := false
	go func() {
		for !quit {
			time.Sleep(1 * time.Second)
			if quit {
				return
			}
			written, _ := p.Stdin.Stats()
			_, read := p.Stdout.Stats()
			os.Stderr.WriteString(fmt.Sprintf("Pipe telemetry: written %d bytes -> .pt() -> read %d bytes\n", written, read))
		}
	}()

	io.Copy(p.Stdout, p.Stdin)
	return nil
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/
