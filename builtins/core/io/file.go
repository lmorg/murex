package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
	"github.com/lmorg/murex/utils/lists"
)

func init() {
	lang.DefineMethod("pt", cmdPipeTelemetry, types.Any, types.Any)
	lang.DefineMethod(">", cmdWriteFile, types.Any, types.Null)
	lang.DefineMethod("fwrite", cmdWriteFile, types.Any, types.Null)
	lang.DefineMethod(">>", cmdAppendFile, types.Any, types.Null)
	lang.DefineMethod("fappend", cmdAppendFile, types.Any, types.Null)
}

func cmdPipeTelemetry(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)
	quit := make(chan bool)
	stats := func() {
		written, _ := p.Stdin.Stats()
		_, read := p.Stdout.Stats()
		os.Stderr.WriteString(
			fmt.Sprintf("Pipe telemetry: `%s` written %s -> pt -> `%s` read %s (Data type: %s)\n",
				p.Previous.Name.String(),
				humannumbers.Bytes(written),
				p.Next.Name.String(),
				humannumbers.Bytes(read),
				dt),
		)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-quit:
				return
			default:
				stats()
			}

		}
	}()

	_, err := io.Copy(p.Stdout, p.Stdin)
	quit <- true
	stats()
	return err
}

const (
	_WAIT_EOF_LONG       = "--wait-for-eof"
	_WAIT_EOF_SHORT      = "-w"
	_DONT_CHECK_PIPELINE = "--ignore-pipeline-check"
)

func cmdWriteFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if filename == _DONT_CHECK_PIPELINE {
		filename, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
		return writeFile(p.Stdin, filename)
	}

	wait := filename == _WAIT_EOF_SHORT || filename == _WAIT_EOF_LONG

	if wait {
		filename, err = p.Parameters.String(1)
		if err != nil {
			return err
		}
	} else {
		wait = isFileOpen(p, filename)
	}

	if wait {
		_, _ = p.Stderr.Writeln([]byte(fmt.Sprintf("warning: '%s' appears as a parameter elsewhere in the pipeline so I'm going to cache the file in RAM before writing to disk.\n       : this message can be suppressed using `%s` or `%s`.", filename, _WAIT_EOF_LONG, _DONT_CHECK_PIPELINE)))
	} else {
		return writeFile(p.Stdin, filename)
	}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	return writeFile(bytes.NewReader(b), filename)
}

func isFileOpen(p *lang.Process, filename string) bool {
	p = p.Previous
	for {
		if p.State.Get() < state.Executing {
			continue
		}
		if lists.Match(p.Parameters.StringArray(), filename) {
			return true
		}
		if !p.IsMethod || p.Id == lang.ShellProcess.Id {
			return false
		}
		p = p.Previous
	}
}

func writeFile(reader io.Reader, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func cmdAppendFile(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, p.Stdin)
	return err
}
