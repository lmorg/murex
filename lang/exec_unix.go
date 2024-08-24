//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package lang

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/which"
)

func osExecGetArgv(p *Process) []string {
	argv := p.Parameters.StringArray()
	argv[0] = which.WhichIgnoreFail(argv[0])
	return argv
}

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

	pwd, _ := os.Getwd()
	p.State.Set(state.Executing)
	/*pid, err := syscall.ForkExec(argv[0], argv, &syscall.ProcAttr{
		Dir:   pwd,
		Files: []uintptr{stdinFd(p), stdoutFd(p), stderrFd(p)},
		//Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
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
			//Noctty:  true,               // Detach fd 0 from controlling terminal
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
		return fmt.Errorf("failed syscall in osExecFork(): %s\nargv: %s",
			err.Error(),
			json.LazyLogging(argv),
		)
	}*/

	unixProcess, err := os.StartProcess(argv[0], argv, &os.ProcAttr{
		Dir:   pwd,
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   p.Envs,
		Sys: &syscall.SysProcAttr{
			//Setsid: true, // Create session.
			// Setpgid sets the process group ID of the child to Pgid,
			// or, if Pgid == 0, to the new child's process ID.
			Setpgid: true,
			// Setctty sets the controlling terminal of the child to
			// file descriptor Ctty. Ctty must be a descriptor number
			// in the child process: an index into ProcAttr.Files.
			// This is only meaningful if Setsid is true.
			Setctty: true,
			//Noctty:  true,               // Detach fd 0 from controlling terminal
			Ctty: int(os.Stdin.Fd()), // Controlling TTY fd
			// Foreground places the child process group in the foreground.
			// This implies Setpgid. The Ctty field must be set to
			// the descriptor of the controlling TTY.
			// Unlike Setctty, in this case Ctty must be a descriptor
			// number in the parent process.
			//Foreground: true,
			Pgid: 0, // Child's process group ID if Setpgid.
		},
	})

	if err != nil {
		return fmt.Errorf("failed os.StartProcess in osExecFork()...\n%s\nargv: %s",
			err.Error(),
			json.LazyLogging(argv),
		)
	}

	sysProc := sysProcUnixT{
		p: unixProcess,
	}

	p.SystemProcess = &sysProc

	sysProc.state, err = sysProc.p.Wait()

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

func stdinFd(_ *Process) uintptr {
	return os.Stdin.Fd()
}

func stdoutFd(p *Process) uintptr {
	f := p.Stdout.File()
	if f != nil {
		return f.Fd()
	}

	//path := fmt.Sprintf("%s%d.1")
	panic("not implemented file creation")
}

func stderrFd(p *Process) uintptr {
	f := p.Stderr.File()
	if f != nil {
		return f.Fd()
	}

	//path := fmt.Sprintf("%s%d.2")
	panic("not implemented file creation")
}

type sysProcUnixT struct {
	p     *os.Process
	state *os.ProcessState
}

func (sp *sysProcUnixT) Pid() int {
	panic("bob00")
	return sp.p.Pid
}
func (sp *sysProcUnixT) ExitNum() int {
	panic("bob01")
	return sp.state.ExitCode()
}
func (sp *sysProcUnixT) Kill() error {
	panic("bob02")
	return sp.p.Kill()
}
func (sp *sysProcUnixT) Signal(sig os.Signal) error {
	panic("bob03")
	return sp.p.Signal(sig)
}

func (sp *sysProcUnixT) UnixT() bool {
	return true
}
