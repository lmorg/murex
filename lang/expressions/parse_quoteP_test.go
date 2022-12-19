package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseQuoteParenthesis(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+))`,
			Args: []string{
				"regexp", "(f#/proc/[0-9]+/fd/([0-9]+))",
			},
			Pipes: []string{
				"!null",
			},
			Exec: false,
		},
		{
			Statement: `regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+))`,
			Args: []string{
				"regexp", "f#/proc/[0-9]+/fd/([0-9]+)",
			},
			Pipes: []string{
				"!null",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseQuoteParenthesisBlock(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `echo <!null> (f#/proc/[0-9]+/fd/([0-9]+))`,
			Stdout: "f#/proc/[0-9]+/fd/([0-9]+)\n",
		},
		{
			Block:  `echo <!null> /proc/0/fd/3 -> regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+)) -> match <!null> 3`,
			Stdout: "3\n",
		},
		{
			Block:  `echo (m/(NAME|PATTERN)/)`,
			Stdout: "m/(NAME|PATTERN)/\n",
		},
	}

	test.RunMurexTests(tests, t)
}
