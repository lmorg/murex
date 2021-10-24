package escape

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/types/string"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestEscape(t *testing.T) {
	// as stdin

	test.RunMethodTest(
		t, cmdEscape, "escape",
		`hello world`, types.String, nil,
		`"hello world"`, nil)

	test.RunMethodTest(
		t, cmdEscape, "escape",
		`"hello world"`, types.String, nil,
		`"\"hello world\""`, nil)

	// as a parameter

	test.RunMethodTest(
		t, cmdEscape, "escape",
		``, types.String, []string{`hello`, `world`},
		`"hello world"`, nil)

	test.RunMethodTest(
		t, cmdEscape, "escape",
		``, types.String, []string{"hello world"},
		`"hello world"`, nil)

	test.RunMethodTest(
		t, cmdEscape, "escape",
		``, types.String, []string{`"hello world"`},
		`"\"hello world\""`, nil)
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

func TestHtml(t *testing.T) {
	// as stdin

	test.RunMethodTest(
		t, cmdHtml, "eschtml",
		`foo & bar`, types.String, nil,
		`foo &amp; bar`, nil)

	test.RunMethodTest(
		t, cmdHtml, "eschtml",
		`"foo & bar"`, types.String, nil,
		`&#34;foo &amp; bar&#34;`, nil)

	// as a parameter

	test.RunMethodTest(
		t, cmdHtml, "escape",
		``, types.String, []string{`foo`, `&`, `bar`},
		`foo &amp; bar`, nil)

	test.RunMethodTest(
		t, cmdHtml, "escape",
		``, types.String, []string{"foo & bar"},
		`foo &amp; bar`, nil)

	test.RunMethodTest(
		t, cmdHtml, "escape",
		``, types.String, []string{`"foo & bar"`},
		`&#34;foo &amp; bar&#34;`, nil)
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
	// as stdin

	test.RunMethodTest(
		t, cmdUrl, "escurl",
		`foo bar`, types.String, nil,
		`foo%20bar`, nil)

	// as a parameter

	test.RunMethodTest(
		t, cmdUrl, "escurl",
		``, types.String, []string{`foo`, `bar`},
		`foo%20bar`, nil)

	test.RunMethodTest(
		t, cmdUrl, "escurl",
		``, types.String, []string{"foo bar"},
		`foo%20bar`, nil)

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
	// as stdin

	test.RunMethodTest(
		t, cmdEscapeCli, "esccli",
		`foo bar`, types.String, nil,
		"foo\\ bar\n", nil)
}

func TestCli2(t *testing.T) {
	tests := []test.MurexTest{
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
			Stderr:  "Error in `!esccli` ( 1,1): exec: \"!esccli\": executable file not found in $PATH\n",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}
