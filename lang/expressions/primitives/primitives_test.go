package primitives

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

// https://github.com/lmorg/murex/issues/845
func TestGenericsIssue845(t *testing.T) {
	count.Tests(t, 1)

	if DataType2Primitive(types.Generic) != String {
		t.Error("generics should be treated like strings")
	}
}
