package tty

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/lmorg/murex/utils/ansi/codes"
	"github.com/lmorg/murex/utils/readline"
)

var (
	Stdin  *os.File = os.Stdin
	Stdout *os.File = os.Stdout
	Stderr *os.File = os.Stderr

	buffer []byte
	hight  int
)

func CreatePTY() {
	p, t, err := pty.Open()
	if err != nil {
		return
	}

	size, err := pty.GetsizeFull(os.Stdout)
	if err != nil {
		return
	}
	pty.Setsize(p, size)
	hight = int(size.Rows)
	//_ = pty.InheritSize(os.Stdout, p)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			size, _ := pty.GetsizeFull(os.Stdout)
			pty.Setsize(p, size)
			hight = int(size.Rows)
			//_ = pty.InheritSize(os.Stdout, p)
		}
	}()

	_, err = readline.MakeRaw(int(p.Fd()))
	if err != nil {
		return
	}

	os.Stdout.WriteString(codes.ClearScreen + codes.Home)
	Stdout, Stderr = p, p
	//readline.SetTTY(p) // this doesn't behave as expected
	readline.ForceCrLf = false
	go func() {
		ptyBuffer(os.Stdout, t)
		signal.Stop(ch)
		close(ch)
	}()
}

func ptyBuffer(dst, src *os.File) {
	p := make([]byte, 10*1024)

	for {
		i, err := src.Read(p)
		if err != nil {
			//continue
			panic(err)
		}
		written, err := dst.Write(p[:i])
		if err != nil {
			panic(err)
		}
		if written != i {
			panic(fmt.Sprintf("write mistmatch: written %d of %d", written, i))
		}

		bufferWrite(p[:i])
	}
}

func bufferWrite(b []byte) {
	buffer = append(buffer, b...)
	var i, count int
	for i = len(buffer) - 1; i != 0; i-- {
		if buffer[i] == '\n' {
			count++
		}
		if count == hight {
			buffer = buffer[i:]
			return
		}
	}
}

func BufferRecall(prompt []byte, line string) {
	if hight == 0 {
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

	_, _ = os.Stdout.Write(buffer)
}
