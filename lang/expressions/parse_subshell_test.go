package expressions_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestParseSubShellScalar(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `${out foobar}`,
			Stdout: `foobar`,
		},
		{
			Block:  `%[1 2 3 ${out foobar}]`,
			Stdout: `[1,2,3,"foobar"]`,
		},
		{
			Block:  `%[1 2 3 ${out foo bar}]`,
			Stdout: `[1,2,3,"foo bar"]`,
		},
		{
			Block:  `%{a: ${out foobar}}`,
			Stdout: `{"a":"foobar"}`,
		},
		{
			Block:  `%{a: ${out foo bar}}`,
			Stdout: `{"a":"foo bar"}`,
		},
		/////
		{
			Block:  `%[1 2 3 "${out foobar}"]`,
			Stdout: `[1,2,3,"foobar"]`,
		},
		{
			Block:  `%[1 2 3 "${out foo bar}"]`,
			Stdout: `[1,2,3,"foo bar"]`,
		},
		{
			Block:  `%{a: "${out foobar}"}`,
			Stdout: `{"a":"foobar"}`,
		},
		{
			Block:  `%{a: "${out foo bar}"}`,
			Stdout: `{"a":"foo bar"}`,
		},
		/////
		{
			Block:  `%[1 2 3 "-${out foobar}-"]`,
			Stdout: `[1,2,3,"-foobar-"]`,
		},
		{
			Block:  `%[1 2 3 "-${out foo bar}-"]`,
			Stdout: `[1,2,3,"-foo bar-"]`,
		},
		{
			Block:  `%{a: "-${out foobar}-"}`,
			Stdout: `{"a":"-foobar-"}`,
		},
		{
			Block:  `%{a: "-${out foo bar}-"}`,
			Stdout: `{"a":"-foo bar-"}`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseSubShellArray(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `@{ja: [1..3]}`,
			Stdout: `[1,2,3]`,
		},
		{
			Block:  `%[1 2 3 @{ja: [1..3]}]`,
			Stdout: `[1,2,3,1,2,3]`,
		},
		{
			Block:  `%[1 2 3 @{ja: [01..3]}]`,
			Stdout: `[1,2,3,"01","02","03"]`,
		},
		{
			Block:  `%{a: @{ja: [1..3]}}`,
			Stdout: `{"a":[1,2,3]}`,
		},
		{
			Block:  `%{a: @{ja: [01..3]}}`,
			Stdout: `{"a":["01","02","03"]}`,
		},
		/////
		{
			Block:  `%[1 2 3 "@{ja: [1..3]}"]`,
			Stdout: `[1,2,3,"@{ja: [1..3]}"]`,
		},
		{
			Block:  `%{a: "@{ja: [1..3]}"}`,
			Stdout: `{"a":"@{ja: [1..3]}"}`,
		},
	}

	test.RunMurexTests(tests, t)
}
