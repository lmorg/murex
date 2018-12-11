package streams

import (
	"io"
)

// ReadCloser is a wrapper around an io.ReadCloser interface
type ReadCloser struct {
	Reader
}

// NewReadCloser creates a new Stdio.Io interface wrapper around a io.ReadCloser interface
func NewReadCloser(reader io.ReadCloser) (r *ReadCloser) {
	if reader == nil {
		panic("streams.ReadCloser interface has nil reader interface")
	}

	r = new(ReadCloser)
	r.reader = reader
	r.readCloser = reader
	return
}
