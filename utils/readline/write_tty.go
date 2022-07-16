//go:build !js
// +build !js

package readline

import "os"

func print(s string) {
	os.Stdout.WriteString(s)
}

func printErr(s string) {
	os.Stderr.WriteString(s)
}
