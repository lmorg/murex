package streams

import (
	"bufio"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sync"
)

/*type Io interface {
	MakeParent()
	UnmakeParent()
	MakePipe()

	Stats() (uint64, uint64)

	GetDataType() string
	SetDataType(string)
	DefaultDataType(bool)

	ReadArray(callback func([]byte))
	ReadMap(*config.Config, func(string, string, bool)) error


	Close()
}*/

type Net struct {
	mutex    sync.Mutex
	buffer   []byte
	closed   bool
	bRead    uint64
	bWritten uint64
	isParent bool
	conn     net.Conn
	dataType string
	dtLock   sync.Mutex
	protocol string
}

func NewNetDialer(protocol, address string) (n *Net, err error) {
	n = new(Net)
	n.protocol = protocol

	if protocol == "udp" || protocol == "tcp" {
		n.dataType = types.Generic
	} else {
		protocol = "tcp"
	}

	n.conn, err = net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}

	return
}

func (n *Net) Stats() (bytesWritten, bytesRead uint64) {
	n.mutex.Lock()
	bytesWritten = n.bWritten
	bytesRead = n.bRead
	n.mutex.Unlock()
	return
}

func (n *Net) Read(p []byte) (i int, err error) {
	i, err = n.conn.Read(p)
	n.mutex.Lock()
	n.bRead += uint64(i)
	n.mutex.Unlock()
	return
}

func (n *Net) ReadLine(callback func([]byte)) {
	scanner := bufio.NewScanner(n)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	if err := scanner.Err(); err != nil {
		//os.Stderr.WriteString("ReadLine: " + err.Error() + utils.NewLineByte)
		//return errors.New("ReadLine: " + err.Error())
		return
	}

	return
}

func (n *Net) ReadAll() (b []byte, err error) {
	b, err = ioutil.ReadAll(n.conn)
	n.mutex.Lock()
	n.bRead += uint64(len(b))
	n.mutex.Unlock()
	return
}

func (n *Net) Write(b []byte) (i int, err error) {
	i, err = n.conn.Write(b)
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

func (n *Net) Writeln(b []byte) (i int, err error) {
	i, err = n.conn.Write(append(b, utils.NewLineByte...))
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

func (n *Net) WriteTo(dst io.Writer) (i int64, err error) {
	i, err = io.Copy(dst, n.conn)
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

func (n *Net) Close() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.isParent {
		// This will legitimately happen a lot since the reason we mark a stream as parent is to prevent
		// accidental closing. However it's worth pushing a message out in debug mode during this alpha build.
		debug.Log("Cannot Close() net marked as parent. We don't want to EOT parent streams multiple times")
		return
	}

	if n.closed {
		os.Stderr.WriteString("Error with murex named pipes: Trying to close an already closed named pipe." + utils.NewLineString)
		return
	}

	n.closed = true
	err := n.conn.Close()
	if err != nil {
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
	}
}
