package lang_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

type testUTPs struct {
	Function  string
	TestBlock string
	Passed    bool
	UTP       lang.UnitTestPlan
}

func TestRunTestPositive(t *testing.T) {
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
		/*{
			Function:  "foobar",
			TestBlock: `<stdin> -> set foo; out $foo`,
			Passed:    true,
			UTP: lang.UnitTestPlan{
				Stdin:     `barrr`,
				PreBlock:  "set foo=bar",
				PostBlock: "!set foo",
				Stdout:    "barrr\n",
			},
		},*/
	}

	testRunTest(t, plans)
}

func testRunTest(t *testing.T, plans []testUTPs) {
	count.Tests(t, len(plans), "testRunTest")

	lang.InitEnv()

	for i := range plans {
		fileRef := &ref.File{
			Source: &ref.Source{
				Filename: "foobar.mx",
				Module:   "foo/bar",
				DateTime: time.Now(),
			},
		}

		lang.MxFunctions.Define(plans[i].Function, []rune(plans[i].TestBlock), fileRef)

		ut := new(lang.UnitTests)
		ut.Add(plans[i].Function, &plans[i].UTP, fileRef)

		if ut.Run(lang.ShellProcess.Tests, plans[i].Function) != plans[i].Passed {
			if plans[i].Passed {
				t.Errorf("Unit test %d failed", i)
				b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
				if err != nil {
					panic(err)
				}
				t.Logf("%s", b)
			} else {
				t.Errorf("Unit test %d passed when expected to fail", i)
			}
		}
	}
}

/*func TestRunTestScopePrivate(t *testing.T) {
	count.Tests(t, 1, "TestRunTestScopePrivate")

	const (
		preBlock  = `set foo=bar`
		testBlock = `out $foo`
		postBlock = `!set foo`
		function  = `foobar`
	)

	lang.InitEnv()

	fileRef := &ref.File{
		Source: &ref.Source{
			Filename: "foobar.mx",
			Module:   "foo/bar",
			DateTime: time.Now(),
		},
	}

	lang.MxFunctions.Define(function, []rune(testBlock), fileRef)

	test := &lang.UnitTestPlan{
		Stdout:    "bar" + utils.NewLineString,
		PreBlock:  preBlock,
		PostBlock: postBlock,
	}

	ut := new(lang.UnitTests)
	ut.Add(function, test, fileRef)

	if !ut.Run(function) {
		t.Error("ut.Run() == false")
	}
}
*/
