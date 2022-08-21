//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package shell

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/humannumbers"
)

// SignalHandler is an internal function to capture and handle OS signals (eg SIGTERM).
func SignalHandler(interactive bool) {
	signalRegister(interactive)
	
	go func() {
		for {
			sig := <-signalChan
			switch sig.String() {

			case syscall.SIGINT.String():
				sigint(interactive)

			case syscall.SIGTERM.String():
				sigterm(interactive)

			case syscall.SIGQUIT.String():
				sigquit(interactive)

			case syscall.SIGTSTP.String():
				sigtstp()

			default:
				os.Stderr.WriteString("Unhandled signal: " + sig.String())
			}
		}
	}()
}

func signalRegister(interactive bool) {
	if interactive {
		// Interactive, so we will handle stop
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTSTP)
	} else {
		// Non-interactive, so lets ignore the stop signal and let the OS / calling shell manage that for us
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}
}

func sigtstp() {
	p := lang.ForegroundProc.Get()
	//debug.Json("p =", p)

	show, err := lang.ShellProcess.Config.Get("shell", "stop-status-enabled", types.Boolean)
	if err != nil {
		show = false
	}

	if show.(bool) {
		stopStatus(p)
	}

	_, cmd := p.Exec.Get()
	if cmd != nil {
		//err = cmd.Process.Signal(syscall.SIGTSTP)
		err = cmd.Process.Signal(syscall.SIGSTOP)
		if err != nil {
			lang.ShellProcess.Stderr.Write([]byte(err.Error()))
		} else {
			p.State.Set(state.Stopped)
			go ShowPrompt()
		}

	} else {
		lang.ShellProcess.Stderr.Write([]byte("(murex functions don't currently support being stopped)"))
	}
}

func stopStatus(p *lang.Process) {
	//if p == nil {
	//	panic("stopStatus received nil p")
	//}

	var (
		stdinR, stdinW   uint64
		stdoutR, stdoutW uint64
		stderrR, stderrW uint64
	)

	if p.Stdin != nil {
		stdinR, stdinW = p.Stdin.Stats()
	}
	if p.Stdout != nil {
		stdoutR, stdoutW = p.Stdout.Stats()
	}
	if p.Stderr != nil {
		stderrR, stderrW = p.Stderr.Stats()
	}

	pipeStatus := fmt.Sprintf(
		"\nSTDIN:  %s read / %s written\nSTDOUT: %s read / %s written\nSTDERR: %s read / %s written",
		humannumbers.Bytes(stdinR), humannumbers.Bytes(stdinW),
		humannumbers.Bytes(stdoutR), humannumbers.Bytes(stdoutW),
		humannumbers.Bytes(stderrR), humannumbers.Bytes(stderrW),
	)
	lang.ShellProcess.Stderr.Writeln([]byte(pipeStatus))

	if p.Exec.Pid() != 0 {
		block, err := lang.ShellProcess.Config.Get("shell", "stop-status-func", types.CodeBlock)
		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			return
		}

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_NO_STDIN)
		fork.Name.Set("(SIGTSTP)")
		fork.Variables.Set(fork.Process, "PID", lang.ForegroundProc.Get().Exec.Pid(), types.Integer)
		_, err = fork.Execute([]rune(block.(string)))

		if err != nil {
			lang.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		}

		lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
			"FID %d has been stopped. Use `fg %d` / `bg %d` to manage the FID or `jobs` or `fid-list` to see a list of processes running on this shell.",
			p.Id, p.Id, p.Id,
		)))
	}
}
