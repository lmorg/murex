package expressions

import (
	"testing"

	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type expTestT struct {
	input    string
	expected string
	pos      int
}

type expTestsT struct {
	tests  []expTestT
	symbol symbols.Exp
}

func getVarMock(name string) (interface{}, string, error) {
	return name, types.String, nil
}

func setVarMock(name string, value interface{}, dataType string) error {
	return nil
}

func testParserSymbol(t *testing.T, tests expTestsT) {
	t.Helper()

	count.Tests(t, len(tests.tests))

	for i, test := range tests.tests {
		tree := newExpTree([]rune(test.input))
		err := tree.parse(true)
		tree.getVar = getVarMock
		tree.setVar = setVarMock

		switch {
		case err != nil:
			t.Errorf("Error: %s", err.Error())
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			continue

		case len(tree.ast) == 0:
			t.Error("No ASTs generated:")
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			continue

		case tree.ast[0].key != tests.symbol:
			t.Error("Unexpected symbol:")

		case tree.ast[0].Value() != test.expected:
			t.Error("Expected doesn't match actual:")

		case tree.ast[0].pos != test.pos:
			t.Errorf("Pos doesn't match expected:")

		default:
			// success
			continue
		}

		t.Logf("  Test:        %d", i)
		t.Logf("  Expression: '%s'", test.input)
		t.Logf("  exp symbol: '%s'", tests.symbol.String())
		t.Logf("  act symbol: '%s'", tree.ast[0].key.String())
		t.Logf("  Expected:   '%s'", test.expected)
		t.Logf("  Actual:     '%s'", tree.ast[0].Value())
		t.Logf("  act bytes:  %v", tree.ast[0].value)
		t.Logf("  Character pos (exp: %d, act: %d)", test.pos, tree.ast[0].pos)
	}
}

type expressionTestT struct {
	Expression string
	Expected   interface{}
	Error      bool
}

func testExpression(t *testing.T, tests []expressionTestT) {
	t.Helper()

	count.Tests(t, len(tests))

	for i, test := range tests {
		tree := newExpTree([]rune(test.Expression))
		tree.getVar = getVarMock
		tree.setVar = setVarMock

		err := tree.parse(true)
		if err != nil {
			t.Errorf("Parser error in test %d: %s", i, err.Error())
		}
		dt, err := tree.execute()

		switch {
		case (err != nil) != test.Error:
			t.Error("err != nil:")

		case len(tree.ast) == 0:
			t.Error("Empty AST tree produced:")

		case dt != nil && dt.Value != test.Expected:
			t.Error("Result doesn't match expected:")

		default:
			// success
			continue
		}

		t.Logf("  Test:       %d", i)
		t.Logf("  Expression: '%s'", test.Expression)
		t.Logf("  Expected:   %s (%T)", json.LazyLogging(test.Expected), test.Expected)
		t.Logf("  Actual:     %v", json.LazyLogging(dt))
		t.Logf("  exp error:  %v", test.Error)
		t.Logf("  Error:      %v", err)
		t.Logf("  Dump():     %s", json.LazyLoggingPretty(tree.Dump()))
		t.Logf("  raw memory: %v", tree)
	}
}
