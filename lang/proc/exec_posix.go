// +build !windows

package proc

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/readline"
)

func getCmdTokens(p *Process) (exe string, parameters []string, err error) {
	exe, err = p.Parameters.String(0)
	if err != nil {
		return
	}

	parameters = p.Parameters.StringArray()[1:]

	return
}

func osSyscalls(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ctty: int(os.Stdout.Fd()),
	}

	return
}

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
