package shell

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/inject"
)

type testSyntaxCompletionsType struct {
	Input    string // Please note that the underscore character must
	Expected string // preceed the character who's cursor is positioned
}

func testSyntaxCompletions(t *testing.T, tests []testSyntaxCompletionsType) {
	t.Helper()
	count.Tests(t, len(tests))
	var failed int

	for i, test := range tests {
		if !strings.Contains(test.Input, "_") {
			t.Errorf("Invalid test. Test should contain an underscore character denoting cursor position")
			continue
		}
		pos := strings.Index(test.Input, "_") - 1
		input := strings.Replace(test.Input, "_", "", 1)

		r, newPos := syntaxCompletion([]rune(input), pos)
		newPos++
		output, err := inject.String(string(r), "_", newPos)
		if err != nil {
			t.Error("Cannot inject '_' into output string:")
			t.Logf("  Test #:  %d (%s)", i, t.Name())
			t.Logf("  Input:  '%s'", test.Input)
			t.Logf("  Error:   %s", err.Error())
			t.Logf("  String: '%s'", string(r))
			t.Logf("  Pos:     %d", newPos)
			failed++
			continue
		}

		if output != test.Expected {
			t.Error("Expected does not match output:")
			t.Logf("  Test #:    %d (%s)", i, t.Name())
			t.Logf("  Input:    '%s'", test.Input)
			t.Logf("  Block:    '%s'", input)
			t.Logf("  Pos:       %d", pos)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", output)
			failed++
		}
	}

	t.Logf("%d test(s) failed", failed)
}

func TestSyntaxCompletionsCurlyBrackets(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Input:    "func: param{_",
			Expected: "func: param{_}",
		},
		{
			Input:    "func: param{f_",
			Expected: "func: param{f_",
		},
		{
			Input:    "func: param{{_",
			Expected: "func: param{{_}",
		},
		{
			Input:    "func: param{{f_",
			Expected: "func: param{{f_",
		},
		{
			Input:    "func: param{{fo_",
			Expected: "func: param{{fo_",
		},
		{
			Input:    "func: param{_fo}}",
			Expected: "func: param{_fo}}",
		},
		{
			Input:    `func: param{_\}`,
			Expected: `func: param{_\}`,
		},
	}

	testSyntaxCompletions(t, tests)
}

func TestSyntaxCompletionsSquareBrackets(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Input:    "func: param[_",
			Expected: "func: param[_]",
		},
		{
			Input:    "func: param[f_",
			Expected: "func: param[f_",
		},
		{
			Input:    "func: param[[_",
			Expected: "func: param[[_]",
		},
		{
			Input:    "func: param[[f_",
			Expected: "func: param[[f_",
		},
		{
			Input:    "func: param[[fo_",
			Expected: "func: param[[fo_",
		},
		{
			Input:    "func: param[_fo]]",
			Expected: "func: param[_fo]]",
		},
		{
			Input:    `func: param[_\[`,
			Expected: `func: param[_\[`,
		},
		{
			Input:    `^[_`,
			Expected: `^[_]`,
		},
		{
			Input:    `^foobar[_`,
			Expected: `^foobar[_]`,
		},
		{
			Input:    `@[_`,
			Expected: `@[_]`,
		},
		{
			Input:    `echo @[_`,
			Expected: `echo @[_]`,
		},
		{
			Input:    `echo: @[_`,
			Expected: `echo: @[_]`,
		},
		{
			Input:    `$[_`,
			Expected: `$[_]`,
		},
		{
			Input:    `echo $[_`,
			Expected: `echo $[_]`,
		},
		{
			Input:    `echo: $[_`,
			Expected: `echo: $[_]`,
		},
	}

	testSyntaxCompletions(t, tests)
}

// https://github.com/lmorg/murex/issues/275
func TestSyntaxCompletionsMixedBrackets(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Input:    "func: param[{[_",
			Expected: "func: param[{[_",
		},
		{
			Input:    "func: param{[{_",
			Expected: "func: param{[{_}",
		},
	}

	testSyntaxCompletions(t, tests)
}
