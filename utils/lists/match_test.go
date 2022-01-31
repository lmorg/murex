package lists_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/lists"
)

func TestMatch(t *testing.T) {
	count.Tests(t, 3)

	a := []string{"foo", "bar"}

	if !lists.Match(a, "foo") {
		t.Error("Did not match foo")
	}

	if !lists.Match(a, "bar") {
		t.Error("Did not match bar")
	}

	if lists.Match(a, "foobar") {
		t.Error("Incorrectly matched foobar")
	}
}
