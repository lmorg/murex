package lang_test

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
				PreBlock:  "global foo=bar",
				PostBlock: "!global foo",
				Stdout:    "bar\n",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out $foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				PreBlock:  "set foo=bar",
				PostBlock: "!set foo",
				Stdout:    "bar\n",
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
				Parameters: []string{"a", "b", "c"},
				Stdout:     `["foobar","a","b","c"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters: []string{"1", "2", "3"},
				Stdout:     `["foobar","1","2","3"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters: []string{"foo bar"},
				Stdout:     `["foobar","foo bar"]` + utils.NewLineString,
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
				Stdout:   `{}`,
				StdoutDT: "json",
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout <err> json {}",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stderr:   `{}`,
				StderrDT: "json",
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout notjson {}",
			Passed:    false,
			UTP: lang.UnitTestPlan{
				Stdout:   `{}`,
				StdoutDT: "json",
			},
		},
		{
			Function:  "foobar",
			TestBlock: "tout <err> notjson {}",
			Passed:    false,
			UTP: lang.UnitTestPlan{
				Stderr:   `{}`,
				StderrDT: "json",
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
				Stdin:  "bar",
				Stdout: "bar",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:    "bar",
				Stdout:   "bar",
				StdoutDT: types.Generic,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:    "bar",
				StdinDT:  "notjson",
				Stdout:   "bar",
				StdoutDT: "notjson",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `-> set foo; out $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:    "bar",
				StdinDT:  "notjson",
				Stdout:   "bar\n",
				StdoutDT: types.String,
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
				Stderr:     "\n",
				ExitNumber: 0,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stderr:     "\n",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				Stderr:     "\n",
				ExitNumber: 2,
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
				StdoutRx: "(foo|bar)",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StdoutRx: "(FOO|BAR)",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRx: "foobar",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRx: "^foobar\n$",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StdoutRx: "^ob$",
			},
		},
		{
			Function:  "foobar",
			TestBlock: `out foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StdoutRx: "ob",
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
				StderrRx:   "(foo|bar)",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foo`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrRx:   "(FOO|BAR)",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRx:   "foobar",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRx:   "^foobar\n$",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    false,
			UTP: lang.UnitTestPlan{
				StderrRx:   "^ob$",
				ExitNumber: 1,
			},
		},
		{
			Function:  "foobar",
			TestBlock: `err foobar`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				StderrRx:   "ob",
				ExitNumber: 1,
			},
		},
	}

	testRunTest(t, plans)
}
