package net

import (
	"bufio"
	"context"
	"io"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils"
)

// Read bytes from net Io interface
func (n *Net) Read(p []byte) (i int, err error) {
	select {
	case <-n.ctx.Done():
		return 0, io.EOF
	default:
	}

	i, err = n.conn.Read(p)
	n.mutex.Lock()
	n.bRead += uint64(i)
	n.mutex.Unlock()
	return
}

// ReadLine reads a line from net Io interface
func (n *Net) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(n)
	for scanner.Scan() {
		b := scanner.Bytes()
		n.mutex.Lock()
		n.bRead += uint64(len(b))
		n.mutex.Unlock()
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll data from net Io interface
func (n *Net) ReadAll() (b []byte, err error) {
	w := streams.NewStdinWithContext(n.ctx, n.forceClose)

	_, err = w.ReadFrom(n.conn)
	if err != nil {
		return
	}

	b, err = w.ReadAll()

	n.mutex.Lock()
	n.bRead += uint64(len(b))
	n.mutex.Unlock()

	return
}

// ReadArray treats net Io interface as an array of data
func (n *Net) ReadArray(ctx context.Context, callback func([]byte)) error {
	return stdio.ReadArray(ctx, n, callback)
}

// ReadArrayWithType treats net Io interface as an array of data
func (n *Net) ReadArrayWithType(ctx context.Context, callback func(interface{}, string)) error {
	return stdio.ReadArrayWithType(ctx, n, callback)
}

// ReadMap treats net Io interface as an hash of data
func (n *Net) ReadMap(config *config.Config, callback func(*stdio.Map)) error {
	return stdio.ReadMap(n, config, callback)
}
