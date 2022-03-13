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

func write(ansi []byte) error {
	if ansi == nil {
		return nil
	}

	_, err := os.Stdout.Write(ansi)
	return err
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
