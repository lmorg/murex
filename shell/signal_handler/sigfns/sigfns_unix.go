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
	//debug.Log("returnFromSigtstp:0", p.Name.String(), p.Parameters.StringAll())

	p.State.Set(state.Stopped)
	if p.SystemProcess.ForcedTTY() {
		lang.UnixPidToFg(0)
	}

	//debug.Log("returnFromSigtstp:1", p.Name.String(), p.Parameters.StringAll())

	show, err := lang.ShellProcess.Config.Get("shell", "stop-status-enabled", types.Boolean)
	if err != nil {
		show = false
	}

	//debug.Log("returnFromSigtstp:2", p.Name.String(), p.Parameters.StringAll())

	if show.(bool) {
		stopStatus(p)
	}

	//debug.Log("returnFromSigtstp:3", p.Name.String(), p.Parameters.StringAll())

	lang.ShowPrompt <- true

	//debug.Log("returnFromSigtstp:4", p.Name.String(), p.Parameters.StringAll())

	p.HasStopped <- true
}

func Sigchld(interactive bool) {
	//debug.Log("Sigchld:0")
	if !interactive {
		return
	}

	//debug.Log("Sigchld:1")
	p := lang.ForegroundProc.Get()

	//debug.Log("Sigchld:2", p.Name.String(), p.Parameters.StringAll())
	if !p.SystemProcess.ForcedTTY() {
		return
	}

	//debug.Log("Sigchld:3", p.Name.String(), p.Parameters.StringAll())
	if p.SystemProcess.State() == nil || p.SystemProcess.State().Sys().(syscall.WaitStatus).Stopped() {
		//debug.Log("Sigchld:4", p.Name.String(), p.Parameters.StringAll())
		if p.State.Get() != state.Stopped {
			//debug.Log("Sigchld:5", p.Name.String(), p.Parameters.StringAll())
			returnFromSigtstp(p)
		}
	}
	//debug.Log("Sigchld:6", p.Name.String(), p.Parameters.StringAll())
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

	lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
		"!!! FID %d has been stopped:\n!!! %s %s\n!!! Use `fg %d` / `bg %d` to manage the FID\n!!! ...or `jobs` to see a list of background and suspended functions",
		p.Id,
		p.Name.String(), p.Parameters.StringAll(),
		p.Id, p.Id,
	)))
}