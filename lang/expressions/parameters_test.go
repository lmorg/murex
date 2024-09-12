package expressions_test

import (
	"fmt"
	"math/rand"
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
			Stdout: "{BAR}\n",
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
			Stdout: "(BAR)\n",
		},
		{
			Block:  `out: $(out)`,
			Stdout: "\n",
		},
		// array
		{
			Block:  `out: @FOO{BAR}`,
			Stdout: "{BAR}\n",
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
			Stdout: "(BAR)\n",
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

func TestParamVarRange(t *testing.T) {
	test1, test2, test3 := rand.Int(), rand.Int(), rand.Int()

	tests := []test.MurexTest{
		{
			Block: fmt.Sprintf(`
				a: [Monday..Friday] -> set: %s%d
				out: @%s%d[2..]
				`,
				t.Name(), test1, t.Name(), test1,
			),
			Stdout: "^Wednesday Thursday Friday\n$",
		},
		{
			Block: fmt.Sprintf(`
				ja: [Monday..Friday] -> set: %s%d
				out: @%s%d[2..3]
				`,
				t.Name(), test2, t.Name(), test2,
			),
			Stdout: "^Wednesday Thursday\n$",
		},
		{
			Block: fmt.Sprintf(`
				ja: [Monday..Friday] -> set: %s%d
				out: @%s%d[..1]
				`,
				t.Name(), test3, t.Name(), test3,
			),
			Stdout: "^Monday Tuesday\n$",
		},
	}

	test.RunMurexTestsRx(tests, t)
}
