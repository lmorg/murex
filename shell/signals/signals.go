package signals

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

const (
	// PromptSIGINT defines the string to write when ctrl+c is pressed
	PromptSIGINT = "^C"

	// PromptSIGQUIT defines the string to write when ctrl+\ is pressed
	PromptSIGQUIT = "^\\"

	// PromptEOF defines the string to write when ctrl+d is pressed
	PromptEOF = "^D"
)

func sigtstp() {
	os.Stderr.WriteString("(murex functions don't currently support being suspended)")

	show, err := proc.ShellProcess.Config.Get("shell", "show-suspend-status", types.Boolean)
	if err != nil {
		show = false
	}
	if !show.(bool) {
		return
	}

	defer func() {
		if debug.Enable {
			return
		}
		if r := recover(); r != nil {
			return
		}
	}()

	stdinR, stdinW := proc.ForegroundProc.Stdin.Stats()
	stdoutR, stdoutW := proc.ForegroundProc.Stdout.Stats()
	stderrR, stderrW := proc.ForegroundProc.Stderr.Stats()
	pipeStatus := fmt.Sprintf(
		"\nSTDIN:  %s/%s bytes r/w\nSTDOUT: %s/%s bytes r/w\nSTDERR: %s/%s bytes r/w\n",
		utils.HumanBytes(stdinR), utils.HumanBytes(stdinW),
		utils.HumanBytes(stdoutR), utils.HumanBytes(stdoutW),
		utils.HumanBytes(stderrR), utils.HumanBytes(stderrW),
	)
	proc.ShellProcess.Stderr.Write([]byte(pipeStatus))

	if proc.ForegroundProc.ExecPid != 0 {
		block, err := proc.ShellProcess.Config.Get("shell", "suspend-status-func", types.CodeBlock)
		if err != nil {
			proc.ShellProcess.Stderr.Writeln([]byte(err.Error()))
			return
		}

		branch := proc.ShellProcess.BranchFID()
		defer branch.Close()
		branch.Process.Variables.Set("PID", proc.ForegroundProc.ExecPid, types.Integer)
		_, err = lang.RunBlockExistingConfigSpace([]rune(block.(string)), nil, proc.ShellProcess.Stdout, proc.ShellProcess.Stderr, branch.Process)
		if err != nil {
			proc.ShellProcess.Stderr.Writeln([]byte(err.Error()))
		}
	}
}

func sigint(interactive bool) {
	//os.Stderr.WriteString(PromptSIGINT)
	sigterm(interactive)
}

func sigterm(interactive bool) {
	if interactive {
		if proc.ForegroundProc != nil && proc.ForegroundProc.Kill != nil {
			proc.ForegroundProc.Kill()
		}

	} else {
		os.Exit(0)
	}
}

func sigquit(interactive bool) {
	if interactive {
		//os.Stderr.WriteString(PromptSIGQUIT)
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)

		fids := proc.GlobalFIDs.ListAll()
		for _, p := range fids {
			if p.Kill != nil {
				procName := p.Name
				procParam, _ := p.Parameters.String(0)
				if p.Name == "exec" {
					procName = procParam
					procParam, _ = p.Parameters.String(1)
				}
				if len(procParam) > 10 {
					procParam = procParam[:10]
				}
				proc.ShellProcess.Stderr.Write([]byte(fmt.Sprintf("!!! Sending kill signal to fid %d: %s %s !!!\n", p.Id, procName, procParam)))
				p.Kill()
			}
		}

	} else {
		os.Stderr.WriteString("Murex received SIGQUIT!" + utils.NewLineString)
		os.Exit(2)
	}
}
