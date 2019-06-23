package null

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestNullArrayWriter(t *testing.T) {
	test.ArrayWriterTest(t, types.Null, "")
}
