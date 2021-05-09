// +build js

package readline

func print(s string) {
	VTerm.Write([]rune(s))
}

func printErr(s string) {
	VTerm.Write([]rune(s))
}
