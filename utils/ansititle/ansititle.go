//go:build !js
// +build !js

package ansititle

import (
	"errors"
	"os"

	"github.com/lmorg/murex/utils/readline"
)

func Write(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return errors.New("not a TTY")
	}

	return write(formatTitle(title))
}

func formatTitle(title []byte) []byte {
	if len(title) == 0 {
		return nil
	}
	ansi := make([]byte, len(title)+5)

	copy(ansi[0:4], []byte{27, ']', '2', ';'})
	copy(ansi[4:len(title)+4], title)
	ansi[len(ansi)-1] = 7

	return ansi
}

func Icon(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return errors.New("not a TTY")
	}

	return write(formatIcon(title))
}

func formatIcon(title []byte) []byte {
	if len(title) == 0 {
		return nil
	}
	ansi := make([]byte, len(title)+5)

	copy(ansi[0:4], []byte{27, ']', '1', ';'})
	copy(ansi[4:len(title)+4], title)
	ansi[len(ansi)-1] = 7

	return ansi
}

func Tmux(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return errors.New("not a TTY")
	}
	if value, exists := os.LookupEnv("TMUX"); !exists || value == "" {
		return errors.New("tmux doesn't appear to be running")
	}

	return write(formatTmux(title))
}

func formatTmux(title []byte) []byte {
	if len(title) == 0 {
		return nil
	}
	ansi := make([]byte, len(title)+4)

	copy(ansi[0:2], []byte{27, 'k'})
	copy(ansi[2:len(title)+2], title)
	copy(ansi[len(ansi)-2:], []byte{27, '\\'})

	return ansi
}
