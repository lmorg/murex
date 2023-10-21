package psuedotty

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/creack/pty"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	stdio.RegisterPipe("pty", registerPipe)
}

func registerPipe(_ string) (stdio.Io, error) {
	return NewPTY(80, 25)
}

type PTY struct {
	in         *os.File
	replica    *os.File
	out        *streams.Stdin
	dependents int32
}

func NewPTY(width, height int) (*PTY, error) {
	primary, replica, err := pty.Open()
	if err != nil {
		return nil, fmt.Errorf("unable to open pty: %s", err.Error())
	}

	size := pty.Winsize{
		Cols: uint16(width),
		Rows: uint16(height),
	}

	err = pty.Setsize(primary, &size)
	if err != nil {
		return nil, fmt.Errorf("unable to set pty size: %s", err.Error())
	}

	_, err = readline.MakeRaw(int(primary.Fd()))
	if err != nil {
		return nil, fmt.Errorf("unable to set pty state: %s", err.Error())
	}

	p := new(PTY)
	p.in = primary
	p.replica = replica
	p.out = streams.NewStdin()
	p.out.Open()

	go func() {
		i, err := io.Copy(p.out, p.replica)
		if err != nil && debug.Enabled {
			os.Stderr.WriteString(fmt.Sprintf("!!! read failed from PTY after %d bytes: %v !!!", i, err))
		}
	}()

	return p, nil
}

func (p *PTY) Read(b []byte) (int, error) {
	return p.out.Read(b)
}

func (p *PTY) ReadLine(callback func([]byte)) error {
	return p.out.ReadLine(callback)
}

func (p *PTY) ReadArray(ctx context.Context, callback func([]byte)) error {
	return p.out.ReadArray(ctx, callback)
}

func (p *PTY) ReadArrayWithType(ctx context.Context, callback func(interface{}, string)) error {
	return p.out.ReadArrayWithType(ctx, callback)
}

func (p *PTY) ReadMap(conf *config.Config, callback func(*stdio.Map)) error {
	return p.out.ReadMap(conf, callback)
}

func (p *PTY) ReadAll() ([]byte, error) {
	b, err := p.out.ReadAll()
	_ = p.in.Close()
	p.out.Close()
	return b, err
}

func (p *PTY) Write(b []byte) (int, error) {
	i, err := p.in.Write(b)
	return i, err
}

func (p *PTY) Writeln(b []byte) (int, error) {
	slice := append(b, utils.NewLineByte...)
	return p.in.Write(slice)
}

func (p *PTY) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(p, dataType)
}

func (p *PTY) WriteTo(w io.Writer) (int64, error) {
	return stdio.WriteTo(p, w)
}
