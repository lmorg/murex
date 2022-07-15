package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParamHangBug(t *testing.T) {
	const config = `
		config: set proc strict-vars false;
		config: set proc strict-arrays false; 
	`

	tests := []test.MurexTest{
		// string
		{
			Block:  `out: $FOO{BAR}`,
			Stdout: "\n",
		},
		{
			Block:  `out: $ FOO{BAR}`,
			Stdout: "$ FOO{BAR}\n",
		},
		{
			Block:  `out: ${out}`,
			Stdout: "\n",
		},
		{
			Block:  `out: $FOO`,
			Stdout: "\n",
		},
		{
			Block:  `out: $FOO[BAR]`,
			Stdout: "\n",
		},
		{
			Block: `out: $FOO[[BAR]]`,
			Stdout: "\n",
			Stderr: "Error in `[[` ( 4,9): murex doesn't know how to lookup `[][]string` (please file a bug with on the murex Github page: https://lmorg/murex)\n",
		},
		{
			Block:  `out: $ FOO[BAR]`,
			Stdout: "$ FOO[BAR]\n",
		},
		{
			Block:  `out: $[out]`,
			Stdout: "$[out]\n",
		},
		{
			Block:  `out: $FOO(BAR)`,
			Stdout: "BAR\n",
		},
		{
			Block:  `out: $ FOO(BAR)`,
			Stdout: "$ FOO(BAR)\n",
		},
		{
			Block:  `out: $(out)`,
			Stdout: "$(out)\n",
		},
		// array
		{
			Block:  `out: @FOO{BAR}`,
			Stdout: "\n",
		},
		{
			Block:  `out: @ FOO{BAR}`,
			Stdout: "@ FOO{BAR}\n",
		},
		{
			Block:  `out: @{out}`,
			Stdout: "\n",
		},
		{
			Block:  `out: @FOO`,
			Stdout: "\n",
		},
		{
			Block:  `out: @FOO[BAR]`,
			Stdout: "\n",
		},
		{
			Block:  `out: @FOO[[BAR]]`,
			Stdout: "\n",
		},
		{
			Block:  `out: @ FOO[BAR]`,
			Stdout: "@ FOO[BAR]\n",
		},
		{
			Block:  `out: @ FOO[[BAR]]`,
			Stdout: "@ FOO[[BAR]]\n",
		},
		{
			Block:  `out: @[out]`,
			Stdout: "@[out]\n",
		},
		{
			Block:  `out: @[[out]]`,
			Stdout: "@[[out]]\n",
		},
		{
			Block:  `out: @FOO(BAR)`,
			Stdout: "BAR\n",
		},
		{
			Block:  `out: @ FOO(BAR)`,
			Stdout: "@ FOO(BAR)\n",
		},
		{
			Block:  `out: @(out)`,
			Stdout: "@(out)\n",
		},
	}

	for i := range tests {
		tests[i].Block = config + tests[i].Block
	}

	test.RunMurexTests(tests, t)
}
