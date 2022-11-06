package spellcheck

import (
	"testing"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestHighlighter(t *testing.T) {
	tests := []struct {
		Line     string
		Words    []string
		Expected string
	}{
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`quick`},
			Expected: `The _quick_ brown fox jumped over the lazy dog`,
		},
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`quick`, `the`},
			Expected: `The _quick_ brown fox jumped over _the_ lazy dog`,
		},
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`dog`},
			Expected: `The quick brown fox jumped over the lazy _dog_`,
		},
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`The`},
			Expected: `_The_ quick brown fox jumped over the lazy dog`,
		},
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`foobar`},
			Expected: `The quick brown fox jumped over the lazy dog`,
		},
		{
			Line:     `The quick brown fox jumped over the lazy dog`,
			Words:    []string{`own`},
			Expected: `The quick brown fox jumped over the lazy dog`,
		},
		///
		{
			Line:     `Hello, 世界, world`,
			Words:    []string{`Hello`},
			Expected: `_Hello_, 世界, world`,
		},
		{
			Line:     `Hello, 世界, world`,
			Words:    []string{`world`},
			Expected: `Hello, 世界, _world_`,
		},
		{
			Line:     `Hello, 世界, world`,
			Words:    []string{`世界`},
			Expected: `Hello, _世界_, world`,
		},
		{
			Line:     `Hello, 世界☺, world`,
			Words:    []string{`世界`},
			Expected: `Hello, _世界_☺, world`,
		},
		/// ignore single character terms
		{
			Line:     `Hello, 世 界☺, world`,
			Words:    []string{`世`},
			Expected: `Hello, 世 界☺, world`,
		},
		{
			Line:     `The quick brown fox jumped o ver the lazy dog`,
			Words:    []string{`o`, `ver`},
			Expected: `The quick brown fox jumped o _ver_ the lazy dog`,
		},
		{
			Line:     `The quick brown fox jumped ov er the lazy dog`,
			Words:    []string{`ov`, `er`},
			Expected: `The quick brown fox jumped _ov_ _er_ the lazy dog`,
		},
		///
		{
			Line:     `foo`,
			Words:    []string{`foo`},
			Expected: `_foo_`,
		},
		{
			Line:     `foo `,
			Words:    []string{`foo`},
			Expected: `_foo_ `,
		},
		{
			Line:     `foo f`,
			Words:    []string{`foo`},
			Expected: `_foo_ f`,
		},
		{
			Line:     `foo fo`,
			Words:    []string{`foo`},
			Expected: `_foo_ fo`,
		},
		{
			Line:     `foo foo`,
			Words:    []string{`foo`},
			Expected: `_foo_ _foo_`,
		},
		{
			Line:     `foo foob`,
			Words:    []string{`foo`},
			Expected: `_foo_ foob`,
		},
		{
			Line:     `foo foobar`,
			Words:    []string{`foo`},
			Expected: `_foo_ foobar`,
		},
		/// bug fix:
		{
			Line:     `am ammend`,
			Words:    []string{`ammend`},
			Expected: `am _ammend_`,
		},
	}

	count.Tests(t, len(tests))

	lang.InitEnv()
	lang.ShellProcess.Config.Define("shell", "color", config.Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc",
		//Default:     (runtime.GOOS != "windows" && isInteractive),
		Default:  true,
		DataType: types.Boolean,
		Global:   true,
	})

	var highlight = &highlightT{
		start: "_",
		end:   "_",
	}

	for i, test := range tests {
		actual := test.Line
		for _, word := range test.Words {
			highlighter(&actual, []rune(word), highlight)
		}

		if actual != test.Expected {
			t.Errorf("Mismatch in test %d", i)
			t.Logf("  Line:     '%s'", test.Line)
			t.Logf("  Words:     %s", json.LazyLogging(test.Words))
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}
