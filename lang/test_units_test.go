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
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/test/count"
)

var uniq int32

type testUTPs struct {
	Function  string
	TestBlock string
	Passed    bool
	UTP       lang.UnitTestPlan
}

func testRunTest(t *testing.T, plans []testUTPs) {
	t.Helper()
	count.Tests(t, len(plans)*2)

	lang.InitEnv()
	lang.ShellProcess.Config.Set("test", "auto-report", false)

	var pubPriv string

	for i := range plans {
		for j := 1; j < 3; j++ { // test public functions the private functions

			fileRef := &ref.File{
				Source: &ref.Source{
					Filename: "foobar.mx",
					Module:   fmt.Sprintf("foobar/mod-%d-%d-%d", atomic.AddInt32(&uniq, 1), i, j),
					DateTime: time.Now(),
				},
			}

			if j == 1 {
				lang.MxFunctions.Define(plans[i].Function, []rune(plans[i].TestBlock), fileRef)
				pubPriv = "public"
			} else {
				lang.PrivateFunctions.Define(plans[i].Function, []rune(plans[i].TestBlock), fileRef)
				plans[i].Function = fileRef.Source.Module + "/" + plans[i].Function
				pubPriv = "private"
			}

			ut := new(lang.UnitTests)
			ut.Add(plans[i].Function, &plans[i].UTP, fileRef)

			if ut.Run(lang.ShellProcess, plans[i].Function) != plans[i].Passed {
				if plans[i].Passed {
					t.Errorf("Unit test %s %d failed", pubPriv, i)
					b, err := json.MarshalIndent(lang.ShellProcess.Tests.Results.Dump(), "", "    ")
					if err != nil {
						panic(err)
					}

					t.Logf("  Test report:\n%s", b)

				} else {
					t.Errorf("Unit test %s %d passed when expected to fail", pubPriv, i)
				}
			}
			lang.ShellProcess.Tests.Results = new(lang.TestResults)
		}
	}
}
