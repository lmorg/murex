package lang_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

func TestRunTestNotATest(t *testing.T) {
	count.Tests(t, 1, "TestRunTestNotATest")

	lang.InitEnv()

	fileRef := &ref.File{
		Source: &ref.Source{
			Filename: "foobar.mx",
			Module:   "foo/bar",
			DateTime: time.Now(),
		},
	}

	plan := &lang.UnitTestPlan{
		Stdout: "foobar\n",
	}

	lang.MxFunctions.Define("foobar", []rune("out foobar"), fileRef)

	ut := new(lang.UnitTests)
	ut.Add("random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs", plan, fileRef)

	if ut.Run(lang.ShellProcess.Tests, "foobar") {
		t.Error("TestRunTestNotATest passed when expected to fail")
	} else {
		b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
		if err != nil {
			panic(err)
		}
		t.Logf("Test report:\n%s", b)
	}
}

func TestRunTestNotAFunction(t *testing.T) {
	count.Tests(t, 1, "TestRunTestNotAFunction")

	lang.InitEnv()

	fileRef := &ref.File{
		Source: &ref.Source{
			Filename: "foobar.mx",
			Module:   "foo/bar",
			DateTime: time.Now(),
		},
	}

	plan := &lang.UnitTestPlan{
		Stdout: "foobar\n",
	}

	lang.MxFunctions.Define("foobar", []rune("out foobar"), fileRef)

	ut := new(lang.UnitTests)
	ut.Add("random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs", plan, fileRef)

	if ut.Run(lang.ShellProcess.Tests, "random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs") {
		t.Error("TestRunTestNotATest passed when expected to fail")
	} else {
		b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
		if err != nil {
			panic(err)
		}
		t.Logf("Test report:\n%s", b)
	}
}
