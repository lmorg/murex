package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func TestRunTestPrivVarScoping(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foo/bar/foobar",
			TestBlock: `out $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				PreBlock:  "global foo=bar",
				PostBlock: "!global foo",
				Stdout:    "bar\n",
			},
		},
		{
			Function:  "foo/bar/foobar",
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

func TestRunTestPrivParameters(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foo/bar/foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters: []string{"a", "b", "c"},
				Stdout:     `["foobar","a","b","c"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foo/bar/foobar",
			TestBlock: "out $ARGS",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Parameters: []string{"1", "2", "3"},
				Stdout:     `["foobar","1","2","3"]` + utils.NewLineString,
			},
		},
		{
			Function:  "foo/bar/foobar",
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

func TestRunTestPrivDataTypes(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foo/bar/foobar",
			TestBlock: "tout json {}",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdout:   `{}`,
				StdoutDT: "json",
			},
		},
		{
			Function:  "foo/bar/foobar",
			TestBlock: "tout <err> json {}",
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stderr:   `{}`,
				StderrDT: "json",
			},
		},
		{
			Function:  "foo/bar/foobar",
			TestBlock: "tout notjson {}",
			Passed:    false,
			UTP: lang.UnitTestPlan{
				Stdout:   `{}`,
				StdoutDT: "json",
			},
		},
		{
			Function:  "foo/bar/foobar",
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

func TestRunTestPrivStdin(t *testing.T) {
	plans := []testUTPs{
		{
			Function:  "foo/bar/foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:  "bar",
				Stdout: "bar",
			},
		},
		{
			Function:  "foo/bar/foobar",
			TestBlock: `-> set foo; $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:    "bar",
				Stdout:   "bar",
				StdoutDT: types.Generic,
			},
		},
		{
			Function:  "foo/bar/foobar",
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
			Function:  "foo/bar/foobar",
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
