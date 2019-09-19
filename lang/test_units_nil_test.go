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
	"encoding/json"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

func TestRunTestNotATest(t *testing.T) {
	count.Tests(t, 1)

	lang.InitEnv()
	lang.ShellProcess.Config.Set("test", "auto-report", false)

	fileRef := &ref.File{
		Source: &ref.Source{
			Filename: "foobar.mx",
			Module:   "foo/bar",
			DateTime: time.Now(),
		},
	}

	plan := &lang.UnitTestPlan{
		StdoutMatch: "foobar\n",
	}

	lang.MxFunctions.Define("foobar", []rune("out foobar"), fileRef)

	ut := new(lang.UnitTests)
	ut.Add("random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs", plan, fileRef)

	if ut.Run(lang.ShellProcess, "foobar") {
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
	count.Tests(t, 1)

	lang.InitEnv()
	lang.ShellProcess.Config.Set("test", "auto-report", false)

	fileRef := &ref.File{
		Source: &ref.Source{
			Filename: "foobar.mx",
			Module:   "foo/bar",
			DateTime: time.Now(),
		},
	}

	plan := &lang.UnitTestPlan{
		StdoutMatch: "foobar\n",
	}

	lang.MxFunctions.Define("foobar", []rune("out foobar"), fileRef)

	ut := new(lang.UnitTests)
	ut.Add("random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs", plan, fileRef)

	if ut.Run(lang.ShellProcess, "random_string_that_shouldnt_exist_kjhadgkjsdhgfksdahfgsadhjsdfjksadfhs") {
		t.Error("TestRunTestNotATest passed when expected to fail")
	} else {
		b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
		if err != nil {
			panic(err)
		}
		t.Logf("Test report:\n%s", b)
	}
}
