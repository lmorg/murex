package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

/*func TestParseVarsLogicalNotUnset(t *testing.T) {
	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)
	p := lang.NewTestProcess()
	if err := p.Config.Set("proc", "strict-vars", true, nil); err != nil {
		t.Fatal(err)
	}

	count.Tests(t, 2)

	var expectedErr string
	for _, expression := range []string{"$(murex_issue_1003_missing)", "!$(murex_issue_1003_missing)"} {
		tree := NewParser(p, []rune(expression), 0)
		if err := tree.parseExpression(true, true); err != nil {
			t.Fatalf("parseExpression(%q): %v", expression, err)
		}

		value, err := tree.executeExpr()
		if err != nil {
			t.Fatalf("executeExpr(%q): %v", expression, err)
		}
		if value == nil {
			t.Fatalf("executeExpr(%q) returned nil value", expression)
		}

		if _, err := value.GetValue(); err == nil {
			t.Errorf("GetValue(%q) returned no error for an unset variable", expression)
		} else if expectedErr == "" {
			expectedErr = err.Error()
		} else if err.Error() != expectedErr {
			t.Errorf("GetValue(%q) error = %q, want %q", expression, err, expectedErr)
		}
	}
}*/

func TestParseVarsScalarSymbol(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Scalar,
		tests: []expTestT{
			{
				input:    `$foo[1]`,
				expected: `$foo[1]`,
				pos:      -1,
			},
			{
				input:    `$foo[[/1]]`,
				expected: `$foo[[/1]]`,
				pos:      -1,
			},
			{
				input:    `!$(foo)`,
				expected: `!$(foo)`,
				pos:      -1,
			},
		},
	}

	testParserSymbol(t, tests)
}

/*func TestParseVarsArraySymbol(t *testing.T) {
	tests := expTestsT{
		symbol: symbols.Calculated,
		tests: []expTestT{
			{
				input:    `@foo[1]`,
				expected: `@foo[1]`,
				pos:      6,
			},
			{
				input:    `@foo[[/1]]`,
				expected: `@foo[[/1]]`,
				pos:      6,
			},
		},
	}

	testParserSymbol(t, tests)
}
*/
