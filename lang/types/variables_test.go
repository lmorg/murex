package types

import (
	"testing"
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

	vars := NewVariableGroup()

	err := vars.Set("number", origNum, Number)
	if err != nil {
		t.Error("Unable to set number in original.")
	}

	err = vars.Set("integer", origInt, Integer)
	if err != nil {
		t.Error("Unable to set integer in original.")
	}

	err = vars.Set("string", origStr, String)
	if err != nil {
		t.Error("Unable to set string in original.")
	}

	err = vars.Set("boolean", origBool, Boolean)
	if err != nil {
		t.Error("Unable to set boolean in original.")
	}

	copy := vars.Copy()

	err = copy.Set("number", copyNum, Number)
	if err != nil {
		t.Error("Unable to set number in copy.")
	}

	err = copy.Set("integer", copyInt, Integer)
	if err != nil {
		t.Error("Unable to set integer in copy.")
	}

	err = copy.Set("string", copyStr, String)
	if err != nil {
		t.Error("Unable to set string in copy.")
	}

	err = copy.Set("boolean", copyBool, Boolean)
	if err != nil {
		t.Error("Unable to set boolean in copy.")
	}

	// test values changed

	if copy.values["number"] == vars.values["number"] {
		t.Error("Copy and original shouldn't share same value for number.")
	}

	if copy.values["integer"] == vars.values["integer"] {
		t.Error("Copy and original shouldn't share same value for integer.")
	}

	if copy.values["string"] == vars.values["string"] {
		t.Error("Copy and original shouldn't share same value for string.")
	}

	if copy.values["boolean"] == vars.values["boolean"] {
		t.Error("Copy and original shouldn't share same value for boolean.")
	}

	// test original values

	if vars.values["number"].(float64) != origNum {
		t.Error("Original number not same as expected value.")
	}

	if vars.values["integer"].(int) != origInt {
		t.Error("Original integer not same as expected value.")
	}

	if vars.values["string"].(string) != origStr {
		t.Error("Original string not same as expected value.")
	}

	if vars.values["boolean"].(bool) != origBool {
		t.Error("Original boolean not same as expected value.")
	}

	// test copy values

	if copy.values["number"].(float64) != copyNum {
		t.Error("Copy number not same as expected value.")
	}

	if copy.values["integer"].(int) != copyInt {
		t.Error("Copy integer not same as expected value.")
	}

	if copy.values["string"].(string) != copyStr {
		t.Error("Copy string not same as expected value.")
	}

	if copy.values["boolean"].(bool) != copyBool {
		t.Error("Copy boolean not same as expected value.")
	}

	// test GetValue on original

	if vars.GetValue("number").(float64) != origNum {
		t.Error("Original var table not returning correct number using GetValue.")
	}

	if vars.GetValue("integer").(int) != origInt {
		t.Error("Original var table not returning correct integer using GetValue.")
	}

	if vars.GetValue("string").(string) != origStr {
		t.Error("Original var table not returning correct string using GetValue.")
	}

	if vars.GetValue("boolean").(bool) != origBool {
		t.Error("Original var table not returning correct boolean using GetValue.")
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

	/*// test GetString on original

	if vars.GetString("number") != FloatToString(origNum) {
		t.Error("Original var table not returning correct numeric converted value using GetString.")
	}

	if vars.GetString("integer") != strconv.Itoa(origInt) {
		t.Error("Original var table not returning correct numeric converted value using GetString.")
	}

	if vars.GetString("string") != origStr {
		t.Error("Original var table not returning correct string converted value using GetString.")
	}

	if IsTrue([]byte(vars.GetString("boolean")), 0) != origBool {
		t.Error("Original var table not returning correct boolean converted value using GetString.")
	}

	// test GetString on copy

	if copy.GetString("number") != FloatToString(copyNum) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if copy.GetString("integer") != strconv.Itoa(copyInt) {
		t.Error("Copy var table not returning correct numeric converted value using GetString.")
	}

	if copy.GetString("string") != copyStr {
		t.Error("Copy var table not returning correct string converted value using GetString.")
	}

	if IsTrue([]byte(copy.GetString("boolean")), 0) != copyBool {
		t.Error("Copy var table not returning correct boolean converted value using GetString.")
	}*/
}
