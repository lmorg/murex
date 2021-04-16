package shell

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/inject"
)

type testSyntaxCompletionsType struct {
	Line     string // Please note that the underscore character must precede the character who's cursor is positioned
	Change   string // The edit made to the line. Line should also include the edit
	Expected string // Please note that the underscore character must precede the character who's cursor is positioned
}

func testSyntaxCompletions(t *testing.T, tests []testSyntaxCompletionsType) {
	t.Helper()
	count.Tests(t, len(tests))
	var failed int

	for i, test := range tests {
		if !strings.Contains(test.Line, "_") {
			t.Errorf("Invalid test. Test should contain an underscore character denoting cursor position")
			continue
		}
		pos := strings.Index(test.Line, "_") - 1
		line := strings.Replace(test.Line, "_", "", 1)
		change := ansi.ForceExpandConsts(test.Change, true)

		r, newPos := syntaxCompletion([]rune(line), change, pos)
		newPos++
		output, err := inject.String(string(r), "_", newPos)
		if err != nil {
			t.Error("Cannot inject '_' into output string:")
			t.Logf("  Test #:  %d (%s)", i, t.Name())
			t.Logf("  Line:   '%s'", test.Line)
			t.Logf("  Change: '%s'", test.Change)
			t.Logf("  Error:   %s", err.Error())
			t.Logf("  String: '%s'", string(r))
			t.Logf("  Pos:     %d", newPos)
			failed++
			continue
		}

		if output != test.Expected {
			t.Error("Expected does not match output:")
			t.Logf("  Test #:    %d (%s)", i, t.Name())
			t.Logf("  Line:     '%s'", test.Line)
			t.Logf("  Change:   '%s'", test.Change)
			t.Logf("  Block:    '%s'", line)
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
			Line:     `func: param\{_`,
			Change:   "{",
			Expected: `func: param\{_`,
		},
		{
			Line:     "func: param{_",
			Change:   "{",
			Expected: "func: param{_}",
		},
		{
			Line:     "func: param{f_",
			Change:   "f",
			Expected: "func: param{f_",
		},
		{
			Line:     "func: param{{_",
			Change:   "{",
			Expected: "func: param{{_}",
		},
		{
			Line:     "func: param{{f_",
			Change:   "f",
			Expected: "func: param{{f_",
		},
		{
			Line:     "func: param{{fo_",
			Change:   "o",
			Expected: "func: param{{fo_",
		},
		{
			Line:     "func: param{_fo}}",
			Change:   "}",
			Expected: "func: param{_fo}}",
		},
		{
			Line:     `func: param{_\}`,
			Change:   "}",
			Expected: `func: param{_\}`,
		},
	}

	testSyntaxCompletions(t, tests)
}

func TestSyntaxCompletionsSquareBrackets(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Line:     "func: param[_",
			Change:   "[",
			Expected: "func: param[_]",
		},
		{
			Line:     "func: param[f_",
			Change:   "f",
			Expected: "func: param[f_",
		},
		{
			Line:     "func: param[[_",
			Change:   "[",
			Expected: "func: param[[_]",
		},
		{
			Line:     "func: param[[f_",
			Change:   "f",
			Expected: "func: param[[f_",
		},
		{
			Line:     "func: param[[fo_",
			Change:   "o",
			Expected: "func: param[[fo_",
		},
		{
			Line:     "func: param[_fo]]",
			Change:   "[",
			Expected: "func: param[_fo]]",
		},
		{
			Line:     `func: param[_\[`,
			Change:   "[",
			Expected: `func: param[_\[`,
		},
		{
			Line:     `^[_`,
			Change:   "[",
			Expected: `^[_]`,
		},
		{
			Line:     `^foobar[_`,
			Change:   "[",
			Expected: `^foobar[_]`,
		},
		{
			Line:     `@[_`,
			Change:   "[",
			Expected: `@[_]`,
		},
		{
			Line:     `echo @[_`,
			Change:   "[",
			Expected: `echo @[_]`,
		},
		{
			Line:     `echo: @[_`,
			Change:   "[",
			Expected: `echo: @[_]`,
		},
		{
			Line:     `$[_`,
			Change:   "[",
			Expected: `$[_]`,
		},
		{
			Line:     `echo $[_`,
			Change:   "[",
			Expected: `echo $[_]`,
		},
		{
			Line:     `echo: $[_`,
			Change:   "[",
			Expected: `echo: $[_]`,
		},
	}

	testSyntaxCompletions(t, tests)
}

func TestSyntaxCompletionsIndexElement(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Line:     "[_",
			Change:   "[",
			Expected: "[_]",
		},
		// I don't know why i need `]]]` for the test when readline only shows
		// `]]` -- there is another bug somewhere
		{
			Line:     "[[_]",
			Change:   "[",
			Expected: "[[ _ ]]]",
		},
		{
			Line:     "test -> [_",
			Change:   "[",
			Expected: "test -> [_]",
		},
		// I don't know why i need `]]]` for the test when readline only shows
		// `]]` -- there is another bug somewhere
		{
			Line:     "test -> [[_]",
			Change:   "[",
			Expected: "test -> [[ _ ]]]",
		},
	}

	testSyntaxCompletions(t, tests)
}

// https://github.com/lmorg/murex/issues/275
func TestSyntaxCompletionsMixedBrackets(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Line:     "func: param[{[_",
			Change:   "[",
			Expected: "func: param[{[_]",
		},
		{
			Line:     "func: param{[{_",
			Change:   "{",
			Expected: "func: param{[{_}",
		},
		{
			Line:     "func: param{_[{",
			Change:   "{",
			Expected: "func: param{_}[{",
		},
		{
			Line:     "func: param(_[{",
			Change:   "(",
			Expected: "func: param(_)[{",
		},
		{
			Line:     "func: param[_[{",
			Change:   "[",
			Expected: "func: param[_[{",
		},
		{
			Line:     "func: [{_]",
			Change:   "{",
			Expected: "func: [{_}]",
		},
		{
			Line:     "func: [(_]",
			Change:   "(",
			Expected: "func: [(_)]",
		},
		{
			Line:     "func: ({_)",
			Change:   "{",
			Expected: "func: ({_})",
		},
		{
			Line:     "func: ([_)",
			Change:   "[",
			Expected: "func: ([_])",
		},
		{
			Line:     "func: {[_}",
			Change:   "[",
			Expected: "func: {[_]}",
		},
		{
			Line:     "func: {(_}",
			Change:   "(",
			Expected: "func: {(_)}",
		},
	}

	testSyntaxCompletions(t, tests)
}

func TestSyntaxCompletionsQuoteBrace(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Line:     "echo: hello world_",
			Change:   "d",
			Expected: "echo: hello world_",
		},
		{
			Line:     "echo: (_hello world",
			Change:   "(",
			Expected: "echo: (_hello world)",
		},
		{
			Line:     "echo: (hello)_ world",
			Change:   ")",
			Expected: "echo: (hello)_ world",
		},
		{
			Line:     "echo: (hello)_ world)",
			Change:   ")",
			Expected: "echo: (hello)_ world",
		},
		{
			Line:     "echo: (hello_ world",
			Change:   "o",
			Expected: "echo: (hello_ world",
		},
		{
			Line:     "echo: ((_hello world",
			Change:   "(",
			Expected: "echo: ((_)hello world",
		},
		{
			Line:     "echo: ((_hello world)",
			Change:   "(",
			Expected: "echo: ((_)hello world)",
		},
		/*{
			Line:     "echo: _hello world)",
			Change:   "{BACKSPACE}",
			Expected: "echo: _hello world",
		},*/
	}

	testSyntaxCompletions(t, tests)
}

func TestSyntaxCompletionsQuotes(t *testing.T) {
	tests := []testSyntaxCompletionsType{
		{
			Line:     `'_`,
			Change:   `'`,
			Expected: `'_'`,
		},
		{
			Line:     `'h_`,
			Change:   `h`,
			Expected: `'h_`,
		},
		{
			Line:     `'_hello`,
			Change:   `'`,
			Expected: `'_hello'`,
		},
		{
			Line:     `echo: '_`,
			Change:   `'`,
			Expected: `echo: '_'`,
		},
		{
			Line:     `echo: 'h_`,
			Change:   `h`,
			Expected: `echo: 'h_`,
		},
		{
			Line:     `echo: '_hello`,
			Change:   `'`,
			Expected: `echo: '_hello'`,
		},
		/////
		{
			Line:     `"_`,
			Change:   `"`,
			Expected: `"_"`,
		},
		{
			Line:     `"h_`,
			Change:   `h`,
			Expected: `"h_`,
		},
		{
			Line:     `"_hello`,
			Change:   `"`,
			Expected: `"_hello"`,
		},
		{
			Line:     `echo: "_`,
			Change:   `"`,
			Expected: `echo: "_"`,
		},
		{
			Line:     `echo: "h_`,
			Change:   `h`,
			Expected: `echo: "h_`,
		},
		{
			Line:     `echo: "_hello`,
			Change:   `"`,
			Expected: `echo: "_hello"`,
		},
	}

	testSyntaxCompletions(t, tests)
}
