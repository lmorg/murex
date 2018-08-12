package streams

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// Reader is a wrapper around an io.Reader interface
type Reader struct {
	mutex      sync.Mutex
	reader     io.Reader
	readCloser io.ReadCloser
	bRead      uint64
	bWritten   uint64
	dependants int
	dataType   string
	dtLock     sync.Mutex
}

// NewReader creates a new Stdio.Io interface wrapper around a io.Reader interface
func NewReader(reader io.Reader) (r *Reader) {
	if reader == nil {
		panic("streams.Reader interface has nil reader interface")
	}

	r = new(Reader)
	r.reader = reader
	return
}

// IsTTY returns false because the reader interface is not a pseudo-TTY
func (r *Reader) IsTTY() bool { return false }

// MakePipe is used for named pipes. Basically just used to relax the exception
// handling since we can make fewer guarantees about the state of named pipes.
func (r *Reader) MakePipe() {
	r.mutex.Lock()
	r.dependants++
	r.mutex.Unlock()
}

// Stats provides real time stream stats. Useful for progress bars etc.
func (r *Reader) Stats() (bytesWritten, bytesRead uint64) {
	r.mutex.Lock()
	bytesWritten = r.bWritten
	bytesRead = r.bRead
	r.mutex.Unlock()
	return
}

// Read is the reader interface Read() method.
func (r *Reader) Read(p []byte) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	i, err := r.reader.Read(p)

	r.bRead += uint64(i)
	return i, err
}

// ReadLine returns each line in the stream as a callback function
func (r *Reader) ReadLine(callback func([]byte)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Bytes()
		r.mutex.Lock()
		r.bRead += uint64(len(b))
		r.mutex.Unlock()
		callback(append(b, utils.NewLineByte...))
	}

	return scanner.Err()
}

// ReadAll reads everything and dump it into one byte slice.
func (r *Reader) ReadAll() ([]byte, error) {
	b, err := ioutil.ReadAll(r)

	r.mutex.Lock()
	r.bRead = uint64(len(b))
	r.mutex.Unlock()

	return b, err
}

// ReadArray returns a data type-specific array returned via a callback function
func (r *Reader) ReadArray(callback func([]byte)) error {
	return readArray(r, callback)
}

// ReadMap returns a data type-specific key/values returned via a callback function
func (r *Reader) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	return readMap(r, config, callback)
}

// Write is a dummy function because it's a reader interface
func (r *Reader) Write(p []byte) (int, error) {
	return 0, errors.New("Cannot write to a reader interface")
}

// Writeln is a dummy function because it's a reader interface
func (r *Reader) Writeln(b []byte) (int, error) {
	return 0, errors.New("Cannot write to a reader interface")
}

// Open the stream.Io interface for another dependant
func (r *Reader) Open() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dependants++
}

// Close the stream.Io interface
func (r *Reader) Close() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.dependants--

	if r.dependants < 0 {
		panic("More closed dependants than open")
	}

	if r.dependants == 0 && r.readCloser != nil {
		r.readCloser.Close()
	}
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (r *Reader) WriteTo(w io.Writer) (int64, error) {
	return writeTo(r, w)
}

// GetDataType returns the murex data type for the stream.Io interface
func (r *Reader) GetDataType() (dt string) {
	for {
		r.dtLock.Lock()
		dt = r.dataType
		r.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

// SetDataType defines the murex data type for the stream.Io interface
func (r *Reader) SetDataType(dt string) {
	r.dtLock.Lock()
	r.dataType = dt
	r.dtLock.Unlock()
	return
}

// DefaultDataType defines the murex data type for the stream.Io interface if
// it's not already set
func (r *Reader) DefaultDataType(err bool) {
	r.dtLock.Lock()
	dt := r.dataType
	r.dtLock.Unlock()

	if dt == "" {
		if err {
			r.dtLock.Lock()
			r.dataType = types.Null
			r.dtLock.Unlock()
		} else {
			r.dtLock.Lock()
			r.dataType = types.Generic
			r.dtLock.Unlock()
		}
	}
}
