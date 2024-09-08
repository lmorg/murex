//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package sigfns

import (
	"fmt"
	"os"
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
	//debug.Log("returnFromSigtstp:0", p.Name.String(), p.Parameters.StringAll())

	p.State.Set(state.Stopped)
	//if p.SystemProcess.ForcedTTY() {
	lang.UnixPidToFg(nil)
	//}

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
	debug.Logf("!!! Sigchld(fid: %d)->session.UnixIsSession(fid: %d, pid: %d): %s %s", lang.ShellProcess.Id, p.Id, os.Getpid()) //, p, p.Parameters.StringAll())
	if p.Id == lang.ShellProcess.Id {
		// Child already exited so we can ignore this signal
		return
	}

	/*debug.Logf("!!! Sigchld()->session.UnixIsSession(pid: %d) == %v", os.Getpid(), session.UnixIsSession())
	if !session.UnixIsSession() {
		return
	}*/

	/*Logf("session.UnixCompareSid() == %v", session.UnixCompareSid())
	if !session.UnixCompareSid() {
		return
	}*/

	debug.Logf("!!! Sigchld()->p.SystemProcess.State(pid: %d) == %v", p.SystemProcess.Pid(), p.SystemProcess.State())
	if p.SystemProcess.State() == nil {
		sid, err := syscall.Getsid(p.SystemProcess.Pid())
		if err != nil {
			debug.Logf("!!! Sigchld()->syscall.Getsid(p.SystemProcess.Pid: %d) failed: %s", p.SystemProcess.Pid(), err.Error())
			return
		}
		/*pgid, err := syscall.Getpgid(p.SystemProcess.Pid())
		if err != nil {
			debug.Logf("!!! Sigchld()->syscall.Getpgid(p.SystemProcess.Pid: %d) failed: %s", p.SystemProcess.Pid(), err.Error())
		}*/

		debug.Logf("!!! syscall.Getsid(p.SystemProcess.Pid: %d) == %d", p.SystemProcess.Pid(), sid)
		if sid != p.SystemProcess.Pid() {
			return
		}

		if p.State.Get() == state.Stopped {
			return
		}

		debug.Log("!!! calling returnFromSigtstp(p)")
		returnFromSigtstp(p)
		debug.Log("!!! returned from returnFromSigtstp(p)")
		/*err = p.SystemProcess.Signal(syscall.SIGSTOP)
		if err != nil {
			debug.Logf("!!! Sigchld()->p.SystemProcess.Signal(pid: %d, syscall.SIGSTOP) failed: %s", p.SystemProcess.Pid(), err.Error())
		}*/
		return
	}

	if p.SystemProcess.State().Sys().(syscall.WaitStatus).Exited() {
		debug.Logf("!!! Sigchld()->p.SystemProcess.State(pid: %d).Sys().(syscall.WaitStatus).Exited() == true", os.Getpid())
		return
	}

	//if p.SystemProcess

	//if p.SystemProcess.Pid()

	//debug.Log("Sigchld:3", p.Name.String(), p.Parameters.StringAll())
	/*if p.SystemProcess.State() == nil || p.SystemProcess.State().Sys().(syscall.WaitStatus).Stopped() {
		//debug.Log("Sigchld:4", p.Name.String(), p.Parameters.StringAll())
		if p.State.Get() != state.Stopped {
			//debug.Log("Sigchld:5", p.Name.String(), p.Parameters.StringAll())
			//returnFromSigtstp(p)
			proc, err := os.FindProcess(os.Getppid())
			if err != nil {
				debug.Logf("!!! Sigchld()->os.FindProcess(os.Getppid()) failed: %s", err.Error())
				return
			}
			debug.Logf("!!! Sigchld()->proc.Signal(Pppid: %d, syscall.SIGTSTP) invoked", proc.Pid)
			err = proc.Signal(syscall.SIGTSTP)
			if err != nil {
				debug.Logf("!!! Sigchld()->proc.Signal(syscall.SIGTSTP) failed: %s", err.Error())
			}
		}
	}*/
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
