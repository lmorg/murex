// +build windows

package proc

// Disable PTY support in Windows.
func ExternalPty(p *Process) error {
	return External(p)
}
