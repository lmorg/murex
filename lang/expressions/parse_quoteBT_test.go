package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseQuoteBackTicks(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "echo `foobar`",
			Stdout: "'foobar'\n",
		},
	}

	test.RunMurexTests(tests, t)
}
