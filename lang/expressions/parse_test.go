package expressions

import (
	"testing"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type expTestT struct {
	input    string
	expected string
	pos      int
	error    bool
}

type expTestsT struct {
	tests  []expTestT
	symbol symbols.Exp
}

func testParserSymbol(t *testing.T, tests expTestsT) {
	t.Helper()

	count.Tests(t, len(tests.tests))

	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)
	p := lang.NewTestProcess()
	p.Config.Set("proc", "strict-vars", false, nil)
	p.Config.Set("proc", "strict-arrays", false, nil)

	for i, test := range tests.tests {
		tree := newExpTree(p, []rune(test.input))
		err := tree.parse(true)

		switch {
		case (err != nil) != test.error:
			t.Errorf("Error: %v", err)
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			continue

		case test.error:
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

func testParserObject(t *testing.T, tests expTestsT) {
	t.Helper()

	count.Tests(t, len(tests.tests))

	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)
	p := lang.NewTestProcess()
	p.Config.Set("proc", "strict-vars", false, nil)
	p.Config.Set("proc", "strict-arrays", false, nil)

	for i, test := range tests.tests {
		tree := newExpTree(p, []rune(test.input))
		err := tree.parse(true)

		switch {
		case (err != nil) != test.error:
			t.Errorf("Error: %v", err)
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			continue

		case test.error:
			continue

		case len(tree.ast) == 0:
			t.Error("No ASTs generated:")
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			continue

		case tree.ast[0].key != tests.symbol:
			t.Error("Unexpected symbol:")

		case json.LazyLogging(tree.ast[0].dt.Value) != test.expected:
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
		t.Logf("  Actual:     '%s'", json.LazyLogging(tree.ast[0].dt.Value))
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

	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)
	p := lang.NewTestProcess()
	p.Config.Set("proc", "strict-vars", false, nil)
	p.Config.Set("proc", "strict-arrays", false, nil)

	for i, test := range tests {
		tree := newExpTree(p, []rune(test.Expression))

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
