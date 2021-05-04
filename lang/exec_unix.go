// +build !windows,!plan9,!js

package lang

import (
	"os"
	"os/exec"
	"syscall"
)

func getCmdTokens(p *Process) (exe string, parameters []string, err error) {
	exe, err = p.Parameters.String(0)
	if err != nil {
		return
	}

	parameters = p.Parameters.StringArray()[1:]

	return
}

func osSyscalls(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ctty: int(os.Stdout.Fd()),
	}
}
