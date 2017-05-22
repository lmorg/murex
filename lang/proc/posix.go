// +build !windows

package proc

import (
	"github.com/kr/pty"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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
func shellExecute(p *Process) error {
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
		b := make([]byte, 1024*1024)
		for active {
			i, err := os.Stdin.Read(b)
			if err != nil {
				return
			}
			if _, err := f.Write(b[:i]); err != nil {
				return
			}
		}
	}()

	//go io.Copy(f, p.Stdin)
	io.Copy(p.Stdout, f)
	active = false
	return nil
}
