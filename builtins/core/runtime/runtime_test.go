package cmdruntime

import (
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/utils/json"
)

func TestRangeByIndex(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `runtime --help`,
			Stdout: marshalHelp(t),
		},
	}

	test.RunMurexTests(tests, t)
}

func marshalHelp(t *testing.T) string {
	t.Helper()

	b, err := json.Marshal(help(), false)
	if err != nil {
		t.Errorf("Cannot marshal help(): %s", err)
	}
	return string(b)
}
