package proc

import (
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// External executes an external process.
func External(p *Process) error {
	if err := execute(p); err != nil {
		// Get exit status. This has only been tested on Linux. May not work on other OSs.
		if strings.HasPrefix(err.Error(), "exit status ") {
			i, _ := strconv.Atoi(strings.Replace(err.Error(), "exit status ", "", 1))
			p.ExitNum = i
		} else {
			p.Stderr.Writeln([]byte(err.Error()))
			p.ExitNum = 1
		}

	}
	return nil
}

func execute(p *Process) error {
	p.Stdout.SetDataType(types.Generic)

	exeName, parameters, err := getCmdTokens(p)
	if err != nil {
		return err
	}
	cmd := exec.Command(exeName, parameters...)

	p.Kill = func() {
		defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		cmd.Process.Kill()
	}
	KillForeground = p.Kill

	if p.Name == consts.CmdPty {
		// If cmd is `pty` then assign function a TTY (on POSIX) and allow it to read and write directly from STDIN et al
		osSyscalls(cmd)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		// If cmd is `exec` then the input and output streams are the murex stdio.Io rather than STD(IN|OUT|ERR)
		cmd.Stdin = p.Stdin
		cmd.Stdout = p.Stdout
		cmd.Stderr = p.Stderr
	}

	if err := cmd.Start(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	return nil
}
