//go:build !js
// +build !js

package readline

func print(s string) {
	primary.WriteString(s)
}

func printErr(s string) {
	primary.WriteString(s)
}
