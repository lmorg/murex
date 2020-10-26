package lang_test

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestRedirection(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   "out control test",
			Stdout:  "control test\n",
			Stderr:  "",
			ExitNum: 0,
		},
		{
			Block:   "out <err> redirect out",
			Stdout:  "",
			Stderr:  "redirect out\n",
			ExitNum: 0,
		},

		{
			Block:   "err control test",
			Stdout:  "",
			Stderr:  "control test\n",
			ExitNum: 1,
		},
		{
			Block:   "out <!out> redirect err",
			Stdout:  "redirect err\n",
			Stderr:  "",
			ExitNum: 0,
		},

		// null pipes

		{
			Block:   "out <null> null pipe",
			Stdout:  "",
			Stderr:  "",
			ExitNum: 0,
		},
		{
			Block:   "err <!null> null pipe",
			Stdout:  "",
			Stderr:  "",
			ExitNum: 1,
		},

		// pipelines

		{
			Block:   "regexp <!null> -> match ' '",
			Stdout:  "",
			Stderr:  "",
			ExitNum: 0,
		},
		/*{
			Block:   "regexp <!out> -> match ' '",
			Stdout:  "Error in `regexp` (0,1): `regexp` expects to be pipelined\n",
			Stderr:  "",
			ExitNum: 0,
		},*/
	}

	test.RunMurexTests(tests, t)
}

func TestRedirectionParserBug(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:   "$ARGS -> [ <!null> 10 ]",
			Stdout:  "",
			Stderr:  "",
			ExitNum: 1,
		},
	}

	test.RunMurexTests(tests, t)
}
