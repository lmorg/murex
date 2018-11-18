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

	if f.ExecPid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `bg`.")
	}

	if f.ExecCmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process.")
	}

	err = f.ExecCmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	updateTree(f, true)

	err = f.ExecCmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	f.State = state.Executing
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

	if f.ExecPid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `fg`.")
	}

	if f.ExecCmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process.")
	}

	updateTree(f, false)

	proc.ForegroundProc = f

	err = f.ExecCmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	f.State = state.Executing
	shell.PromptGoProc.Set(f.PromptGoProc)
	return nil
}
