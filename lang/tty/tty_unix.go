//go:build !windows && !plan && !js
// +build !windows,!plan,!js

package tty

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/readline"
)

var (
	buffer   []byte
	bufMutex sync.Mutex
	height   int
)

func ConfigRead() (interface{}, error) {
	return height > 0, nil
}

func ConfigWrite(v interface{}) error {
	switch t := v.(type) {
	case bool:
		return EnableDisable(v.(bool))
	case string, []byte:
		v, err := types.ConvertGoType(t, types.Boolean)
		if err != nil {
			return err
		}
		return EnableDisable(v.(bool))
	default:
		return fmt.Errorf("expectig a bool, instead got %T", t)
	}
}

func EnableDisable(v bool) error {
	if v == (height > 0) {
		return nil
	} else {
		if v {
			return CreatePTY()
		} else {
			DestroyPty()
			return nil
		}
	}
}

func CreatePTY() error {
	pOut, tOut, err := pty.Open()
	if err != nil {
		return fmt.Errorf("unable to open pty for stdout: %s", err.Error())
	}

	size, err := pty.GetsizeFull(os.Stdout)
	if err != nil {
		return fmt.Errorf("unable to get tty size for stdout: %s", err.Error())
	}

	err = pty.Setsize(pOut, size)
	if err != nil {
		return fmt.Errorf("unable to set pty size for stdout: %s", err.Error())
	}

	height = int(size.Rows)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			size, _ := pty.GetsizeFull(os.Stdout)
			pty.Setsize(pOut, size)
			height = int(size.Rows)
		}
	}()

	_, err = readline.MakeRaw(int(pOut.Fd()))
	if err != nil {
		return fmt.Errorf("unable to set put pty for stdout into 'raw' mode: %s", err.Error())
	}

	os.Stdout.WriteString(codes.ClearScreen + codes.Home)
	Stdin, Stdout, Stderr = os.Stdin, pOut, pOut
	readline.ForceCrLf = false
	go func() {
		ptyBuffer(os.Stdout, tOut)
		signal.Stop(ch)
		close(ch)
	}()
	return nil
}

func DestroyPty() {
	_ = Stdout.Close()
	Stdout, Stderr = os.Stdout, os.Stderr
	height = 0
	bufMutex.Lock()
	buffer = []byte{}
	bufMutex.Unlock()
}

func ptyBuffer(dst, src *os.File) {
	p := make([]byte, 10*1024)

	for {
		i, err := src.Read(p)
		if err != nil {
			if err.Error() != io.EOF.Error() {
				if _, err := os.Stderr.WriteString("error reading from PTY: " + err.Error()); err != nil {
					panic(err)
				}
			}
			return
		}
		written, err := dst.Write(p[:i])
		if err != nil {
			if _, err := os.Stderr.WriteString("error writing to term: " + err.Error()); err != nil {
				panic(err)
			}
			return
		}
		if written != i {
			if _, err := os.Stderr.WriteString(fmt.Sprintf("write mistmatch: written %d of %d", written, i)); err != nil {
				panic(err)
			}
			return
		}

		bufferWrite(p[:i])
	}
}

func bufferWrite(b []byte) {
	bufMutex.Lock()
	buffer = append(buffer, b...)
	var i, count int
	for i = len(buffer) - 1; i != 0; i-- {
		if buffer[i] == '\n' {
			count++
		}
		if count == height {
			buffer = buffer[i:]
			bufMutex.Unlock()
			return
		}
	}
	bufMutex.Unlock()
}

func BufferRecall(prompt []byte, line string) {
	if height == 0 {
		// height unset so lets assume no PTY created
		return
	}

	/*if len(buffer) > 0 && buffer[len(buffer)-1] != '\n' {
		//Stdout.Write([]byte{'\r', '\n'})
		buffer = append(buffer, '\r', '\n')
	}*/

	Stdout.WriteString(codes.Reset)
	Stdout.Write(prompt)
	Stdout.WriteString(line)
	Stdout.Write([]byte{'\r', '\n'})
	Stdout.WriteString(codes.BgBlackBright + codes.FgWhiteBright)
	Stdout.WriteString(time.Now().Format(time.RubyDate))
	Stdout.WriteString(codes.Reset)
	Stdout.Write([]byte{'\r', '\n'})

	_, _ = os.Stdout.WriteString(codes.Reset)
	_, _ = os.Stdout.WriteString(codes.Home + codes.ClearScreenBelow)

	bufMutex.Lock()
	_, _ = os.Stdout.Write(buffer)
	bufMutex.Unlock()
}
