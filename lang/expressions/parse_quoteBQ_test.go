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

// https://github.com/lmorg/murex/issues/697
func TestParseQuoteBlockIssue697(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "%[1] -> foreach var { out '%{foo}\n%{bar}' }",
			Stdout: "%{foo}\n%{bar}\n",
		},
	}

	test.RunMurexTests(tests, t)
}
