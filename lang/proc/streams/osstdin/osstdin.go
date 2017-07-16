package osstdin

// This package might seem a bit weird as it's just a wrapper around os.Stdin,
// but it is needed to work around the issue with PTY's stealing focus.

import (
	"os"
	"sync"
)

var BuffSize int = 1024 * 1024 * 10

type stdin struct {
	mutex sync.Mutex
	//mutex  debug.BadMutex
	data   chan []byte
	buffer []byte
}

var Stdin *stdin

func init() {
	Stdin = new(stdin)
	Stdin.data = make(chan []byte)
}

func (in *stdin) Prepend(b []byte) {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	in.buffer = append(b, in.buffer...)
}

func (in *stdin) Read(p []byte) (i int, err error) {
	err = nil

	in.mutex.Lock()
	defer in.mutex.Unlock()

	go read()

	if len(in.buffer) == 0 {
		in.buffer = <-in.data
	}

	if cap(p) < len(in.buffer) {
		copy(p, in.buffer[:cap(p)])
		in.buffer = in.buffer[cap(p)-1:]
		return cap(p), nil
	}

	copy(p, in.buffer)
	i = len(in.buffer)
	in.buffer = make([]byte, 0)
	return
}

func read() {
	p := make([]byte, BuffSize)

	var in []byte

	for {
		i, err := os.Stdin.Read(p)

		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

		in = append(in, p[:i]...)

		if i < BuffSize {
			break
		}
	}

	Stdin.data <- in
}
