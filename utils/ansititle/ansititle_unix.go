//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package ansititle

import (
	"bytes"
	"errors"
	"os"

	"github.com/lmorg/murex/utils/readline"
)

var ErrNotTTY = errors.New("not a TTY")

func Write(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return ErrNotTTY
	}

	return write(formatTitle(title))
}

func formatTitle(title []byte) []byte {
	if len(title) == 0 {
		return nil
	}
	title = sanatise(title)
	ansi := make([]byte, len(title)+6)

	copy(ansi[0:4], []byte{27, ']', '2', ';'})
	copy(ansi[4:len(title)+4], title)
	//ansi[len(ansi)-1] = 7
	ansi[len(ansi)-2] = 'S'
	ansi[len(ansi)-1] = 'T'

	return ansi
}

func Icon(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return ErrNotTTY
	}

	return write(formatIcon(title))
}

func formatIcon(title []byte) []byte {
	if len(title) == 0 {
		return nil
	}
	title = sanatise(title)
	ansi := make([]byte, len(title)+6)

	copy(ansi[0:4], []byte{27, ']', '1', ';'})
	copy(ansi[4:len(title)+4], title)
	//ansi[len(ansi)-1] = 7
	ansi[len(ansi)-2] = 'S'
	ansi[len(ansi)-1] = 'T'

	return ansi
}

func Tmux(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return ErrNotTTY
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
	title = sanatise(title)
	ansi := make([]byte, len(title)+4)

	copy(ansi[0:2], []byte{27, 'k'})
	copy(ansi[2:len(title)+2], title)
	copy(ansi[len(ansi)-2:], []byte{27, '\\'})

	return ansi
}

func sanatise(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte{'\r'}, nil)
	// replace all control characters with space
	for i := range b {
		if b[i] < 32 || b[i] == 127 {
			b[i] = 32
		}
	}

	return b
}
