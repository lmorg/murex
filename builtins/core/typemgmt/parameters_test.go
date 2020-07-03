package typemgmt

import (
	"encoding/json"
	"testing"
)

func TestSplitVarString(t *testing.T) {
	splitTest := []struct {
		Input    []string
		DataType string
		Name     string
		Value    string
		Error    bool
	}{
		{
			Input:    []string{"foo=bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input: []string{"-=bar"},
			Error: true,
		},
		{
			Input: []string{"-", "=bar"},
			Error: true,
		},
		{
			Input: []string{"=bar"},
			Error: true,
		},
		{
			Input: []string{"", "=bar"},
			Error: true,
		},
		{
			Input:    []string{"foo =bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo= bar"},
			DataType: "",
			Name:     "foo",
			Value:    " bar",
		},
		{
			Input:    []string{"foo = bar"},
			DataType: "",
			Name:     "foo",
			Value:    " bar",
		},
		{
			Input:    []string{"foo  =bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo=  bar"},
			DataType: "",
			Name:     "foo",
			Value:    "  bar",
		},
		{
			Input:    []string{"foo  =  bar"},
			DataType: "",
			Name:     "foo",
			Value:    "  bar",
		},
		/////
		{
			Input: []string{" foo=bar"},
			Error: true,
		},
		{
			Input: []string{" foo =bar"},
			Error: true,
		},
		{
			Input: []string{" foo= bar"},
			Error: true,
		},
		{
			Input: []string{" foo = bar"},
			Error: true,
		},
		{
			Input:    []string{"foo  =bar "},
			DataType: "",
			Name:     "foo",
			Value:    "bar ",
		},
		{
			Input:    []string{"foo=  bar "},
			DataType: "",
			Name:     "foo",
			Value:    "  bar ",
		},
		{
			Input:    []string{"foo  =  bar "},
			DataType: "",
			Name:     "foo",
			Value:    "  bar ",
		},

		/////
		{
			Input: []string{"dt foo=bar"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar=value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar =value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar= value"},
			Error: true,
		},
		{
			Input: []string{"dt foo bar = value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar=value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar =value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar= value"},
			Error: true,
		},
		{
			Input: []string{"d t foo bar = value"},
			Error: true,
		},
		{
			Input:    []string{"dt", "foo=bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input: []string{"dt", "foo", "bar"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar=value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar =value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar= value"},
			Error: true,
		},
		{
			Input: []string{"dt", "foo", "bar = value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar=value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar =value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar= value"},
			Error: true,
		},
		{
			Input: []string{"d", "t", "foo", "bar = value"},
			Error: true,
		},
		/////
		{
			Input:    []string{"foo", "=", "bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo=", "bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"foo", "=bar"},
			DataType: "",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo", "=", "bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo=", "bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "foo", "=bar"},
			DataType: "dt",
			Name:     "foo",
			Value:    "bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo ", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo  bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "bar "},
			DataType: "dt",
			Name:     "test",
			Value:    "foo bar ",
		},
		{
			Input:    []string{"dt", "test", "=", "foo=bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo=bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo=", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo= bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "=bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo =bar",
		},
		{
			Input:    []string{"dt", "test", "=", "foo", "=", "bar"},
			DataType: "dt",
			Name:     "test",
			Value:    "foo = bar",
		},
	}

	for i := range splitTest {
		name, value, dataType, err := splitVarString(splitTest[i].Input)

		switch {
		case (err != nil) != splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: error condition doesn't match expected")
		case name != splitTest[i].Name && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: name doesn't match expected")
		case value != splitTest[i].Value && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: value doesn't match expected")
		case dataType != splitTest[i].DataType && !splitTest[i].Error:
			t.Error("Variable block incorrectly parsed: data type doesn't match expected")
		default:
			continue
		}

		b, jErr := json.Marshal(splitTest[i].Input)
		if jErr != nil {
			panic(jErr)
		}

		t.Logf("  Test #:    %d", i)
		t.Logf("  Input:     %s", string(b))
		t.Logf("  Exp Name: '%s'", splitTest[i].Name)
		t.Logf("  Act Name: '%s'", name)
		t.Logf("  Exp Val:  '%s'", splitTest[i].Value)
		t.Logf("  Act Val:  '%s'", value)
		t.Logf("  Exp DT:   '%s'", splitTest[i].DataType)
		t.Logf("  Act DT:   '%s'", dataType)
		t.Logf("  Exp Err:   %v (boolean state)", splitTest[i].Error)
		t.Logf("  Act Err:   %s (error message)", err)
	}
}
