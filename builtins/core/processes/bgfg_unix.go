// +build !windows,!plan9

package processes

import (
	"errors"
	"syscall"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
)

func mkbg(p *lang.Process) error {
	fid, err := p.Parameters.Uint32(0)
	if err != nil {
		return errors.New("Invalid parameters. Expecting either a code block or FID of a stopped process")
	}

	f, err := lang.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.State.Get() != state.Stopped {
		return errors.New("FID is not a stopped process. Run `jobs` or `fid-list` to see a list of stopped processes")
	}

	if f.Exec.Pid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `bg`")
	}

	if f.Exec.Cmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process")
	}

	updateTree(f, true)

	if !f.IsMethod {
		// This doesn't work. But we would need something clever like this:
		//f.Exec.Cmd.Stdin = streams.NewStdin()
	}

	// This doesn't belong here. It should only be called if the requesting program tries to access STDIN while backgrounded:
	/*err = f.Exec.Cmd.Process.Signal(syscall.SIGTTIN)
	if err != nil {
		return err
	}*/

	/*f.Exec.PipeR, f.Exec.PipeW, err = os.Pipe()
	if err != nil {
		return err
	}

	f.Stdin = streams.NewReadCloser(f.Exec.PipeR)

	f.Exec.Cmd.SysProcAttr.Setctty = true
	f.Exec.Cmd.SysProcAttr.Ctty = int(f.Exec.PipeR.Fd())
	f.Exec.Cmd.SysProcAttr.Foreground = false*/

	/*= &syscall.SysProcAttr{
		Setsid:     true,
		Setctty:    true,
		Ctty:       int(f.Exec.PipeR.Fd()),
		Foreground: false,
	}*/

	err = f.Exec.Cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}

	f.State.Set(state.Executing)

	shell.ShowPrompt()
	return nil
}

func cmdForeground(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	fid, err := p.Parameters.Uint32(0)
	if err != nil {
		return err
	}

	f, err := lang.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.Exec.Pid == 0 {
		return errors.New("This FID doesn't have an associated PID. Murex functions currently don't support `fg`")
	}

	if f.Exec.Cmd == nil {
		return errors.New("Something went wrong trying to communicate back to the OS process")
	}

	updateTree(f, false)

	//lang.ForegroundProc = f
	lang.ForegroundProc.Set(f)

	//if !f.Exec.Cmd.ProcessState.Exited() {
	err = f.Exec.Cmd.Process.Signal(syscall.SIGCONT)
	if err != nil {
		return err
	}
	//}

	f.State.Set(state.Executing)
	shell.PromptId.Set(f.PromptId)
	return nil
}
