// +build !windows

package proc

import (
	"github.com/lmorg/murex/lang/types"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Execute an external process. Don't give it a TTY
func External(p *Process) error {
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
	p.Stdout.SetDataType(types.String)

	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	parameters := p.Parameters.StringArray()
	cmd := exec.Command(exeName, parameters[1:]...)

	p.Kill = func() {
		defer func() { recover() }() // I don't care about errors.
		cmd.Process.Kill()
	}
	KillForeground = p.Kill

	cmd.Stdin = p.Stdin
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stderr

	if err := cmd.Start(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	return nil
}

// Execute an external process. Give it a TTY
func ExternalPty(p *Process) error {
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

func shellExecute(p *Process) error {
	p.Stdout.SetDataType(types.String)

	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	parameters := p.Parameters.StringArray()
	cmd := exec.Command(exeName, parameters[1:]...)

	cmd.SysProcAttr = &syscall.SysProcAttr{Ctty: int(os.Stdout.Fd())}

	p.Kill = func() {
		defer func() { recover() }() // I don't care about errors.
		cmd.Process.Kill()
	}
	KillForeground = p.Kill

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	return nil
}
