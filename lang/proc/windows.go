// +build windows

package proc

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func External(p *Process) error {
	// External executable
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
	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	// Horrible fudge adding cmd /c to the parameters but this is to get around cmd.exe builtins.
	parameters := append([]string{"cmd", "/c"}, p.Parameters.StringArray()...)
	cmd := exec.Command(exeName, parameters[1:]...)

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

func ExternalPty(p *Process) error {
	// External executable
	if err := shellExecute(p); err != nil {
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

func shellExecute(p *Process) error {
	// Create an object for the executable we wish to invoke.
	exeName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	// Horrible fudge adding cmd /c to the parameters but this is to get around cmd.exe builtins.
	parameters := append([]string{"cmd", "/c"}, p.Parameters.StringArray()...)
	cmd := exec.Command(exeName, parameters[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// Create an STDIN function, copying 1KB blocks at a time.
	active := true
	go func() {
		defer stdin.Close()
		b := make([]byte, 1024)
		for active {
			i, err := os.Stdin.Read(b)
			if err != nil {
				return
			}
			if _, err := stdin.Write(b[:i]); err != nil {
				return
			}
		}
		stdin.Close()
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	active = false
	return nil
}
