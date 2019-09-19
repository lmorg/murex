package history

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestNoColon(t *testing.T) {
	tests := []string{
		"command param1 param2 param3",
		"command: param1 param2 param3",
		"command : param1 param2 param3",
		"command :param1 param2 param3",
		"comman:d param1 param2 param3",
		"command param1: param2 param3",
		"command param1: param2: param3",
		"command param1: param2: param3:",
		"command: param1: param2: param3:",
		":command param1 param2 param3",
		":command: param1 param2 param3",
	}

	expected := []string{
		"command param1 param2 param3",
		"command param1 param2 param3",
		"command : param1 param2 param3",
		"command :param1 param2 param3",
		"comman d param1 param2 param3",
		"command param1: param2 param3",
		"command param1: param2: param3",
		"command param1: param2: param3:",
		"command param1: param2: param3:",
		"command param1 param2 param3",
		"command param1 param2 param3",
	}

	count.Tests(t, len(tests))

	for i := range tests {
		actual := noColon(tests[i])
		if actual != expected[i] {
			t.Errorf("Output does not match expected in test %d:", i)
			t.Log("  Original:", tests[i])
			t.Log("  Expected:", expected[i])
			t.Log("  Actual:  ", actual)
		}
	}
}
