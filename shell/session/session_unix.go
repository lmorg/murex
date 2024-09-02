//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package session

import (
	"fmt"
	"syscall"

	"github.com/lmorg/murex/debug"
)

var unixSid int

func UnixSetSid() {
	var err error

	// Create a new session
	unixSid, err = syscall.Setsid()
	if err == nil {
		return
	}

	debug.Log(fmt.Sprintf("!!! syscall.Setsid() failed: %s", err.Error()))

}

func UnixIsSession() bool {
	return unixSid > 0
}

/*
	p, t, err := pty.Open()
	if err != nil {
		debug.Log(fmt.Sprintf(`!!! pty.Open() failed: %s`, err.Error()))
		return
	}

	f, err := os.Open("/dev/tty")
	if err != nil {
		debug.Log(fmt.Sprintf(`!!! os.Open("/dev/tty") failed: %s`, err.Error()))
		return
	}

	os.Stdin = p
	os.Stderr = p
	os.Stdout = p

	go io.Copy(os.Stdout, t)
	go io.Copy(os.Stdin, f)

	_, err = syscall.Setsid()
	if err == nil {
		debug.Log(fmt.Sprintf("!!! syscall.Setsid() failed again: %s", err.Error()))
	}
*/
