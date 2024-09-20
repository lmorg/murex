package io

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/lists"
)

func init() {
	lang.DefineMethod(">", cmdTruncateFile, types.Any, types.Null)
	lang.DefineMethod("fwrite", cmdTruncateFile, types.Any, types.Null)
	lang.DefineMethod(">>", cmdAppendFile, types.Any, types.Null)
	lang.DefineMethod("fappend", cmdAppendFile, types.Any, types.Null)
}

const (
	_WAIT_EOF_LONG       = "--wait-for-eof"
	_WAIT_EOF_SHORT      = "-w"
	_DONT_CHECK_PIPELINE = "--ignore-pipeline-check"
)

func cmdTruncateFile(p *lang.Process) error { return writeFile(p, truncateFile) }
func cmdAppendFile(p *lang.Process) error   { return writeFile(p, appendFile) }

func writeFile(p *lang.Process, fn func(io.Reader, string) error) error {
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
		return fn(p.Stdin, filename)
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
		return fn(p.Stdin, filename)
	}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	return fn(bytes.NewReader(b), filename)
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

func truncateFile(reader io.Reader, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func appendFile(reader io.Reader, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}
