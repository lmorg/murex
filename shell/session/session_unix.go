//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package session

import (
	"os"

	"github.com/lmorg/murex/debug"
	"golang.org/x/sys/unix"
)

var (
	unixSid int
	tty     = os.Stdin
)

func UnixOpenTTY() {
	var err error

	// Opening /dev/tty feels like a bit of a kludge when we already know
	// the tty of stdin. However we often see the following error when
	// attempting to tcsetpgrp the file descriptor of stdin:
	//
	//    inappropriate ioctl for device
	//
	// Where as opening /dev/tty and using that file descriptor resolves
	// that error.
	tty, err = os.Open(`/dev/tty`)
	if err != nil {
		debug.Logf("!!! UnixSetSid()->os.Open(`/dev/tty`) failed: %s", err.Error())
	} else {
		debug.Log("!!! UnixSetSid()->os.Open(`/dev/tty`) success")
	}
}

func UnixIsSession() bool {
	return unixSid > 0
}

func UnixTTY() *os.File {
	return tty
}

func UnixCreateSession() {
	debug.Log("!!! Entering UnixSetSid()")

	var err error
	//pid := os.Getpid()

	// create a new group
	/*err = unix.Setpgid(pid, os.Getpid())
	if err != nil {
		debug.Logf("!!! UnixSetSid()->unix.Setpgid():1 failed: %s", err.Error())
	}*/

	// Create a new session
	unixSid, err = unix.Setsid()
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Setsid():1 failed: %s", err.Error())
	}

	/*pgid, err := unix.Getpgid(pid)
	debug.Logf("pid: %d, ppid: %d, pgid: %d, err: %v", pid, os.Getppid(), pgid, err)
	pgrp := syscall.Getpgrp()
	debug.Logf("pid: %d, ppid: %d, pgid: %d, err: --", pid, os.Getppid(), pgrp)
	sid, err := syscall.Getsid(pid)
	debug.Logf("pid: %d, ppid: %d, sid: %d, err: %v", pid, os.Getppid(), sid, err)*/
}
