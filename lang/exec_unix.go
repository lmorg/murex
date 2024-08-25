//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/state"
	signalhandler "github.com/lmorg/murex/shell/signal_handler"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/which"
	"golang.org/x/sys/unix"
)

func osExecFork(p *Process, argv []string) error {
	if p.HasCancelled() {
		return nil
	}
	p.Kill = func() {
		if !debug.Enabled {
			defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		}

		//ctxCancel()
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

	unixProcess, err := os.StartProcess(which.WhichIgnoreFail(argv[0]), argv, &os.ProcAttr{
		//Dir:   pwd,
		Files: []*os.File{
			os.Stdin,
			//os.Stdout,
			//os.Stderr,
			//p.Stdin.File(),
			p.Stdout.File(),
			p.Stderr.File(),
		},
		Env: p.Envs,
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
			Ctty: int(os.Stdin.Fd()), // Controlling TTY fd
			// Foreground places the child process group in the foreground.
			// This implies Setpgid. The Ctty field must be set to
			// the descriptor of the controlling TTY.
			// Unlike Setctty, in this case Ctty must be a descriptor
			// number in the parent process.
			Foreground: true,
			Pgid:       0, // Child's process group ID if Setpgid.
		},
	})

	if err != nil {
		return fmt.Errorf("failed fork in os.StartProcess -> osExecFork()...\n%s\nargv: %s",
			err.Error(),
			json.LazyLogging(argv),
		)
	}

	sysProc := sysProcUnixT{
		p: unixProcess,
	}

	p.SystemProcess = &sysProc

	signal.Ignore(syscall.SIGTTOU)
	defer signal.Reset(syscall.SIGTTOU)
	UnixPidToFg(sysProc.p.Pid)
	signalhandler.Register(Interactive)
	sysProc.state, err = sysProc.p.Wait()
	UnixPidToFg(0)

	if err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	return nil
}

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
		//Pgid: 0, // Child's process group ID if Setpgid.
	}
}

type sysProcUnixT struct {
	p     *os.Process
	state *os.ProcessState
}

func (sp *sysProcUnixT) Pid() int                   { return sp.p.Pid }
func (sp *sysProcUnixT) ExitNum() int               { return sp.state.ExitCode() }
func (sp *sysProcUnixT) Kill() error                { return sp.p.Kill() }
func (sp *sysProcUnixT) Signal(sig os.Signal) error { return sp.p.Signal(sig) }

// UnixPidToFg brings a UNIX process to the foreground.
// If pid == 0 then UnixPidToFg will assume Murex Pid instead.
func UnixPidToFg(pid int) {
	if pid == 0 {
		pid = os.Getpid()
	}

	err := unix.IoctlSetPointerInt(0, unix.TIOCSPGRP, pid)
	if err != nil {
		os.Stderr.WriteString(
			fmt.Sprintf("!!! failed syscall in unix.IoctlSetPointerInt -> unixResetFg()...\n!!! %s\n",
				err.Error(),
			),
		)
	}
}
