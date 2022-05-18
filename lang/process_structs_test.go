package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestArgs(t *testing.T) {
	type testT struct {
		NameIn    string
		ParamsIn  []string
		NameOut   string
		ParamsOut []string
	}

	tests := []testT{
		{
			NameIn:  "bob",
			NameOut: "bob",
		},
		{
			NameIn:    "bob",
			ParamsIn:  []string{"bob"},
			NameOut:   "bob",
			ParamsOut: []string{"bob"},
		},
		{
			NameIn:    "bob",
			ParamsIn:  []string{"exec"},
			NameOut:   "bob",
			ParamsOut: []string{"exec"},
		},

		{
			NameIn:  "exec",
			NameOut: "exec",
		},
		{
			NameIn:   "exec",
			ParamsIn: []string{"exec"},
			NameOut:  "exec",
		},
		{
			NameIn:   "exec",
			ParamsIn: []string{"bob"},
			NameOut:  "bob",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		nameAct, paramsAct := args(test.NameIn, test.ParamsIn)

		if len(test.ParamsIn) == 0 {
			test.ParamsIn = []string{}
		}

		if len(test.ParamsOut) == 0 {
			test.ParamsOut = []string{}
		}

		if nameAct != test.NameOut ||
			json.LazyLogging(paramsAct) != json.LazyLogging(test.ParamsOut) {
			t.Errorf("Test %d failed:", i)
			t.Logf("  Name:       `%s`", test.NameIn)
			t.Logf("  Params:      %s", json.LazyLogging(test.ParamsIn))
			t.Logf("  exp name:   `%s`", test.NameOut)
			t.Logf("  exp params:  %s", json.LazyLogging(test.ParamsOut))
			t.Logf("  act name:   `%s`", nameAct)
			t.Logf("  act params:  %s", json.LazyLogging(paramsAct))
		}
	}
}
