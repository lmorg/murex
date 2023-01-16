package lang

import (
	"strconv"
	"strings"
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

// TestVariablesDefault tests with the F_DEFAULTS fork flag set
func TestVariablesDefaults(t *testing.T) {
	testVariables(t, F_DEFAULTS, "F_DEFAULTS")
}

// testVariables is the main testing function for variables
func testVariables(t *testing.T, flags int, details string) {
	t.Log("Testing with flags: " + details)
	const (
		origNum  = 123.123
		origInt  = 45678
		origStr  = "abcABC"
		origBool = true
		copyNum  = 987.789
		copyInt  = 65432
		copyStr  = "xyzXYZ"
		copyBool = false
	)

	InitEnv()

	p := NewTestProcess()

	// Create a referenced variable table
	count.Tests(t, 4)
	fork := p.Fork(flags)
	copy := fork.Variables

	err := copy.Set(p, "number", copyNum, types.Number)
	if err != nil {
		t.Error("Unable to set number in copy. " + err.Error())
	}

	err = copy.Set(p, "integer", copyInt, types.Integer)
	if err != nil {
		t.Error("Unable to set integer in copy. " + err.Error())
	}

	err = copy.Set(p, "string", copyStr, types.String)
	if err != nil {
		t.Error("Unable to set string in copy. " + err.Error())
	}

	err = copy.Set(p, "boolean", copyBool, types.Boolean)
	if err != nil {
		t.Error("Unable to set boolean in copy. " + err.Error())
	}

	// test GetValue
	count.Tests(t, 4)

	v := panicErr(copy.GetValue("number"))
	if v.(float64) != copyNum {
		t.Error("Copy var table not returning correct number using GetValue.")
	}

	v = panicErr(copy.GetValue("integer"))
	if v.(int) != copyInt {
		t.Error("Copy var table not returning correct integer using GetValue.")
	}

	v = panicErr(copy.GetValue("string"))
	if v.(string) != copyStr {
		t.Error("Copy var table not returning correct string using GetValue.")
	}

	v = panicErr(copy.GetValue("boolean"))
	if v.(bool) != copyBool {
		t.Error("Copy var table not returning correct boolean using GetValue.")
	}

	// test GetString on copy
	count.Tests(t, 4)

	if v, err := copy.GetString("number"); err != nil || v != types.FloatToString(copyNum) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if v, err := copy.GetString("integer"); err != nil || v != strconv.Itoa(copyInt) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if v, err := copy.GetString("string"); err != nil || v != copyStr {
		t.Error("Copy var table not returning correct string converted value using GetString.")
	}

	s, err := copy.GetString("boolean")
	if types.IsTrue([]byte(s), 0) != copyBool || err != nil {
		t.Error("Copy var table not returning correct boolean converted value using GetString.")
	}
}

func panicErr(v interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return v
}

// TestReservedVariables tests the Vars structure
func TestReservedVariables(t *testing.T) {
	p := NewTestProcess()

	reserved := []string{
		"SELF",
		"ARGS",
		"PARAMS",
		"MUREX_EXE",
		"MUREX_ARGS",
		"HOSTNAME",
		"PWD",
		"0", "1", "8", "9",
		"00", "10", "50", "99", "100",
		"_",
	}

	count.Tests(t, len(reserved))

	for _, name := range reserved {
		err := GlobalVariables.Set(p, name, "foobar", types.String)
		if err == nil || !strings.Contains(err.Error(), "reserved") {
			t.Errorf("`%s` is not a reserved variable", name)
		}
	}
}
