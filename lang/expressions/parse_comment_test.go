package expressions

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParseCommentInLineStatement(t *testing.T) {
	tests := []testParseStatementT{
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo/# this is an infixed comment #/ foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
		///
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: false,
		},
		{
			Statement: `echo /# this is an infixed comment #/foo bar`,
			Args: []string{
				"echo", "foo", "bar",
			},
			Exec: true,
		},
	}

	testParseStatement(t, tests)
}

func TestParseCommentInLine(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `3 /#* 4 #/+ 5`,
			Stdout: `8`,
		},
		{
			Block:  `/# foobar #/ out test`,
			Stdout: "test\n",
		},
		{
			Block:  `%[ 1 2 /#3#/ 4 ]`,
			Stdout: `[1,2,4]`,
		},
		{
			Block:  `%{ a:1, b:2, /#c:3#/ d:4 }`,
			Stdout: `{"a":1,"b":2,"d":4}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseCommentMultiLine(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				out 1.0
				/#
				out 2
				#/
				out 3
				`,
			Stdout: "1.0\n3\n",
		},
		{
			Block: `
				out 1.1
				/#out 2
				#/
				out 3
				`,
			Stdout: "1.1\n3\n",
		},
		{
			Block: `
				out 1.2
				/# out 2
				#/
				out 3
				`,
			Stdout: "1.2\n3\n",
		},
		{
			Block: `
				out 1.3
				/#
				out 2#/
				out 3
				`,
			Stdout: "1.3\n3\n",
		},
		{
			Block: `
				out 1.4
				/#
				out 2 #/
				out 3
				`,
			Stdout: "1.4\n3\n",
		},
		/////
		{
			Block: `
				out 1.5
				try {
					/# if {true} then {
						out 2
					}#/
					out 3
				}
				`,
			Stdout: "1.5\n3\n",
		},
		{
			Block: `
				out 1.6
				try {
					/#if {true} then {
						out 2
					} #/
					out 3
				}
				`,
			Stdout: "1.6\n3\n",
		},
		{
			Block: `
				out 1.7
				try {
					/# if {true} then {
						out 2
					} #/
					out 3
				}
				`,
			Stdout: "1.7\n3\n",
		},
		{
			Block: `
				out 1.8
				try {
					/#
					if {true} then {
						out 2
					} #/
					out 3
				}
				`,
			Stdout: "1.8\n3\n",
		},
		{
			Block: `
				out 1.9
				try {
					/#
					if {true} then {
						out 2
					}
					#/
					out 3
				}
				`,
			Stdout: "1.9\n3\n",
		},
	}

	test.RunMurexTests(tests, t)
}
