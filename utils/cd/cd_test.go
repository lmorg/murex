package cd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/cd"
)

func TestCd(t *testing.T) {
	// Just test we can actually change directories in Go first.
	// This pre-test also has the benefit of fixing any issues that symlinks
	// might cause the later test - which might result in the test working but
	// producing a fail condition.
	count.Tests(t, 1)
	new := ".."
	before, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = os.Chdir(new)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	new, err = os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	err = os.Chdir(before)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	before, err = os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// initialise test environment
	count.Tests(t, 2)

	lang.InitEnv()
	p := lang.NewTestProcess()

	// add some history to the path history global
	err = lang.GlobalVariables.Set(p, cd.GlobalVarName, []string{"/1/1/1", "/2/2/2"}, types.Json)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = cd.Chdir(p, new)
	if err != nil {
		t.Errorf(err.Error())
	}

	pwdHist, _ := lang.GlobalVariables.GetValue(cd.GlobalVarName)

	switch pwdHist.(type) {
	case []string:
		l := len(pwdHist.([]string))
		if l < 2 {
			t.Errorf("1st pass: $%s len() too short: %d", cd.GlobalVarName, l)
		}
		if pwdHist.([]string)[l-1] != new {
			t.Errorf("1st pass: $%s not appending new directory", cd.GlobalVarName)
			t.Logf("  Expected: '%s'", new)
			t.Logf("  Actual:   '%s'", pwdHist.([]string)[l-1])
			t.Logf("  Array:     %s", pwdHist.([]string))
			t.Logf("  Location:  %d of %d", l-1, len(pwdHist.([]string))-1)
		}

	default:
		t.Errorf("1st pass: $%s.(type) != []string", cd.GlobalVarName)
	}

	// break the path history
	err = lang.GlobalVariables.Set(p, cd.GlobalVarName, "", types.String)
	if err != nil {
		t.Errorf(err.Error())
	}

	// set the path back again
	err = cd.Chdir(p, before)
	if err != nil {
		t.Errorf(err.Error())
	}

	pwdHist, _ = lang.GlobalVariables.GetValue(cd.GlobalVarName)
	match2 := "murex/utils/cd"

	switch pwdHist.(type) {
	case []string:
		l := len(pwdHist.([]string))
		if l != 1 {
			t.Errorf("2nd pass: $%s len() incorrect. Expecting 1, got %d", cd.GlobalVarName, l)
		}
		s := pwdHist.([]string)[l-1]
		if s != before && !strings.HasSuffix(s, match2) {
			t.Errorf("2nd pass: $%s not appending new directory", cd.GlobalVarName)
			t.Logf("  Expected1: '%s'", before)
			t.Logf("  Actual1:   '%s'", s)
			t.Logf("  Expected2: '%s'", match2)
			t.Logf("  Actual2:   '%s'", s[len(s)-len(match2):])
			t.Logf("  Array:      %s", pwdHist.([]string))
			t.Logf("  Location:   %d of %d", l-1, len(s)-1)
		}

	default:
		t.Errorf("2nd pass: $%s.(type) != []string", cd.GlobalVarName)
	}
}
