package proc

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/readline"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
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

	if p.HasCancelled() {
		return nil
	}

	ctxCancel := p.Kill
	p.Kill = func() {
		if !debug.Enable {
			defer func() { recover() }() // I don't care about errors in this instance since we are just killing the proc anyway
		}

		ctxCancel()
		cmd.Process.Kill()
	}

	if p.IsMethod {
		cmd.Stdin = p.Stdin
	} else {
		cmd.Stdin = os.Stdin

		/*p.Exec.Stdin, err = newStdinPipe(p, cmd)
		if err != nil {
			return err
		}
		p.Exec.Stdin.Foreground()*/
	}

	if p.Stdout.IsTTY() {
		// If Stdout is a TTY then set the appropriate syscalls to allow the calling program to own the TTY....
		osSyscalls(cmd)
		cmd.Stdout = os.Stdout
	} else {
		// ....otherwise we just treat the program as a regular piped util
		cmd.Stdout = p.Stdout
	}

	// Pipe STDERR irrespective of whether the exec process is execting a TTY or not.
	// The reason for this is so that we can do some post-processing on the error stream (namely add colour to it),
	// however this might cause some bugs. If so please raise on github: https://github.com/lmorg/murex
	// In the meantime, you can force exec processes to write STDERR to the TTY via the `config` command in the shell:
	//
	//     config set proc force-tty true
	if p.Stderr.IsTTY() && forceTTY(p) {
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stderr = p.Stderr
	}

	if err := cmd.Start(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	p.Exec.Pid = cmd.Process.Pid
	p.Exec.Cmd = cmd

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	//debug.Log("exec env:",cmd.Env)

	return nil
}

func forceTTY(p *Process) bool {
	v, err := p.Config.Get("proc", "force-tty", types.Boolean)
	if err != nil {
		return false
	}
	return v.(bool)
}

///////////////////////////////////

type StdinPipe struct {
	f    *os.File
	proc *Process
	bg   chan bool
}

func newStdinPipe(p *Process, cmd *exec.Cmd) (pipe *StdinPipe, err error) {
	pipe = new(StdinPipe)

	pipe.bg = make(chan bool)

	//var r *os.File
	//r, pipe.f, err = os.Pipe()
	//if err != nil {
	//	return nil, err
	//}

	path := consts.TempDir + strconv.Itoa(p.Id) + ".stdin"
	//os.Create(path)
	//err = syscall.Mkfifo(path, 0644) //uint32(os.O_APPEND|os.O_CREATE|os.O_WRONLY))
	err = syscall.Mknod(path, 0644|syscall.S_IFCHR, 0)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	defer f.Close()

	fd := int(f.Fd())

	_, err = readline.MakeRaw(fd)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	cmd.Stdin = f
	pipe.f = f

	return pipe, nil
}

// Background disables the stdin listener
func (pipe *StdinPipe) Background() {
	go func() {
		pipe.bg <- true
	}()
}

// Foreground enables the stdin listener
func (pipe *StdinPipe) Foreground() {
	go pipe.read(os.Stdin)
}

func (pipe *StdinPipe) read(r io.Reader) {
	b := make([]byte, 1024)
	for {
		select {
		case <-pipe.proc.Context.Done():
			return

		case <-pipe.bg:
			return

		default:
			i, err := r.Read(b)
			if err != nil {
				panic(err.Error())
			}

			_, err = pipe.f.Write(b[:i])
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
