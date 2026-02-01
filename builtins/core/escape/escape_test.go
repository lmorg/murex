package escape_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestEscape(t *testing.T) {
	tests := []test.MurexTest{
		// method

		{
			Block:  `tout str hello world -> escape`,
			Stdout: `"hello world"`,
		},
		{
			Block:  `tout str '"hello world"' -> escape`,
			Stdout: `"\"hello world\""`,
		},

		// parameter

		{
			Block:  `escape hello world`,
			Stdout: `"hello world"`,
		},
		{
			Block:  `escape "hello world"`,
			Stdout: `"hello world"`,
		},
		{
			Block:  `escape '"hello world"'`,
			Stdout: `"\"hello world\""`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestEscapeBang(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `!escape "hello world"`,
			Stdout: `hello world`,
		},
		{
			Block:  `!escape ("hello world")`,
			Stdout: `hello world`,
		},
		{
			Block:  `!escape '\"hello world\"'`,
			Stdout: `\"hello world\"`,
		},
		{
			Block:  `!escape '\"hello\ world\"'`,
			Stdout: `\"hello\ world\"`,
		},
		{
			Block:  `!escape "foo &amp bar"`,
			Stdout: `foo & bar`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestHTML(t *testing.T) {
	tests := []test.MurexTest{
		// method

		{
			Block:  `tout str foo & bar -> eschtml`,
			Stdout: `foo &amp; bar`,
		},
		{
			Block:  `tout str '"foo & bar"' -> eschtml`,
			Stdout: `&#34;foo &amp; bar&#34;`,
		},

		// parameter

		{
			Block:  `eschtml foo & bar`,
			Stdout: `foo &amp; bar`,
		},
		{
			Block:  `eschtml "foo & bar"`,
			Stdout: `foo &amp; bar`,
		},
		{
			Block:  `eschtml '"foo & bar"'`,
			Stdout: `&#34;foo &amp; bar&#34;`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestHtmlBang(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `!eschtml "foo &amp bar"`,
			Stdout: `foo & bar`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestUrl(t *testing.T) {
	tests := []test.MurexTest{
		// method

		{
			Block:  `tout str foo bar -> escurl`,
			Stdout: `foo%20bar`,
		},

		// parameter

		{
			Block:  `escurl foo bar`,
			Stdout: `foo%20bar`,
		},
		{
			Block:  `escurl "foo bar"`,
			Stdout: `foo%20bar`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestUrlBang(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `!escurl foo%20bar`,
			Stdout: `foo bar`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestCli(t *testing.T) {
	tests := []test.MurexTest{
		// method

		{
			Block:  `tout str foo bar -> esccli`,
			Stdout: "foo\\ bar\n",
		},
		{
			Block:  `tout str '"foo bar"' -> esccli`,
			Stdout: "\\\"foo\\ bar\\\"\n",
		},

		// parameter

		{
			Block:  `esccli foo bar`,
			Stdout: "foo bar\n",
		},
		{
			Block:  `esccli "foo bar"`,
			Stdout: "foo\\ bar\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestCliBang(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   `!esccli foo\ bar`,
			Stderr:  "executable file not found",
			ExitNum: 1,
		},
	}

	test.RunMurexTestsRx(tests, t)
}
