package inject_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/inject"
)

type TestType struct {
	Input    string
	Inject   string
	Pos      int
	Expected string
	Error    bool
}

var tests = []TestType{
	{
		Input:    "",
		Inject:   "_",
		Pos:      0,
		Expected: "_",
		Error:    false,
	},
	{
		Input:    "",
		Inject:   "_",
		Pos:      1,
		Expected: "",
		Error:    true,
	},
	{
		Input:    "1",
		Inject:   "_",
		Pos:      0,
		Expected: "_1",
		Error:    false,
	},
	{
		Input:    "1",
		Inject:   "_",
		Pos:      1,
		Expected: "1_",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      -1,
		Expected: "",
		Error:    true,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      0,
		Expected: "_foobar",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      1,
		Expected: "f_oobar",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      2,
		Expected: "fo_obar",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      3,
		Expected: "foo_bar",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      4,
		Expected: "foob_ar",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      5,
		Expected: "fooba_r",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      6,
		Expected: "foobar_",
		Error:    false,
	},
	{
		Input:    "foobar",
		Inject:   "_",
		Pos:      7,
		Expected: "",
		Error:    true,
	},
}

func TestString(t *testing.T) {
	count.Tests(t, len(tests))

	for i, test := range tests {
		actual, err := inject.String(test.Input, test.Inject, test.Pos)

		if (err == nil) == test.Error {
			t.Errorf("%s failed:", t.Name())
			t.Logf("  Test #:    %d", i)
			t.Logf("  Input:    '%s'", test.Input)
			t.Logf("  Inject:   '%s'", test.Inject)
			t.Logf("  Pos:       %d", test.Pos)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Error:    '%T'", test.Error)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  Err Msg:  '%s'", err.Error())
			continue
		}

		if actual != test.Expected {
			t.Errorf("%s failed:", t.Name())
			t.Logf("  Test #:    %d", i)
			t.Logf("  Input:    '%s'", test.Input)
			t.Logf("  Inject:   '%s'", test.Inject)
			t.Logf("  Pos:       %d", test.Pos)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Error:    '%T'", test.Error)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  Err Msg:  '%s'", err.Error())
		}
	}
}

func TestRune(t *testing.T) {
	count.Tests(t, len(tests))

	for i, test := range tests {
		r, err := inject.Rune([]rune(test.Input), []rune(test.Inject), test.Pos)
		actual := string(r)

		if (err == nil) == test.Error {
			t.Errorf("%s failed:", t.Name())
			t.Logf("  Test #:    %d", i)
			t.Logf("  Input:    '%s'", test.Input)
			t.Logf("  Inject:   '%s'", test.Inject)
			t.Logf("  Pos:       %d", test.Pos)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Error:    '%T'", test.Error)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  Err Msg:  '%s'", err.Error())
			continue
		}

		if actual != test.Expected {
			t.Errorf("%s failed:", t.Name())
			t.Logf("  Test #:    %d", i)
			t.Logf("  Input:    '%s'", test.Input)
			t.Logf("  Inject:   '%s'", test.Inject)
			t.Logf("  Pos:       %d", test.Pos)
			t.Logf("  Expected: '%s'", test.Expected)
			t.Logf("  Error:    '%T'", test.Error)
			t.Logf("  Actual:   '%s'", actual)
			t.Logf("  Err Msg:  '%s'", err.Error())
		}
	}
}
