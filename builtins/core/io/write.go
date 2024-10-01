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
	_WAIT_EOF_SHORT        = "-w"
	_WAIT_EOF_LONG         = "--wait-for-eof"
	_IGNORE_PIPELINE_SHORT = "-i"
	_IGNORE_PIPELINE_LONG  = "--ignore-pipeline-check"
)

func cmdTruncateFile(p *lang.Process) error { return writeFile(p, truncateFile) }
func cmdAppendFile(p *lang.Process) error   { return writeFile(p, appendFile) }

func writeFile(p *lang.Process, fn func(io.Reader, string) error) error {
	p.Stdout.SetDataType(types.Null)

	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if !p.IsMethod {
		// nothing to write
		return fn(bytes.NewBuffer([]byte{}), filename)
	}

	if filename == _IGNORE_PIPELINE_SHORT || filename == _IGNORE_PIPELINE_LONG {
		parameter2, err := p.Parameters.String(1)
		if err == nil {
			return fn(p.Stdin, parameter2)
		}
		// no second parameter so lets assume the flag was actually a file name
	}

	wait := filename == _WAIT_EOF_SHORT || filename == _WAIT_EOF_LONG

	if wait {
		parameter2, err := p.Parameters.String(1)
		if err == nil {
			filename = parameter2

		} else {
			// no second parameter so lets assume the flag was actually a file name
			wait = false
		}
	}

	if !wait {
		wait = isFileOpen(p, filename)
		if wait {
			_, _ = p.Stderr.Writeln([]byte(fmt.Sprintf("warning: '%s' appears as a parameter elsewhere in the pipeline so I'm going to cache the file in RAM before writing to disk.\n       : This message can be suppressed using `%s` or `%s`.", filename, _WAIT_EOF_LONG, _IGNORE_PIPELINE_LONG)))
		}
	}

	if !wait {
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
