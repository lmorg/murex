package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseQuoteBlock(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "echo { %() }",
			Stdout: "{ %() }\n",
		},
		{
			Block:  "echo { %[] }",
			Stdout: "{ %[] }\n",
		},
		{
			Block:  "echo { %{} }",
			Stdout: "{ %{} }\n",
		},
	}

	test.RunMurexTests(tests, t)
}
