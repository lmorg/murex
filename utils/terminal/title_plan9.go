//go:build plan9
// +build plan9

package terminal

func Title(title []byte) error     { return nil }
func Icon(title []byte) error      { return nil }
func TitleTmux(title []byte) error { return nil }
