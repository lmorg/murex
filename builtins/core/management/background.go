package management

import (
	"errors"
	"syscall"

	"github.com/lmorg/murex/lang/proc/state"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
)

func init() {
	proc.GoFunctions["bg"] = cmdBackground
	proc.GoFunctions["fg"] = cmdForeground
}

func cmdBackground(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	var block []rune

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

	} else {
		block, err = p.Parameters.Block(0)
		if err != nil {
			return err
		}
	}

	p.IsBackground = true
	p.WaitForTermination <- false
	lang.RunBlockExistingConfigSpace(block, p.Stdin, p.Stdout, p.Stderr, p)

	return nil
}

func cmdForeground(p *proc.Process) error {
	fid, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	f, err := proc.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.ExecPid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `fg`.")
	}

	if f.ExecCmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process.")
	}

	err = f.ExecCmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	f.State = state.Executing
	f.IsBackground = false
	shell.PromptGoProc.Set(f.PromptGoProc)
	return nil
}
