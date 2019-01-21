package net

import (
	"io"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils"
)

// Write bytes to net Io interface
func (n *Net) Write(b []byte) (i int, err error) {
	select {
	case <-n.ctx.Done():
		return 0, io.ErrClosedPipe
	default:
	}

	i, err = n.conn.Write(b)
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

// Writeln writes a line to net Io interface
func (n *Net) Writeln(b []byte) (i int, err error) {
	i, err = n.conn.Write(append(b, utils.NewLineByte...))
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (n *Net) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(n, dataType)
}

// WriteTo reads from net Io interface and then writes that to foreign Writer interface
func (n *Net) WriteTo(dst io.Writer) (i int64, err error) {
	i, err = io.Copy(dst, n.conn)
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}
