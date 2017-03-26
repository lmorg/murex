package proc

import (
	"github.com/kr/pty"
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
	cmd := exec.Command(p.Parameters[0], p.Parameters[1:]...)

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

// Prototype call with support for PTYs. Highly experimental, doesn't really work yet.
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

// Prototype call with support for PTYs. Highly experimental, doesn't really work yet.
func shellExecute(p *Process) error {
	cmd := exec.Command(p.Parameters[0], p.Parameters[1:]...)

	//cmd.Stdin = p.Stdin
	//cmd.Stdout = p.Stdout
	//cmd.Stderr = p.Stderr

	f, err := pty.Start(cmd)
	if err != nil {
		return err
	}

	go io.Copy(f, p.Stdin)
	io.Copy(p.Stdout, f)

	return nil
}
