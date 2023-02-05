//go:build !js
// +build !js

package readline

import "os"

func print(s string) {
	term.WriteString(s)
}

func printErr(s string) {
	term.WriteString(s)
}

func (rl *Instance) SetTTY(tty *os.File) {
	term = tty
}
