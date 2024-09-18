package expressions

import (
	"testing"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/test"
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
	p.Name.Set("(test)")
	p.Config.Set("proc", "strict-vars", false, nil)
	p.Config.Set("proc", "strict-arrays", false, nil)

	for i, test := range tests.tests {
		tree := NewParser(p, []rune(test.input), 0)
		err := tree.parseExpression(true, true)

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
		tree := NewParser(p, []rune(test.input), 0)
		err := tree.parseExpression(true, true)
		var actVal string
		var failed bool

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
			failed = true

		case tree.ast[0].pos != test.pos:
			t.Errorf("Pos doesn't match expected:")
			failed = true

		default:
			v, err := tree.ast[0].dt.GetValue()
			if (err != nil) != test.error {
				t.Errorf("Error: %v", err)
				failed = true
			} else {
				actVal = json.LazyLogging(v.Value)
				if actVal != test.expected {
					t.Error("Expected doesn't match actual:")
					failed = true
				}
			}
		}

		if failed {
			t.Logf("  Test:        %d", i)
			t.Logf("  Expression: '%s'", test.input)
			t.Logf("  exp symbol: '%s'", tests.symbol.String())
			t.Logf("  act symbol: '%s'", tree.ast[0].key.String())
			t.Logf("  Expected:   '%s'", test.expected)
			t.Logf("  Actual:     '%s'", actVal)
			t.Logf("  Character pos (exp: %d, act: %d)", test.pos, tree.ast[0].pos)
		}
	}
}

type expressionTestT struct {
	Expression string
	Expected   any
	Error      bool
}

func testExpression(t *testing.T, tests []expressionTestT, strictTypes bool) {
	t.Helper()

	count.Tests(t, len(tests))

	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)
	p := lang.NewTestProcess()
	if err := p.Config.Set("proc", "strict-vars", false, nil); err != nil {
		panic(err)
	}
	if err := p.Config.Set("proc", "strict-arrays", false, nil); err != nil {
		panic(err)
	}
	if err := p.Config.Set("proc", "strict-types", strictTypes, nil); err != nil {
		panic(err)
	}

	for i, test := range tests {
		tree := NewParser(p, []rune(test.Expression), 0)

		err := tree.parseExpression(true, true)
		if err != nil {
			t.Errorf("Parser error in test %d: %s", i, err.Error())
			continue
		}
		dt, err := tree.executeExpr()
		var val *primitives.Value

		switch {
		default:
			// success
			continue

		case (err != nil) != test.Error:
			t.Error("tree.executeExpr() err != nil:")

		case len(tree.ast) == 0:
			t.Error("Empty AST tree produced:")

		case dt != nil:
			val, err = dt.GetValue()
			switch {
			default:
				// success
				continue
			case (err != nil) != test.Error:
				t.Error("dt.GetValue() err != nil:")
			case val.Value != test.Expected:
				t.Error("Result doesn't match expected:")
			}
		}

		t.Logf("  Test:       %d", i)
		t.Logf("  Expression: '%s'", test.Expression)
		t.Logf("  Expected:   %s (%T)", json.LazyLogging(test.Expected), test.Expected)
		t.Logf("  Actual:     %s (%T)", json.LazyLogging(val.Value), val.Value)
		t.Logf("  Struct:     %v", json.LazyLogging(val))
		t.Logf("  exp error:  %v", test.Error)
		t.Logf("  Error:      %v", err)
		t.Logf("  Dump():     %s", json.LazyLoggingPretty(tree.Dump()))
		t.Logf("  raw memory: %v", tree)
	}
}

func TestParseExprLogicalAnd(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "true == true && true == true",
			Stdout: "true",
		},
		{
			Block:   "true == false && true == true",
			Stdout:  "",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestParseExprPlusPlus(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "TestParseExprPlusPlus0 = 0; $TestParseExprPlusPlus0++; $TestParseExprPlusPlus0",
			Stdout: "^1$",
		},
		{
			Block:  "TestParseExprPlusPlus1 = 1; TestParseExprPlusPlus1++; $TestParseExprPlusPlus1",
			Stdout: "^1$",
			Stderr: "Error",
		},
		{
			Block:  "TestParseExprPlusPlus2 = 2; $TestParseExprPlusPlus2 ++; $TestParseExprPlusPlus2",
			Stdout: "^3$",
		},
		/*{ // doesn't compile
			Block:  "TestParseExprPlusPlus3 = 3; $TestParseExprPlusPlus3 ++ 4; $TestParseExprPlusPlus3",
			Stdout: "^3$",
			Stderr: "commands cannot begin with",
		},*/
	}

	test.RunMurexTestsRx(tests, t)
}

func TestParseExprMinusMinus(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  "TestParseExprMinusMinus0 = 0; $TestParseExprMinusMinus0--; $TestParseExprMinusMinus0",
			Stdout: "^-1$",
		},
		{
			Block:  "TestParseExprMinusMinus1 = 1; TestParseExprMinusMinus1--; $TestParseExprMinusMinus1",
			Stdout: "^1$",
			Stderr: "Error",
		},
		{
			Block:  "TestParseExprMinusMinus2 = 2; $TestParseExprMinusMinus2 --; $TestParseExprMinusMinus2",
			Stdout: "^1$",
		},
		/*{ // doesn't compile
			Block:  "TestParseExprMinusMinus3 = 3; $TestParseExprMinusMinus3 -- 4; $TestParseExprMinusMinus3",
			Stdout: "^7$",
		},*/
		{
			Block:  "TestParseExprMinusMinus4 = 4; $TestParseExprMinusMinus4 --5; $TestParseExprMinusMinus4",
			Stdout: "^94$",
		},
		{
			Block:  "TestParseExprMinusMinus5 = 5; $TestParseExprMinusMinus5--6; $TestParseExprMinusMinus5",
			Stdout: "^115$",
		},
	}

	test.RunMurexTestsRx(tests, t)
}
