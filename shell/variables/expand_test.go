package variables

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/home"
)

var (
	testString = "|.|$foo|.|$bar|.|~|.|"
	expString  = "|.|oof|.|rab|.|" + home.MyDir + "|.|"
)

// TestExpand tests the ExpandString function
func TestExpand(t *testing.T) {
	count.Tests(t, 1, "TestExpand")

	lang.InitEnv()

	err := lang.ShellProcess.Variables.Set("foo", "oof", types.String)
	if err != nil {
		t.Error(err)
	}

	err = lang.ShellProcess.Variables.Set("bar", "rab", types.String)
	if err != nil {
		t.Error(err)
	}

	r := Expand([]rune(testString))
	if string(r) != expString {
		t.Error("String didn't expand as expected")
		t.Log("  testString:", testString)
		t.Log("  expString: ", expString)
		t.Log("  string(r): ", string(r))
	}
}

// TestExpandString tests the ExpandString function
func TestExpandString(t *testing.T) {
	count.Tests(t, 1, "TestExpandString")

	lang.InitEnv()

	err := lang.ShellProcess.Variables.Set("foo", "oof", types.String)
	if err != nil {
		t.Error(err)
	}

	err = lang.ShellProcess.Variables.Set("bar", "rab", types.String)
	if err != nil {
		t.Error(err)
	}

	s := ExpandString(testString)
	if s != expString {
		t.Error("String didn't expand as expected")
		t.Log("  testString:", testString)
		t.Log("  expString :", expString)
		t.Log("  s:         ", s)
	}
}

// TestCompare checks the Expand and ExpandString functions returns the same data (albeit in different data types)
func TestCompare(t *testing.T) {
	count.Tests(t, 1, "TestCompare")

	r := Expand([]rune(testString))
	s := ExpandString(testString)

	if string(r) != s {
		t.Error("Expand and ExpandString are not the same after data type conversion")
	}
}
