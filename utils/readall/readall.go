package readall

import (
	"context"
	"io"

	"github.com/lmorg/murex/builtins/pipes/streams"
)

// ReadAll is a context aware equivalent to ioutils.ReadAll
func ReadAll(ctx context.Context, r io.Reader) (b []byte, err error) {
	w := streams.NewStdinWithContext(ctx, nil)

	_, err = w.ReadFrom(r)
	if err != nil {
		return
	}

	return w.ReadAll()
}

// WithCancel is a context aware equivalent to ioutils.ReadAll
func WithCancel(ctx context.Context, cancel context.CancelFunc, r io.Reader) (b []byte, err error) {
	w := streams.NewStdinWithContext(ctx, cancel)

	_, err = w.ReadFrom(r)
	if err != nil {
		return
	}

	return w.ReadAll()
}
