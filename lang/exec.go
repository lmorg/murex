package lang

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/which"
)

const (
	envMethodTrue  = consts.EnvMethod + "=" + consts.EnvTrue
	envMethodFalse = consts.EnvMethod + "=" + consts.EnvFalse

	envBackgroundTrue  = consts.EnvBackground + "=" + consts.EnvTrue
	envBackgroundFalse = consts.EnvBackground + "=" + consts.EnvFalse

	envDataType = consts.EnvDataType + "="
)

var (
	envMurexPid = fmt.Sprintf("%s=%d", consts.EnvMurexPid, os.Getpid())
	//termOut     = reflect.TypeOf(new(term.Out)).String()
)

func forceTTY(p *Process) bool {
	v, err := p.Config.Get("proc", "force-tty", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

// External executes an external process.
func External(p *Process) error {
	err := execute(p)
	if err != nil {
		if p.SystemProcess != nil {
			p.ExitNum = p.SystemProcess.ExitNum()
		} else {
			p.ExitNum = 1
		}
		return err
	}

	return nil
}

func execute(p *Process) error {
	argv := osExecGetArgv(p)
	if len(argv) == 0 {
		return fmt.Errorf("no parameters specified for `exec`\nExpecting a name of command to run")
	}

	if !p.IsMethod && p.Stdout.IsTTY() && p.Stderr.IsTTY() && forceTTY(p) {
		return osExecFork(p, argv)
	}

	return execForkFallback(p, argv)
}

func execForkFallback(p *Process, argv []string) error {
	cmd := exec.Command(which.WhichIgnoreFail(argv[0]), argv[1:]...)

	if p.HasCancelled() {
		return nil
	}
	p.Kill = func() {
		if !debug.Enabled {
			defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		}

		//ctxCancel()
		err := cmd.Process.Signal(syscall.SIGTERM)
		if err != nil {
			if err.Error() == os.ErrProcessDone.Error() {
				return
			}
			name, _ := p.Args()
			os.Stderr.WriteString(
				fmt.Sprintf("\n!!! Error sending SIGTERM to `%s`: %s\n", name, err.Error()))
		}
	}

	// ***
	// Define STANDARD IN (fd 0)
	// ***

	switch {
	case p.IsMethod:
		cmd.Stdin = p.Stdin
		if p.Background.Get() {
			cmd.Env = append(os.Environ(), envMurexPid, envMethodTrue, envBackgroundTrue, envDataType+p.Stdin.GetDataType())
		} else {
			cmd.Env = append(os.Environ(), envMurexPid, envMethodTrue, envBackgroundFalse, envDataType+p.Stdin.GetDataType())
		}
	case p.Background.Get():
		cmd.Stdin = new(null.Null)
		cmd.Env = append(os.Environ(), envMurexPid, envMethodFalse, envBackgroundTrue, envDataType+p.Stdin.GetDataType())
	default:
		cmd.Stdin = os.Stdin
		cmd.Env = append(os.Environ(), envMurexPid, envMethodFalse, envBackgroundFalse, envDataType+p.Stdin.GetDataType())
	}

	/*if len(p.Envs) > 0 {
		for i := range p.Envs {
			if !strings.Contains(p.Envs[i], "=") {
				v, err := p.Variables.GetString(p.Envs[i])
				if err != nil {
					continue
				}
				p.Envs[i] += "=" + v
			}
		}
		cmd.Env = append(cmd.Env, p.Envs...)
	}*/

	// ***
	// Define STANDARD OUT (fd 1)
	// ***

	if p.Stdout.IsTTY() {
		// If Stdout is a TTY then set the appropriate syscalls to allow the calling program to own the TTY....
		//osSyscalls(cmd, int(p.ttyout.Fd()))
		//osSyscalls(cmd, int(os.Stdin.Fd()))
		//if reflect.TypeOf(p.Stdout).String() == termOut {
		//	osSyscalls(cmd, int(p.ttyout.Fd()))
		//	cmd.Stdout = p.ttyout
		//} else {

		cmd.SysProcAttr = osSysProcAttr(int(p.Stdout.File().Fd()))

		cmd.Stdout = p.Stdout.File()
		//}
	} else {
		// ....otherwise we just treat the program as a regular piped util
		cmd.Stdout = p.Stdout
	}

	// ***
	// Define STANDARD ERR (fd 2)
	// ***

	// Pipe STDERR irrespective of whether the exec process is outputting to a TTY or not.
	// The reason for this is so that we can do some post-processing on the error stream (namely add colour to it),
	// however this might cause some bugs. If so please raise on github: https://github.com/lmorg/murex
	// In the meantime, you can force exec processes to write STDERR to the TTY via the `config` command in the shell:
	//
	//     config set proc force-tty true
	if p.Stderr.IsTTY() && forceTTY(p) {
		cmd.Stderr = p.Stderr.File()
	} else {
		cmd.Stderr = p.Stderr
	}

	// ***
	// Define MUREX DATA TYPE (fd 3)
	// ***

	/*var failedPipe bool
	mxdtR, mxdtW, err := os.Pipe()
	if err != nil {
		tty.Stderr.WriteString("unable to create murex data type output file for external process: " + err.Error() + "\n")
		failedPipe = true
		mxdtR = new(os.File)
		mxdtW = new(os.File)

	} else {
		cmd.ExtraFiles = []*os.File{mxdtW}
	}*/

	// ***
	// Start process
	// ***

	p.State.Set(state.Executing)
	if err := cmd.Start(); err != nil {
		//if !strings.HasPrefix(err.Error(), "signal:") {
		//mxdtW.Close()
		//mxdtR.Close()
		return err
		//}
	}

	// ***
	// Get murex data type
	// ***

	/*go func() {
		if failedPipe {
			p.Stdout.SetDataType(types.Generic)
			return
		}

		var dt string

		scanner := bufio.NewScanner(mxdtR)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			dt = scanner.Text()
			break
		}

		if scanner.Err() != nil || dt == "" {
			dt = types.Generic
		}

		p.Stdout.SetDataType(dt)
		mxdtR.Close()
	}()*/

	/////////

	p.SystemProcess = &sysProcT{cmd}

	/*if err := mxdtW.Close(); err != nil {
		tty.Stderr.WriteString("error closing murex data type output file write pipe:" + err.Error() + "\n")
	}*/

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			//mxdtR.Close()
			return err
		}
	}

	//mxdtR.Close()
	return nil
}

type sysProcT struct{ cmd *exec.Cmd }

func (sp *sysProcT) Pid() int                   { return sp.cmd.Process.Pid }
func (sp *sysProcT) ExitNum() int               { return sp.cmd.ProcessState.ExitCode() }
func (sp *sysProcT) Kill() error                { return sp.cmd.Process.Kill() }
func (sp *sysProcT) Signal(sig os.Signal) error { return sp.cmd.Process.Signal(sig) }
