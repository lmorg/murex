//go:build windows
// +build windows

package lang

import "os/exec"

func getCmdTokens(p *Process) (exe string, parameters []string, err error) {
	_, err = p.Parameters.String(0)
	if err != nil {
		return
	}

	exe = "cmd"
	parameters = append([]string{"/c"}, p.Parameters.StringArray()...)

	return
}

func osSyscalls(cmd *exec.Cmd) {
	return
}
