package null

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestArrayWriter(t *testing.T) {
	input := []string{"foo", "bar"}
	test.ArrayWriterTest(t, types.Null, input, "")
}
