// +build windows

package proc

import (
	"github.com/lmorg/murex/lang/types"
	"os/exec"
	"strconv"
	"strings"
)

// Execute an external (Windows) process
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
	p.Stdout.SetDataType(types.String)

	// Horrible fudge adding cmd /c to the parameters,
	// this is to get around half the useful features in cmd.exe being builtins.
	parameters := append([]string{"cmd", "/c"}, p.Parameters.StringArray()...)
	cmd := exec.Command(parameters[0], parameters...)

	cmd.Stdin = p.Stdin
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// This function exists for POSIX builds. Since Windows doesn't have TTYs, lets just make it a wrapper around the
// first function in this file
func ExternalPty(p *Process) error {
	return External(p)
}
