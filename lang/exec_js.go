// +build js

package lang

import (
	"os/exec"
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
	return
}
