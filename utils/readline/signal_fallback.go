//go:build windows || js
// +build windows js

package readline

func (rl *Instance) sigwinch() {
	rl.closeSigwinch = func() {}
}
