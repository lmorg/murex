package mail

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/smtp"
	"sync"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

// Since I don't want you to create null pipes, lets not regester it
func init() {
	stdio.RegisterPipe("mail", NewMail)
}

// Mail is null interface for named pipes
type Mail struct {
	mutex      sync.Mutex
	bWritten   uint64
	dependants int
	client     *smtp.Client
	body       io.WriteCloser
	bodyAuth   *bytes.Buffer
	bAuth      []byte
	useAuth    bool
	smtpAuth   smtp.Auth
	subject    string
	recipients []string
}

// Read is an empty method because you cannot read a sent email
func (m *Mail) Read(p []byte) (int, error) {
	return 0, nil
}

// ReadLine is an empty method because you cannot read a sent email
func (m *Mail) ReadLine(func([]byte)) error {
	return errors.New("ReadLine() is not supported by mail pipes")
}

// ReadArray is an empty method because you cannot read a sent email
func (m *Mail) ReadArray(context.Context, func([]byte)) error {
	return errors.New("ReadArray() is not supported by mail pipes")
}

// ReadArrayWithType is an empty method because you cannot read a sent email
func (m *Mail) ReadArrayWithType(context.Context, func(interface{}, string)) error {
	return errors.New("ReadArrayWithType() is not supported by mail pipes")
}

// ReadMap is an empty method because you cannot read a sent email
func (m *Mail) ReadMap(*config.Config, func(*stdio.Map)) error {
	return errors.New("ReadMap() is not supported by mail pipes")
}

// ReadAll is an empty method because you cannot read a sent email
func (m *Mail) ReadAll() ([]byte, error) {
	return nil, nil
}

// WriteTo is an empty method because you cannot read a sent email
func (m *Mail) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// Write - caches the data before send
func (m *Mail) Write(b []byte) (i int, err error) {
	m.mutex.Lock()
	if m.useAuth {
		i, err = m.bodyAuth.Write(b)
	} else {
		i, err = m.body.Write(b)
	}
	m.bWritten += uint64(i)
	m.mutex.Unlock()
	return i, err
}

// Writeln - caches the data before send
func (m *Mail) Writeln(b []byte) (int, error) {
	return m.Write(append(b, '\n'))
}

// WriteArray - caches the data before send
func (m *Mail) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(m, dataType)
}

// Stats - return data size
func (m *Mail) Stats() (bytesWritten, bytesRead uint64) {
	m.mutex.Lock()
	bytesWritten = m.bWritten
	//bytesRead = m.bRead
	m.mutex.Unlock()
	return
}

// GetDataType returns null because you cannot read a sent email
func (m *Mail) GetDataType() string { return types.Null }

// SetDataType - not required as data is emailed
func (m *Mail) SetDataType(string) {}

// DefaultDataType - not required as data is emailed
func (m *Mail) DefaultDataType(bool) {}

// IsTTY - Stdio.Io is an email not a terminal
func (m *Mail) IsTTY() bool { return false }

// Open email for sending
func (m *Mail) Open() {
	m.mutex.Lock()
	m.dependants++
	m.mutex.Unlock()
}

// Close email and send
func (m *Mail) Close() {
	m.mutex.Lock()

	m.dependants--
	if m.dependants == 0 {
		m.send()
	}

	if m.dependants < 0 {
		panic("more closed dependants than open")
	}

	m.mutex.Unlock()
}

// ForceClose is not required on this occasion
func (m *Mail) ForceClose() {}
