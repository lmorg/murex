//go:build windows
// +build windows

package ansititle

import (
	"errors"
	"os"
	"syscall"
	"unsafe"

	"github.com/lmorg/readline/v4"
)

var ErrNotTTY = errors.New("not a TTY")

func Write(title []byte) error {
	fd := os.Stdout.Fd()
	if !readline.IsTerminal(int(fd)) {
		return ErrNotTTY
	}

	title = sanitize(title)

	ptr, err := syscall.UTF16PtrFromString(string(title))
	if err != nil {
		return err
	}

	ret, _, err := syscall.
		NewLazyDLL("kernel32.dll").
		NewProc("SetConsoleTitleW").
		Call(uintptr(unsafe.Pointer(ptr)))

	if ret == 0 {
		return err
	}

	return nil
}

func Icon(title []byte) error { return nil }
func Tmux(title []byte) error { return errors.New("tmux not supported on Windows") }

func sanitize(b []byte) []byte {
	newb := make([]byte, len(b))
	copy(newb, b)

	for i := range newb {
		if newb[i] == '\r' {
			newb[i] = ' '
		}
		if newb[i] < 32 || newb[i] == 127 {
			newb[i] = 32
		}
	}

	return newb
}
