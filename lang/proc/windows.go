// +build windows

package proc

// Disable PTY support in Windows.
func ExternalPty(p *Process) error {
	// Nasty fudge to work around the fact that a lot of Windows's common commands are built into cmd.exe
	p.Parameters.Params = append([]string{"cmd", "/c"}, p.Parameters.Params...)
	return External(p)
}
