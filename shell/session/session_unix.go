//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package session

import (
	"os"
	"syscall"

	"github.com/lmorg/murex/debug"
	signalhandler "github.com/lmorg/murex/shell/signal_handler"
)

var (
	unixSid int
	tty     = os.Stdin
)

func UnixSetSid() {
	var err error
	debug.Enabled = true
	debug.Log("!!! Entering UnixSetSid()")

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


	pid := os.Getpid()

	//sid, _ := syscall.Getsid(pid)
	//if syscall.Getpgrp() == sid {
	//	return
	//}

	// create a new group
	err = syscall.Setpgid(pid, os.Getppid())
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Setpgid():1 failed: %s", err.Error())
	}

	// Create a new session
	unixSid, err = syscall.Setsid()
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Setsid():1 failed: %s", err.Error())
	}

	// create a new group
	err = syscall.Setpgid(pid, pid)
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Setpgid():2 failed: %s", err.Error())
	}

	// Create a new session
	unixSid, err = syscall.Setsid()
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Setsid():2 failed: %s", err.Error())
	}

	signalhandler.Register(true)

	pgid, err := syscall.Getpgid(pid)
	debug.Logf("pid: %d, ppid: %d, pgid: %d, err: %v", pid, os.Getppid(), pgid, err)
	pgrp := syscall.Getpgrp()
	debug.Logf("pid: %d, ppid: %d, pgid: %d, err: --", pid, os.Getppid(), pgrp)
	sid, err := syscall.Getsid(pid)
	debug.Logf("pid: %d, ppid: %d, sid: %d, err: %v", pid, os.Getppid(), sid, err)

}

func UnixIsSession() bool {
	return unixSid > 0
}

func UnixTTY() *os.File {
	return tty
}

/*
func relaunchMurex() error {
	if os.Getenv("MUREX_SESSION") != "" {
		return fmt.Errorf("session already nested")
	}

	cmd := exec.Command(which.WhichIgnoreFail(os.Args[0]), os.Args[1:]...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("MUREX_SESSION=%d", os.Getpid()))
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	size, err := pty.GetsizeFull(UnixTTY())
	if err != nil {
		return fmt.Errorf("cannot get size of terminal: %s", err.Error())
	}

	tty, err := pty.StartWithSize(cmd, size)
	if err != nil {
		return fmt.Errorf("error starting process: %s", err.Error())
	}

	mxState, err := readline.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("cannot put TTY into raw mode: %s", err.Error())
	}

	defer readline.Restore(int(os.Stdin.Fd()), mxState)

	go io.Copy(os.Stdout, tty)
	go io.Copy(tty, os.Stdin)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	os.Exit(0)
	return nil // this is silly but go doesn't compile without it
}
*/
