package expressions_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/home"
	"github.com/lmorg/murex/utils/json"
)

type testParseSwitchT struct {
	Expression string
	Expected   []testParseSwitchNodeT
	Error      bool
}

type testParseSwitchNodeT struct {
	Condition  string
	Parameters []string
}

func (t *testParseSwitchT) Json() string { return json.LazyLoggingPretty(t.Expected) }

func runTestParseSwitch(t *testing.T, tests []testParseSwitchT) {
	t.Helper()

	count.Tests(t, len(tests))

	lang.InitEnv()
	defaults.Config(lang.ShellProcess.Config, false)

	for i, test := range tests {
		p := lang.NewTestProcess()
		p.Name.Set(t.Name())
		defaults.Config(p.Config, false)

		if err := p.Config.Set("proc", "strict-vars", false, nil); err != nil {
			panic(err)
		}
		if err := p.Config.Set("proc", "strict-arrays", false, nil); err != nil {
			panic(err)
		}
		expSwitch, err := expressions.ParseSwitch(p, []rune(test.Expression))

		if (err != nil) != test.Error {
			t.Errorf("Unexpected mismatch in errors in test %d", i)
			t.Logf("  Expression: '%s'", test.Expression)
			t.Logf("  Expected:   %s", test.Json())
			t.Logf("  err exp:    %v", test.Error)
			t.Logf("  err act:    %v", err)
			continue
		}

		var actual testParseSwitchT
		for j := range expSwitch {
			node := testParseSwitchNodeT{
				Condition:  expSwitch[j].Condition,
				Parameters: expSwitch[j].ParametersAll(),
			}
			actual.Expected = append(actual.Expected, node)
		}

		if test.Json() != actual.Json() {
			t.Errorf("Expected does not match actual in test %d", i)
			t.Logf("  Expression: '%s'", test.Expression)
			t.Logf("  Expected:   %s", test.Json())
			t.Logf("  Actual:     %s", actual.Json())
			t.Logf("  err exp:    %v", test.Error)
			t.Logf("  err act:    %v", err)
		}
	}
}

func TestParseSwitchBasic(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case "foo" { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo", "{ bar }"},
				},
			},
		},
		{
			Expression: `if "foo" { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "if",
					Parameters: []string{"foo", "{ bar }"},
				},
			},
		},
		{
			Expression: `case "foo" then { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo", "then", "{ bar }"},
				},
			},
		},
		{
			Expression: `if "foo" then { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "if",
					Parameters: []string{"foo", "then", "{ bar }"},
				},
			},
		},
		{
			Expression: `default { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "default",
					Parameters: []string{"{ bar }"},
				},
			},
		},
		{
			Expression: `if:"foo" {bar}`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "if",
					Parameters: []string{"foo", "{bar}"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchQuotes(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case 'foo'`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo"},
				},
			},
		},
		{
			Expression: `case "foo"`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo"},
				},
			},
		},
		{
			Expression: `case (foo)`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo"},
				},
			},
		},
		{
			Expression: `case (f(o)o)`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"f(o)o"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchMultiStatement(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `if "a" { b }; case "c" { d }; default { e }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "if",
					Parameters: []string{"a", "{ b }"},
				},
				{
					Condition:  "case",
					Parameters: []string{"c", "{ d }"},
				},
				{
					Condition:  "default",
					Parameters: []string{"{ e }"},
				},
			},
		},
		{
			Expression: `
				if "a" { b }
				case "c" { d }
				default { e }
			`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "if",
					Parameters: []string{"a", "{ b }"},
				},
				{
					Condition:  "case",
					Parameters: []string{"c", "{ d }"},
				},
				{
					Condition:  "default",
					Parameters: []string{"{ e }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchEscape(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case \s { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{" ", "{ bar }"},
				},
			},
		},
		{
			Expression: `case \t { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"\t", "{ bar }"},
				},
			},
		},
		{
			Expression: `case \r { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"\r", "{ bar }"},
				},
			},
		},
		{
			Expression: `case \n { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"\n", "{ bar }"},
				},
			},
		},
		{
			Expression: `case \q { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"q", "{ bar }"},
				},
			},
		},
		{
			Expression: `case \\ { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"\\", "{ bar }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchComments(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `
				case 1 { a }
				# case 2 { b }
				case 3 { c }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"1", "{ a }"},
				},
				{
					Condition:  "case",
					Parameters: []string{"3", "{ c }"},
				},
			},
		},
		{
			Expression: `
				case 1 { a }
				/# case 2 { b } #/
				case 3 { c }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"1", "{ a }"},
				},
				{
					Condition:  "case",
					Parameters: []string{"3", "{ c }"},
				},
			},
		},
		{
			Expression: `case 1 /#2#/ 3 { a }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"1", "3", "{ a }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchVariables(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case $foo { $bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"", "{ $bar }"},
				},
			},
		},
		{
			Expression: `case @foo { $bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"{ $bar }"},
				},
			},
		},
		{
			Expression: `case ~ { ~ }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{home.MyDir, "{ ~ }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchSubshells(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case ${out foo} { $bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo", "{ $bar }"},
				},
			},
		},
		{
			Expression: `case @{ja: [1..3]} { ~ }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"1", "2", "3", "{ ~ }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}

func TestParseSwitchCreators(t *testing.T) {
	tests := []testParseSwitchT{
		{
			Expression: `case %(foo) { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"foo", "{ bar }"},
				},
			},
		},
		{
			Expression: `case %[1 2 3] { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{"[1,2,3]", "{ bar }"},
				},
			},
		},
		{
			Expression: `case %{a:1, b:2, c:3} { bar }`,
			Expected: []testParseSwitchNodeT{
				{
					Condition:  "case",
					Parameters: []string{`{"a":1,"b":2,"c":3}`, "{ bar }"},
				},
			},
		},
	}

	runTestParseSwitch(t, tests)
}
