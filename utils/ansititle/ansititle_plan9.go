//go:build plan9
// +build plan9

package ansititle

func Write(title []byte) error { return nil }
func Icon(title []byte) error  { return nil }
func Tmux(title []byte) error  { return nil }
