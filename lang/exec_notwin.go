//go:build !windows
// +build !windows

package lang

func osExecGetArgv(p *Process) []string {
	return p.Parameters.StringArray()
}
