package lang_test

/*
	This test library relates to using the Go testing framework to test murex's
	framework for unit testing shell scripts.

	The naming convention here is basically the inverse of Go's test naming
	convention. ie Go source files will be named "test_unit.go" (because
	calling it unit_test.go would mean it's a Go test rather than murex test)
	and the code is named UnitTestPlans (etc) rather than TestUnitPlans (etc)
	because the latter might suggest they would be used by `go test`. This
	naming convention is a little counterintuitive but it at least avoids
	naming conflicts with `go test`.
*/

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func TestRunTestVarScoping(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: `out $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				PreBlock:    "global foo=bar",
				PostBlock:   "!global foo",
				StdoutMatch: "bar\n",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out $foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				PreBlock:    "set foo=bar",
				PostBlock:   "!set foo",
				StdoutMatch: "bar\n",
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestParameters(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters:  []string{"a", "b", "c"},
				StdoutMatch: `["foobar","a","b","c"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters:  []string{"1", "2", "3"},
				StdoutMatch: `["foobar","1","2","3"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters:  []string{"foo bar"},
				StdoutMatch: `["foobar","foo bar"]` + utils.NewLineString,
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestDataTypes(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: "tout json {}",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutMatch: `{}`,
				StdoutType:  types.Json,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout <err> json {}",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrMatch: `{}`,
				StderrType:  types.Json,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout notjson {}",
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StdoutMatch: `{}`,
				StdoutType:  types.Json,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout <err> notjson {}",
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrMatch: `{}`,
				StderrType:  types.Json,
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestStdin(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:       "bar",
				StdoutMatch: "bar",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:       "bar",
				StdoutMatch: "bar",
				StdoutType:  types.Generic,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:       "bar",
				StdinType:   "notjson",
				StdoutMatch: "bar",
				StdoutType:  "notjson",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; out $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:       "bar",
				StdinType:   "notjson",
				StdoutMatch: "bar\n",
				StdoutType:  types.String,
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestExitNumber(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: `err`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrMatch: "\n",
				ExitNum:     0,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrMatch: "\n",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrMatch: "\n",
				ExitNum:     2,
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestRegexStdout(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: `out foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "(foo|bar)",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "(FOO|BAR)",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "foobar",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "^foobar\n$",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "^ob$",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRegex: "ob",
			},
		},
	}

	testRunTest(t, plans)
}

func TestRunTestRegexStderr(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foobar",
			TestBlock: `err foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRegex: "(foo|bar)",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrRegex: "(FOO|BAR)",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRegex: "foobar",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRegex: "^foobar\n$",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrRegex: "^ob$",
				ExitNum:     1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRegex: "ob",
				ExitNum:     1,
			},
		},
	}

	testRunTest(t, plans)
}
