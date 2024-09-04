//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/shell/session"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/which"
	"golang.org/x/sys/unix"
)

func osExecFork(p *Process, argv []string) error {
	if !session.UnixIsSession() {
		debug.Logf("!!! session not defined, falling back to non-unix ttys")
		return execForkFallback(p, argv)
	}

	if p.HasCancelled() {
		return nil
	}
	p.Kill = func() {
		if !debug.Enabled {
			defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		}

		err := p.SystemProcess.Signal(syscall.SIGTERM)
		if err != nil {
			if err.Error() == os.ErrProcessDone.Error() {
				return
			}
			name, _ := p.Args()
			os.Stderr.WriteString(
				fmt.Sprintf("\n!!! Error sending SIGTERM to `%s`: %s\n", name, err.Error()))
		}
	}

	p.State.Set(state.Executing)
	unixProcess, err := os.StartProcess(which.WhichIgnoreFail(argv[0]), argv, unixProcAttr(p.Envs))
	if err != nil {
		return fmt.Errorf("failed fork in os.StartProcess -> osExecFork()...\n%s\nargv: %s",
			err.Error(),
			json.LazyLogging(argv),
		)
	}

	sysProc := sysProcUnixT{p: unixProcess}
	p.SystemProcess.Set(&sysProc)

	UnixPidToFg(sysProc.p.Pid)
	return sysProc.wait()
	/*if err != nil {
		//if !strings.HasPrefix(err.Error(), "signal:") {
		return err
		//}
	}

	return nil*/
}

type sysProcUnixT struct {
	p     *os.Process
	state *os.ProcessState
	mutex sync.Mutex
}

func (sp *sysProcUnixT) Pid() int                   { return sp.p.Pid }
func (sp *sysProcUnixT) ExitNum() int               { return sp.state.ExitCode() }
func (sp *sysProcUnixT) Kill() error                { return sp.p.Kill() }
func (sp *sysProcUnixT) Signal(sig os.Signal) error { return sp.p.Signal(sig) }
func (sp *sysProcUnixT) ForcedTTY() bool            { return true }

func (sp *sysProcUnixT) State() *os.ProcessState {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()
	return sp.state
}

func (sp *sysProcUnixT) wait() error {
	state, err := sp.p.Wait()
	sp.mutex.Lock()
	sp.state = state
	sp.mutex.Unlock()

	/*if state.Sys().(syscall.WaitStatus).Stopped() {
		syscallErr := syscall.Kill(syscall.Getpid(), syscall.SIGTSTP)
		if err != nil {
			return syscallErr
		}
	}*/
	return err
}

// UnixPidToFg brings a UNIX process to the foreground.
// If pid == 0 then UnixPidToFg will assume Murex Pid instead.
func UnixPidToFg(pid int) {
	var err error

	pid, err = syscall.Getpgid(unix.Getpid())
	if err != nil {
		debug.Logf("!!! UnixSetSid()->syscall.Getpgid(unix.Getpid()) failed: %v", err)
		pid = syscall.Getpid()
	}

	err = unixPidToFg(pid, int(os.Stdin.Fd()))
	if err == nil {
		// success, no need to retry
		return
	}

	err = unixPidToFg(pid, int(session.UnixTTY().Fd()))
	if err != nil {
		debug.Logf("!!! UnixPidToFg(%d)->session.UnixTTY(): %s", pid, err.Error())
	}
}

func unixPidToFg(pid int, tty int) error {
	err := unix.IoctlSetPointerInt(tty, unix.TIOCSPGRP, pid)
	if err != nil {
		debug.Log(fmt.Sprintf("!!! unixPidToFg(%d, %d): %s", pid, tty, err.Error()))
	}

	return err
}

/////

func osSysProcAttr(fd int) *syscall.SysProcAttr {
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
		Ctty: fd, // Controlling TTY fd
		// Foreground places the child process group in the foreground.
		// This implies Setpgid. The Ctty field must be set to
		// the descriptor of the controlling TTY.
		// Unlike Setctty, in this case Ctty must be a descriptor
		// number in the parent process.
		//Foreground: true,
		//Pgid:       0, // Child's process group ID if Setpgid.
	}
}

func unixProcAttr(envs []string) *os.ProcAttr {
	return &os.ProcAttr{
		//Files: []*os.File{session.UnixTTY(), session.UnixTTY(), session.UnixTTY()},
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   envs,
		Sys: &syscall.SysProcAttr{
			//Setsid: true, // Create session.
			// Setpgid sets the process group ID of the child to Pgid,
			// or, if Pgid == 0, to the new child's process ID.
			Setpgid: true,
			// Setctty sets the controlling terminal of the child to
			// file descriptor Ctty. Ctty must be a descriptor number
			// in the child process: an index into ProcAttr.Files.
			// This is only meaningful if Setsid is true.
			//Setctty: true,
			//Noctty: true,               // Detach fd 0 from controlling terminal
			//Ctty: 0, // Controlling TTY fd
			// Foreground places the child process group in the foreground.
			// This implies Setpgid. The Ctty field must be set to
			// the descriptor of the controlling TTY.
			// Unlike Setctty, in this case Ctty must be a descriptor
			// number in the parent process.
			//Foreground: true,
			Pgid: 0, // Child's process group ID if Setpgid.
		},
	}
}
