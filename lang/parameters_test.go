package lang_test

import (
	"testing"

	"github.com/lmorg/murex/test"
)

func TestParamHangBug(t *testing.T) {
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
			Block:  `out: @ FOO[BAR]`,
			Stdout: "@ FOO[BAR]\n",
		},
		{
			Block:  `out: @[out]`,
			Stdout: "@[out]\n",
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

	test.RunMurexTests(tests, t)
}
