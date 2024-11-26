package lang

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestGetExtType(t *testing.T) {
	count.Tests(t, 2)

	fileExts = map[string]string{"foo": "bar"}

	if GetExtType("FOO") != "bar" {
		t.Error("foo != bar")
	}

	if GetExtType("l;dskjforijf;sdj;oweirnfifunweodijn") != types.Generic {
		t.Error("Undefined ext != generic")
	}
}
