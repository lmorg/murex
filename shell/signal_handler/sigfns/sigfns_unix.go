//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package sigfns

import (
	"fmt"
	"syscall"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

func Sigtstp(_ bool) {
	p := lang.ForegroundProc.Get()
	if p.SystemProcess.External() {
		err := p.SystemProcess.Signal(syscall.SIGSTOP)
		if err != nil {
			lang.ShellProcess.Stderr.Write([]byte(err.Error()))
		}
	}

	if p.State.Get() != state.Stopped {
		returnFromSigtstp(p)
	}
}

func returnFromSigtstp(p *lang.Process) {
	p.State.Set(state.Stopped)
	lang.UnixPidToFg(nil)

	show, err := lang.ShellProcess.Config.Get("shell", "stop-status-enabled", types.Boolean)
	if err != nil {
		show = false
	}

	if show.(bool) {
		stopStatus(p)
	}

	lang.ShowPrompt <- true

	p.HasStopped <- true
}

func Sigchld(interactive bool) {
	if !interactive {
		return
	}

	p := lang.ForegroundProc.Get()
	if p.Id == lang.ShellProcess.Id {
		// Child already exited so we can ignore this signal
		return
	}

	var (
		status syscall.WaitStatus
		pid    = p.SystemProcess.Pid()
	)

	if pid == -1 {
		return
	}

	wpid, err := syscall.Wait4(pid, &status, syscall.WNOHANG|syscall.WUNTRACED, nil)
	if err != nil {
		debug.Logf("!!! error in syscall.Wait4(pid: %d, &status, syscall.WNOHANG|syscall.WUNTRACED, nil):\n!!! %v",
			pid, err)
	}

	if wpid == 0 {
		return
	}

	switch {
	case status.Stopped():
		returnFromSigtstp(p)
	}

	/*if p.SystemProcess.State() == nil {
		sid, err := unix.Getsid(p.SystemProcess.Pid())
		if err != nil {
			debug.Logf("!!! Sigchld()->unix.Getsid(p.SystemProcess.Pid: %d) failed: %s", p.SystemProcess.Pid(), err.Error())
			return
		}

		debug.Logf("!!! unix.Getsid(p.SystemProcess.Pid: %d) == %d", p.SystemProcess.Pid(), sid)
		if sid != p.SystemProcess.Pid() {
			return
		}

		// on macOS, Sigchld seems to get called multiple times when a process
		// is stopped. This is likely a bug in Murex, but the following line
		// code avoids any side effects regardless of the root cause.
		if p.State.Get() == state.Stopped {
			return
		}

		if err = p.SystemProcess.Signal(syscall.Signal(0)); err != nil {
			returnFromSigtstp(p)
		}

		return
	}

	if p.SystemProcess.State().Sys().(syscall.WaitStatus).Exited() {
		return // TODO: eventually we should have a clean up of old PIDs
	}*/
}

func stopStatus(p *lang.Process) {
	var (
		stdinR, stdinW   uint64
		stdoutR, stdoutW uint64
		stderrR, stderrW uint64
	)

	if p.Stdin != nil {
		stdinW, stdinR = p.Stdin.Stats()
	}
	if p.Stdout != nil {
		stdoutW, stdoutR = p.Stdout.Stats()
	}
	if p.Stderr != nil {
		stderrW, stderrR = p.Stderr.Stats()
	}

	lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
		"\n!!! FID %d has been stopped:\n!!! %s %s\n!!! Use `fg %d` / `bg %d` to manage the FID\n!!! ...or `jobs` to see a list of background and suspended functions",
		p.Id,
		p.Name.String(), p.Parameters.StringAll(),
		p.Id, p.Id,
	)))

	lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
		"!!!\n!!! STDIN:  %s written / %s read\n!!! STDOUT: %s written / %s read\n!!! STDERR: %s written / %s read",
		humannumbers.Bytes(stdinW), humannumbers.Bytes(stdinR),
		humannumbers.Bytes(stdoutW), humannumbers.Bytes(stdoutR),
		humannumbers.Bytes(stderrW), humannumbers.Bytes(stderrR),
	)))

	if p.SystemProcess.External() {
		block, fileRef, err := lang.ShellProcess.Config.GetFileRef("shell", "stop-status-func", types.CodeBlock)
		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			return
		}

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN)
		fork.Name.Set("(SIGSTOP)")
		fork.FileRef = fileRef
		fork.Variables.Set(fork.Process, "PID", p.SystemProcess.Pid(), types.Integer)
		_, err = fork.Execute([]rune(block.(string)))

		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		}
	}

}

func Sigcont(_ bool) {
	p := lang.ForegroundProc.Get()
	if p.SystemProcess.External() {
		err := p.SystemProcess.Signal(syscall.SIGCONT)
		if err != nil {
			lang.ShellProcess.Stderr.Write([]byte(err.Error()))
		}
	}
}
