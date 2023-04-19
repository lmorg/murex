//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package tty

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"

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

func Enabled() bool {
	return height > 0
}

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
		return fmt.Errorf("expecting a bool, instead got %T", t)
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
	primary, replica, err := pty.Open()
	if err != nil {
		return fmt.Errorf("unable to open pty: %s", err.Error())
	}

	size, err := pty.GetsizeFull(os.Stdout)
	if err != nil {
		return fmt.Errorf("unable to get tty size: %s", err.Error())
	}

	err = pty.Setsize(primary, size)
	if err != nil {
		return fmt.Errorf("unable to set pty size: %s", err.Error())
	}
	height = int(size.Rows)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			size, _ := pty.GetsizeFull(os.Stdin)
			pty.Setsize(primary, size)
			height = int(size.Rows)
		}
	}()

	_, err = readline.MakeRaw(int(primary.Fd()))
	if err != nil {
		return fmt.Errorf("unable to put pty into 'raw' mode: %s", err.Error())
	}

	// move cursor to bottom of screen
	//buffer = bytes.Repeat([]byte{'\n'}, height)
	//_, _ = primary.Write(buffer)

	//os.Stdout.WriteString(codes.ClearScreen + codes.Home)
	Stdin, Stdout, Stderr = os.Stdin, primary, primary
	readline.ForceCrLf = false
	readline.SetTTY(primary, os.Stdin)
	go func() {
		ptyBuffer(os.Stdout, replica)
		signal.Stop(ch)
		close(ch)
	}()
	return nil
}

func DestroyPty() {
	Stdin, Stdout, Stderr = os.Stdin, os.Stdout, os.Stderr
	readline.SetTTY(os.Stdout, os.Stdin)
	readline.ForceCrLf = true
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
		go bufferWrite(p[:i])
		written, err := dst.Write(p[:i])
		if err != nil {
			if _, err := os.Stderr.WriteString("error writing to term: " + err.Error()); err != nil {
				panic(err)
			}
			return
		}
		if written != i {
			if _, err := os.Stderr.WriteString(fmt.Sprintf("write mismatch: written %d of %d", written, i)); err != nil {
				panic(err)
			}
			return
		}
	}
}

func bufferWrite(b []byte) {
	bufMutex.Lock()

	buffer = append(buffer, b...)

	var i, count int

	for i = len(buffer) - 1; i != 0; i-- {
		if buffer[i] == '\n' {
			count++
			if count == height {
				buffer = buffer[i:]

				bufMutex.Unlock()
				return
			}
		}

		// clear / cls
		if buffer[i] == 'J' && i > 3 &&
			buffer[i-3] == 27 && buffer[i-2] == '[' &&
			(buffer[i-1] == '2' || buffer[i-1] == '3') {

			/*if i == len(buffer)-1 {
				buffer = bytes.Repeat([]byte{'\n'}, height)
			} else {
				buffer = append(bytes.Repeat([]byte{'\n'}, height), buffer[i+1:]...)
			}*/
			buffer = buffer[i+1:]

			bufMutex.Unlock()
			return
		}

		// disable alternative screen buffer
		// \e[?1049l
		if buffer[i] == 'l' && i > 7 &&
			buffer[i-7] == 27 && buffer[i-6] == '[' && buffer[i-5] == '?' &&
			buffer[i-4] == '1' && buffer[i-3] == '0' && buffer[i-2] == '4' && buffer[i-1] == '9' {

			/*if i == len(buffer)-1 {
				buffer = bytes.Repeat([]byte{'\n'}, height)
			} else {
				buffer = append(bytes.Repeat([]byte{'\n'}, height), buffer[i+1:]...)
			}*/
			buffer = buffer[i+1:]

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

	bufMutex.Lock()

	/*Stdout.WriteString(codes.Reset)
	Stdout.Write(prompt)
	Stdout.WriteString(line)
	Stdout.Write([]byte{'\r', '\n'})
	Stdout.WriteString(codes.BgBlackBright + codes.FgWhiteBright)
	Stdout.WriteString(time.Now().Format(time.RubyDate))
	Stdout.WriteString(codes.Reset)
	Stdout.Write([]byte{'\r', '\n'})*/

	_, _ = os.Stdout.WriteString(codes.Reset)
	_, _ = os.Stdout.WriteString(codes.Home + codes.ClearScreenBelow)

	_, _ = os.Stdout.Write(buffer)

	bufMutex.Unlock()
}

func BufferGet() {
	_, _ = os.Stdout.WriteString(codes.Reset + codes.Home + codes.ClearScreenBelow)
	_, _ = os.Stdout.Write(buffer)
}

var rxEsc = regexp.MustCompile(string([]byte{27, ']', '.', '*', '?', 7})) // no titlebar ANSI escape sequences

func MissingCrLf() bool {
	if height == 0 {
		return false
	}

	bufMutex.Lock()
	buffer := rxEsc.ReplaceAll(buffer, []byte{})

	if len(buffer) > 0 && buffer[len(buffer)-1] != '\n' {
		bufMutex.Unlock()
		return true
	}

	bufMutex.Unlock()
	return false
}

func WriteCrLf() {
	bufMutex.Lock()
	_, _ = Stdout.Write(NewLine)
	bufMutex.Unlock()
}
