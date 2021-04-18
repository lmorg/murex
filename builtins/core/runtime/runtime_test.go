package cmdruntime_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	cmdruntime "github.com/lmorg/murex/builtins/core/runtime"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/json"
)

func marshalHelp(t *testing.T) string {
	t.Helper()

	b, err := json.Marshal(cmdruntime.Help(), false)
	if err != nil {
		t.Errorf("Cannot marshal help(): %s", err)
	}
	return string(b)
}

func TestRuntimeHelp(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `runtime --help`,
			Stdout: marshalHelp(t),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestRuntimeNoPanic(t *testing.T) {
	lang.InitEnv()

	tests := []test.MurexTest{
		{
			Block:  `runtime @{runtime --help}`,
			Stdout: `^.+$`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
