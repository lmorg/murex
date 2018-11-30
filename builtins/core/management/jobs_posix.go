// +build !windows

package management

import (
	"errors"
	"syscall"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
)

func mkbg(p *proc.Process) error {
	fid, err := p.Parameters.Int(0)
	if err != nil {
		return errors.New("Invalid parameters. Expecting either a code block or FID of a suspended process.")
	}

	f, err := proc.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.State != state.Suspended {
		return errors.New("FID is not a suspended process. Run `jobs` or `fid-list` to see a list of suspended processes.")
	}

	if f.Exec.Pid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `bg`.")
	}

	if f.Exec.Cmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process.")
	}

	updateTree(f, true)

	if !f.IsMethod {
		// This doesn't work. But we would need something clever like this:
		//f.Exec.Cmd.Stdin = streams.NewStdin()
	}

	err = f.Exec.Cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	// This doesn't belong here. It should only be called if the requesting program tries to access STDIN while backgrounded:
	/*err = f.Exec.Cmd.Process.Signal(syscall.SIGTTIN)
	if err != nil {
		return err
	}*/

	f.State = state.Executing

	shell.StartPrompt()
	return nil
}

func cmdForeground(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	fid, err := p.Parameters.Int(0)
	if err != nil {
		return err
	}

	f, err := proc.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.Exec.Pid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `fg`.")
	}

	if f.Exec.Cmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process.")
	}

	updateTree(f, false)

	proc.ForegroundProc = f

	err = f.Exec.Cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	f.State = state.Executing
	shell.PromptGoProc.Set(f.PromptGoProc)
	return nil
}
