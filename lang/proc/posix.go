// +build !windows

package proc

import (
	"github.com/kr/pty"
	"github.com/lmorg/murex/lang/proc/streams/osstdin"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func External(p *Process) error {
	// External executable
	if err := execute(p); err != nil {
		// Get exit status. This has only been tested on Linux. May not work on other OSs.
		if strings.HasPrefix(err.Error(), "exit status ") {
			i, _ := strconv.Atoi(strings.Replace(err.Error(), "exit status ", "", 1))
			p.ExitNum = i
		} else {
			p.Stderr.Writeln([]byte(err.Error()))
			p.ExitNum = 1
		}

	}
	return nil
}

func execute(p *Process) error {
	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	parameters := p.Parameters.StringArray()
	cmd := exec.Command(exeName, parameters[1:]...)

	cmd.Stdin = p.Stdin
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// Prototype call with support for PTYs. Highly experimental.
func ExternalPty(p *Process) error {
	// External executable
	if err := shellExecute(p); err != nil {
		// Get exit status. This has only been tested on Linux. May not work on other OSs.
		if strings.HasPrefix(err.Error(), "exit status ") {
			i, _ := strconv.Atoi(strings.Replace(err.Error(), "exit status ", "", 1))
			p.ExitNum = i
		} else {
			p.Stderr.Writeln([]byte(err.Error()))
			p.ExitNum = 1
		}

	}
	return nil
}

// Prototype call with support for PTYs. Highly experimental.
func shellExecute(p *Process) (err error) {
	// Create an object for the executable we wish to invoke.
	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	parameters := p.Parameters.StringArray()
	cmd := exec.Command(exeName, parameters[1:]...)

	// Create a PTY for the executable.
	f, err := pty.Start(cmd)
	if err != nil {
		return err
	}

	// Create an STDIN function, copying 1KB blocks at a time.
	active := true
	go func() {
		b := make([]byte, 1024)
		for active {
			var i int

			i, err := osstdin.Stdin.Read(b)
			if err != nil {
				return
			}
			// oops the program has closed but this goroutine is still active.
			// So lets push the []bytes back into the stack.
			if !active {
				osstdin.Stdin.Prepend(b[:i])
				return
			}
			if _, err = f.Write(b[:i]); err != nil {
				return
			}
		}
	}()

	io.Copy(p.Stdout, f)
	active = false
	return
}
