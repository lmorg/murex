//go:build !js
// +build !js

package readline

func print(s string) {
	term.WriteString(s)
}

func printErr(s string) {
	term.WriteString(s)
}
