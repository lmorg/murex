package psuedotty

import (
	"fmt"
	"io"
	"os"

	"github.com/creack/pty"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

func init() {
	stdio.RegisterPipe("pty", registerPipe)
}

func registerPipe(_ string) (stdio.Io, error) {
	return NewPTY(80, 25)
}

type PTY struct {
	in         *os.File
	out        stdio.Io
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

	p := new(PTY)
	p.in = primary
	p.out = streams.NewReader(replica)

	return p, nil
}

func (p *PTY) Write(b []byte) (int, error) {
	return p.in.Write(b)
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
