//go:build windows
// +build windows

package ansititle

func Write(title []byte) error { return nil }
func Icon(title []byte) error  { return nil }
func Tmux(title []byte) error  { return nil }
