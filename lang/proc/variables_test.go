// +build ignore

package proc

import (
	"strconv"
	"testing"

	"github.com/lmorg/murex/lang/types"
)

// TestVariables tests the Vars structure
func TestVariables(t *testing.T) {
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
	orig := NewVariables(ShellProcess)

	err := orig.Set("number", origNum, types.Number)
	if err != nil {
		t.Error("Unable to set number in original. " + err.Error())
	}

	err = orig.Set("integer", origInt, types.Integer)
	if err != nil {
		t.Error("Unable to set integer in original. " + err.Error())
	}

	err = orig.Set("string", origStr, types.String)
	if err != nil {
		t.Error("Unable to set string in original. " + err.Error())
	}

	err = orig.Set("boolean", origBool, types.Boolean)
	if err != nil {
		t.Error("Unable to set boolean in original. " + err.Error())
	}

	p := &Process{Id: 1, FidTree: []int{0, 1}}
	p.Parent = p
	copy := NewVariables(p)

	err = copy.Set("number", copyNum, types.Number)
	if err != nil {
		t.Error("Unable to set number in copy. " + err.Error())
	}

	err = copy.Set("integer", copyInt, types.Integer)
	if err != nil {
		t.Error("Unable to set integer in copy. " + err.Error())
	}

	err = copy.Set("string", copyStr, types.String)
	if err != nil {
		t.Error("Unable to set string in copy. " + err.Error())
	}

	err = copy.Set("boolean", copyBool, types.Boolean)
	if err != nil {
		t.Error("Unable to set boolean in copy. " + err.Error())
	}

	// test values changed

	if copy.varTable.vars[0].Value != orig.varTable.vars[0].Value {
		t.Error("Copy and original shouldn't share same value for number.")
	}

	if copy.varTable.vars[1].Value != orig.varTable.vars[1].Value {
		t.Error("Copy and original shouldn't share same value for integer.")
	}

	if copy.varTable.vars[2].Value != orig.varTable.vars[2].Value {
		t.Error("Copy and original shouldn't share same value for string.")
	}

	if copy.varTable.vars[3].Value != orig.varTable.vars[3].Value {
		t.Error("Copy and original shouldn't share same value for boolean.")
	}

	// test copy values

	if copy.varTable.vars[0].Value.(float64) != copyNum {
		t.Error("Copy number not same as expected value.")
	}

	if copy.varTable.vars[1].Value.(int) != copyInt {
		t.Error("Copy integer not same as expected value.")
	}

	if copy.varTable.vars[2].Value.(string) != copyStr {
		t.Error("Copy string not same as expected value.")
	}

	if copy.varTable.vars[3].Value.(bool) != copyBool {
		t.Error("Copy boolean not same as expected value.")
	}

	// test GetValue on copy

	if copy.GetValue("number").(float64) != copyNum {
		t.Error("Copy var table not returning correct number using GetValue.")
	}

	if copy.GetValue("integer").(int) != copyInt {
		t.Error("Copy var table not returning correct integer using GetValue.")
	}

	if copy.GetValue("string").(string) != copyStr {
		t.Error("Copy var table not returning correct string using GetValue.")
	}

	if copy.GetValue("boolean").(bool) != copyBool {
		t.Error("Copy var table not returning correct boolean using GetValue.")
	}

	// test GetString on copy

	if copy.GetString("number") != types.FloatToString(copyNum) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if copy.GetString("integer") != strconv.Itoa(copyInt) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if copy.GetString("string") != copyStr {
		t.Error("Copy var table not returning correct string converted value using GetString.")
	}

	if types.IsTrue([]byte(copy.GetString("boolean")), 0) != copyBool {
		t.Error("Copy var table not returning correct boolean converted value using GetString.")
	}

	err = copy.Set("new", "string", types.String)
	if err != nil {
		t.Error("Unable to create new string. " + err.Error())
	}

	if orig.GetString("new") != "" {
		t.Error("New string exists on original when not expected.")
	}

	if orig.GetString("new") == "string" {
		t.Error("New string saved on copy was replicated on original - this shouldn't happen.")
	}

	if copy.GetString("new") != "string" {
		t.Error("New string on copy not retriving correctly: '" + copy.GetString("new") + "'")
	}
}
