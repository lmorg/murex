// +build windows

package proc

// Disable PTY support in Windows.
func ExternalPty(p *Process) error {
	// Nasty fudge to work around the fact that a lot of Windows's common commands are built into cmd.exe
	p.Parameters.SetPrepend("/c")
	p.Parameters.SetPrepend("cmd")
	return External(p)
}
