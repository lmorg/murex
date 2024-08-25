//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package sigfns

import (
	"fmt"
	"syscall"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

func Sigtstp(interactive bool) {
	p := lang.ForegroundProc.Get()
	if p.SystemProcess != nil {
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
	if p.SystemProcess != nil && p.SystemProcess.ForcedTTY() {
		lang.UnixPidToFg(0)
	}

	show, err := lang.ShellProcess.Config.Get("shell", "stop-status-enabled", types.Boolean)
	if err != nil {
		show = false
	}

	if show.(bool) {
		stopStatus(p)
	}

	p.State.Set(state.Stopped)
	go func() { p.HasStopped <- true }()

	lang.ShowPrompt <- true
}

func Sigchld(interactive bool) {
	if !interactive {
		return
	}

	p := lang.ForegroundProc.Get()
	if p.SystemProcess == nil {
		return
	}

	if !p.SystemProcess.ForcedTTY() {
		return
	}

	if p.SystemProcess.State() == nil || p.SystemProcess.State().Sys().(syscall.WaitStatus).Stopped() {
		if p.State.Get() != state.Stopped {
			returnFromSigtstp(p)
		}
	}
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

	pipeStatus := fmt.Sprintf(
		"\n!!! STDIN:  %s written / %s read\n!!! STDOUT: %s written / %s read\n!!! STDERR: %s written / %s read",
		humannumbers.Bytes(stdinW), humannumbers.Bytes(stdinR),
		humannumbers.Bytes(stdoutW), humannumbers.Bytes(stdoutR),
		humannumbers.Bytes(stderrW), humannumbers.Bytes(stderrR),
	)
	lang.ShellProcess.Stderr.Writeln([]byte(pipeStatus))

	if p.SystemProcess != nil {
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

	lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
		"!!! FID %d has been stopped:\n!!! %s %s\n!!! Use `fg %d` / `bg %d` to manage the FID\n!!! ...or `jobs` to see a list of background and suspended functions",
		p.Id,
		p.Name.String(), p.Parameters.StringAll(),
		p.Id, p.Id,
	)))
}
