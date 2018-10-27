package proc

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang/types"
)

// External executes an external process.
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
	p.Stdout.SetDataType(types.Generic)

	exeName, parameters, err := getCmdTokens(p)
	if err != nil {
		return err
	}
	cmd := exec.Command(exeName, parameters...)

	p.Kill = func() {
		defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		cmd.Process.Kill()
	}
	KillForeground = p.Kill

	if p.IsMethod {
		cmd.Stdin = p.Stdin
	} else {
		cmd.Stdin = os.Stdin
	}

	if p.Stdout.IsTTY() {
		// If Stdout is a TTY then set the appropiate syscalls to allow the calling program to own the TTY....
		osSyscalls(cmd)
		cmd.Stdout = os.Stdout
	} else {
		// ....othwise we just treat the program as a regular piped util
		cmd.Stdout = p.Stdout
	}

	if p.Stderr.IsTTY() {
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stderr = p.Stderr
	}

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

	//debug.Log("exec env:",cmd.Env)

	return nil
}
