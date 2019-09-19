package define_test

import (
	"fmt"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/test/count"
)

func TestUnmarshalArrayJsonString(t *testing.T) {
	count.Tests(t, 1)

	input := `["e","d","c","b","a"]` // lets prove the output retains sorting
	output := `[e d c b a]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_CREATE_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	_, err := fork.Stdin.Write([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	v, err := define.UnmarshalData(fork.Process, types.Json)
	if err != nil {
		t.Error(err)
		return
	}

	if fmt.Sprintf("%v", v) != output {
		t.Error("Unmarshaller output doesn't match expected:")
		t.Logf("  Input:    %s", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%v'", v)
	}
}

func TestUnmarshalArrayJsonInt(t *testing.T) {
	count.Tests(t, 1)

	input := `[5,4,3,2,1]` // lets prove the output retains sorting
	output := `[5 4 3 2 1]`

	lang.InitEnv()
	fork := lang.ShellProcess.Fork(lang.F_CREATE_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)

	_, err := fork.Stdin.Write([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	v, err := define.UnmarshalData(fork.Process, types.Json)
	if err != nil {
		t.Error(err)
		return
	}

	if fmt.Sprintf("%v", v) != output {
		t.Error("Unmarshaller output doesn't match expected:")
		t.Logf("  Input:    %s", input)
		t.Logf("  Expected: '%s'", output)
		t.Logf("  Actual:   '%v'", v)
	}
}
