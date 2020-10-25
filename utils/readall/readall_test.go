package readall_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/readall"
)

func TestReadAll(t *testing.T) {
	count.Tests(t, 1)

	s := "foobar"

	buf := bytes.NewBuffer([]byte(s))
	ctx := context.Background()

	b, err := readall.ReadAll(ctx, buf)
	if err != nil {
		t.Error(err)
	}

	if string(b) != s {
		t.Errorf("`%s` != `%s`", string(b), s)
	}
}

func TestWithCancel(t *testing.T) {
	count.Tests(t, 1)

	s := "foobar"

	buf := bytes.NewBuffer([]byte(s))
	ctx := context.Background()
	f := func() {}

	b, err := readall.WithCancel(ctx, f, buf)
	if err != nil {
		t.Error(err)
	}

	if string(b) != s {
		t.Errorf("`%s` != `%s`", string(b), s)
	}
}
