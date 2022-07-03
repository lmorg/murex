package json_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testCase struct {
	Json     string
	Expected string
	Error    bool
}

func runTestCases(t *testing.T, tests []testCase) {
	count.Tests(t, len(tests))

	for i := range tests {
		testMx(t, tests[i].Json, tests[i].Expected, tests[i].Error, i)
	}
}

func testMx(t *testing.T, src string, sExp string, fail bool, testNum int) {
	var v interface{}
	err := json.UnmarshalMurex([]byte(src), &v)
	/*bAct, jsonErr := corejson.Marshal(v)
	if jsonErr != nil {
		t.Errorf("Unable to marshal Go struct from mxjson unmarshaller, this is possibly an error with the standard library: %s", jsonErr.Error())
	}

	sAct := string(bAct)
	bExp := []byte(sExp)*/

	if (err != nil) != fail {
		t.Errorf("Error response not as expected in test %d: ", testNum)
		t.Logf("  mxjson:  %s", src)
		t.Logf("  exp err: %v", fail)
		t.Logf("  act err: %v", err)
		/*t.Logf("  exp str: %s", sExp)
		t.Logf("  act str: %s", sAct)
		t.Log("  exp byt: ", bExp)
		t.Log("  act byt: ", bAct)*/
	}

	/*if sExp == "" && v == nil {
		return
	}

	if sExp != sAct {
		t.Errorf("Output doesn't match expected in test %d: ", testNum)
		t.Logf("  mxjson:  %s", src)
		t.Logf("  exp err: %v", fail)
		t.Logf("  act err: %v", err)
		t.Logf("  exp str: %s", sExp)
		t.Logf("  act str: %s", sAct)
		t.Log("  exp byt: ", bExp)
		t.Log("  act byt: ", bAct)
	}*/
}
