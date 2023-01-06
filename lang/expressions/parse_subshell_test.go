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

func TestParseSubShellScalarParam(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `echo ${out foobar}`,
			Stdout: "foobar\n",
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

func TestParseSubShellArrayParams(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `echo @{ja: [1..3]}`,
			Stdout: "1 2 3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseSubShellEmptyArrayBugFix(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				config: set proc strict-arrays false
				echo: @{g: <!null> pseudo-random-string-that-will-not-be-matched-with-anything}
			`,
			Stdout: "\n",
		},
		{
			Block: `
				config: set proc strict-arrays false
				bob = @{g: <!null> pseudo-random-string-that-will-not-be-matched-with-anything}
				$bob
			`,
			Stdout: "null",
		},
		{
			Block: `
				config: set proc strict-arrays false
				bob = @{g: <!null> pseudo-random-string-that-will-not-be-matched-with-anything}
				echo bob = $bob
			`,
			Stdout: "bob = null\n",
		},
		{
			Block: `
				config: set proc strict-arrays false
				bob = @{g: <!null> pseudo-random-string-that-will-not-be-matched-with-anything}
				echo bob = @bob
			`,
			Stdout: "bob =\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseSubShellArrayJson(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `echo @{tout json [1,2,3] }`,
			Stdout: "1 2 3\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseSubShellBugFixJsonStr(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `TestParseSubShellBugFixJsonStr0 = %[1 2 3]; echo @{ $TestParseSubShellBugFixJsonStr0 }`,
			Stdout: "1 2 3\n",
		},
		{
			Block:  `TestParseSubShellBugFixJsonStr1 = %{a:1, b:2, c:[1 2 3]}; echo @{ $TestParseSubShellBugFixJsonStr1[c] }`,
			Stdout: "1 2 3\n",
		},
	}

	test.RunMurexTests(tests, t)
}
