package null_test

import (
	"testing"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/test/count"
)

func TestNull(t *testing.T) {
	count.Tests(t, 5)

	n := new(null.Null)

	n.Open()

	i, err := n.Writeln([]byte("foobar"))
	if i != 6 {
		t.Errorf("i should be 6: %d", i)
	}
	if err != nil {
		t.Error(err)
	}

	b, err := n.ReadAll()
	if err != nil {
		t.Error(err)
	}
	if len(b) > 0 {
		t.Errorf("ReadAll should be empty: %s", string(b))
	}
	if b == nil {
		t.Error("b == nil. It should be an empty slice")
	}

	n.Close()
}
