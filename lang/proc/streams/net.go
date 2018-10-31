package streams

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"net"
	"os"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// Net Io interface
type Net struct {
	mutex      sync.Mutex
	ctx        context.Context
	forceClose func()
	bRead      uint64
	bWritten   uint64
	dependants int
	conn       net.Conn
	dataType   string
	protocol   string
}

// DefaultDataType is unavailable for net Io interfaces
func (n *Net) DefaultDataType(bool) {}

// IsTTY always returns false because net Io interfaces are not a pseudo-TTY
func (n *Net) IsTTY() bool { return false }

// NewDialer creates a new net.Dial-based stream.Io pipe
func NewDialer(protocol, address string) (n *Net, err error) {
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

	n.ctx, n.forceClose = context.WithCancel(context.Background())

	return
}

// NewListener creates a new net.Listen-based stream.Io pipe
func NewListener(protocol, address string) (n *Net, err error) {
	n = new(Net)
	n.protocol = protocol

	if protocol == "udp" || protocol == "tcp" {
		n.dataType = types.Generic
	} else {
		protocol = "tcp"
	}

	listen, err := net.Listen(protocol, address)
	if err != nil {
		return nil, err
	}
	defer listen.Close()

	n.conn, err = listen.Accept()
	if err != nil {
		return nil, err
	}

	n.ctx, n.forceClose = context.WithCancel(context.Background())

	return
}

// MakePipe turns the stream.Io interface into a named pipe
func (n *Net) MakePipe() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	//n.isParent = true
	n.dependants++
}

// SetDataType assigns a data type to the stream.Io interface
func (n *Net) SetDataType(dt string) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dataType = dt
}

// GetDataType read the stream.Io interface's data type
func (n *Net) GetDataType() string {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	return n.dataType
}

// Stats returns the bytes written and bytes read to the net Io interface
func (n *Net) Stats() (bytesWritten, bytesRead uint64) {
	n.mutex.Lock()
	bytesWritten = n.bWritten
	bytesRead = n.bRead
	n.mutex.Unlock()
	return
}

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
	b, err = ioutil.ReadAll(n.conn)
	n.mutex.Lock()
	n.bRead += uint64(len(b))
	n.mutex.Unlock()
	return
}

// Write bytes to net Io interface
func (n *Net) Write(b []byte) (i int, err error) {
	/*select {
	case <-n.ctx.Done():
		return 0, io.EOF
	}*/

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

// WriteTo reads from net Io interface and then writes that to foreign Writer interface
func (n *Net) WriteTo(dst io.Writer) (i int64, err error) {
	i, err = io.Copy(dst, n.conn)
	n.mutex.Lock()
	n.bWritten += uint64(i)
	n.mutex.Unlock()
	return
}

// Open net Io interface
func (n *Net) Open() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dependants++
}

// Close net Io interface
func (n *Net) Close() {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.dependants--

	if n.dependants == 0 {
		err := n.conn.Close()
		if err != nil {
			os.Stderr.WriteString(err.Error() + utils.NewLineString)
		}
	}

	if n.dependants < 0 {
		panic("more closed dependants than open")
	}
}

// ForceClose forces the net Io interface to close. This should only be called on reader interfaces
func (n *Net) ForceClose() {
	n.forceClose()
}

// ReadArray treats net Io interface as an array of data
func (n *Net) ReadArray(callback func([]byte)) error {
	return readArray(n, callback)
}

// ReadMap treats net Io interface as an hash of data
func (n *Net) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(n, config, callback)
}
