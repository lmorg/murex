//go:build !windows && !js && !plan9 && !no_pty
// +build !windows,!js,!plan9,!no_pty

package psuedotty

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"

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

// GetDataType returns the murex data type for the stream.Io interface
func (p *PTY) GetDataType() (dt string) { return p.out.GetDataType() }

// SetDataType defines the murex data type for the stream.Io interface
func (p *PTY) SetDataType(dt string) { p.out.SetDataType(dt) }

// Stats provides real time stream stats. Useful for progress bars etc.
func (p *PTY) Stats() (uint64, uint64) { return p.out.Stats() }

// IsTTY returns true because the PTY stream is a pseudo-TTY
func (p *PTY) IsTTY() bool { return true }

// File returns the os.File struct for the stream.Io interface if a TTY
func (p *PTY) File() *os.File { return p.in }

// Open the stream.Io interface for another dependant
func (p *PTY) Open() {
	atomic.AddInt32(&p.dependents, 1)
	p.out.Open()
}

// Close the stream.Io interface
func (p *PTY) Close() {
	i := atomic.AddInt32(&p.dependents, -1)
	if i < 0 {
		panic("More closed dependents than open")
	}
	p.out.Close()
	if i == 0 {
		go p.close()
	}
}

func (p *PTY) close() {
	defer p.out.ForceClose()

	for {
		time.Sleep(1 * time.Second)
		w, r := p.out.Stats()
		if r >= w {
			err := p.in.Close()
			if err != nil {
				panic(err)
			}
			err = p.replica.Close()
			if err != nil {
				panic(err)
			}
			return
		}
	}
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (p *PTY) ForceClose() {
	p.in.Close()
	p.replica.Close()
	p.out.ForceClose()
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
