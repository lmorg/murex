package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestMarshalArrayJsonString(t *testing.T) {
	count.Tests(t, 1)

	input := []string{"e", "d", "c", "b", "a"} // lets prove the output retains sorting
	output := `["e","d","c","b","a"]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	b, err := lang.MarshalData(fork.Process, types.Json, input)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != output {
		t.Error("Marshaller output doesn't match expected:")
		t.Logf("  Input:    %v", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%s'", b)
	}
}

func TestMarshalArrayJsonInt(t *testing.T) {
	count.Tests(t, 1)

	input := []int{5, 4, 3, 2, 1} // lets prove the output retains sorting
	output := `[5,4,3,2,1]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	b, err := lang.MarshalData(fork.Process, types.Json, input)
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != output {
		t.Error("Marshaller output doesn't match expected:")
		t.Logf("  Input:    %v", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%s'", b)
	}
}
