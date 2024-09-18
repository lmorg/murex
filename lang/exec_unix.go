//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"fmt"
	"os"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/shell/session"
	"golang.org/x/sys/unix"
)

// UnixPidToFg brings a UNIX process to the foreground.
// If p == nil then UnixPidToFg will assume Murex Pid instead.
func UnixPidToFg(p *Process) {
	if !session.UnixIsSession() {
		return
	}

	var (
		pid int
		err error
	)

	if p == nil { // Put Murex in the foreground

		pid, err = unix.Getpgid(unix.Getpid())
		if err != nil {
			//debug.Logf("!!! UnixSetSid()->unix.Getpgid(unix.Getpid()) failed: %v", err)
			pid = unix.Getpid()
		}

		// This is only required because some badly behaving programs run
		// setsid() themselves despite not technically needing to be a session
		// leader eg shell.
		unix.Setsid()

	} else { // Put a system process in the foreground

		// Check if its system process, if not, then there's no point proceeding
		pid = p.SystemProcess.Pid()
		if pid <= 0 {
			return
		}
	}

	err = unixPidToFg(pid, int(os.Stdin.Fd()))
	if err == nil {
		// Success, no need to retry with a different file descriptor
		return
	}

	// fallback is to use /dev/tty. This seems the default recommendation in a
	// lot of the example code and documentation on this topic but it still
	// feels "wrong" not to at least try os.Stdin first.
	_ = unixPidToFg(pid, int(session.UnixTTY().Fd()))
}

func unixPidToFg(pid int, tty int) error {
	err := unix.IoctlSetPointerInt(tty, unix.TIOCSPGRP, pid)
	if err != nil {
		debug.Log(fmt.Sprintf("!!! unixPidToFg(%d, %d): %s", pid, tty, err.Error()))
	}

	return err
}

/////

func unixProcAttrFauxTTY() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		//Setsid: true, // Create session.
		// Setpgid sets the process group ID of the child to Pgid,
		// or, if Pgid == 0, to the new child's process ID.
		//Setpgid: true,
		// Setctty sets the controlling terminal of the child to
		// file descriptor Ctty. Ctty must be a descriptor number
		// in the child process: an index into ProcAttr.Files.
		// This is only meaningful if Setsid is true.
		//Setctty: true,
		//Noctty:  true, // Detach fd 0 from controlling terminal
		//Ctty: 0, // Controlling TTY fd
		// Foreground places the child process group in the foreground.
		// This implies Setpgid. The Ctty field must be set to
		// the descriptor of the controlling TTY.
		// Unlike Setctty, in this case Ctty must be a descriptor
		// number in the parent process.
		//Foreground: true,
		//Pgid: 0, // Child's process group ID if Setpgid.
	}
}
