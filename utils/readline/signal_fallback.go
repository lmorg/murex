//go:build windows || js || plan9
// +build windows js plan9

package readline

func (rl *Instance) sigwinch() {
	rl.closeSigwinch = func() {
		// empty function because SIGWINCH isn't supported on these platforms
	}
}
