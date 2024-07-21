//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
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

func osSyscalls(cmd *exec.Cmd, fd int) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ctty: fd,
		//Noctty:  false,
		//Setctty: true,
		//Setsid:  true,
	}
}
