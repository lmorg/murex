package random

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func TestRandInteger(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.Integer},
			`^[0-9]+$`,
		)
	}
}

func TestRandIntegerMax(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.Integer, "9"},
			`^[0-9]$`,
		)
	}
}

func TestRandNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.Number},
			`^[0-9]+$`,
		)
	}
}

func TestRandNumberMax(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.Number, "9"},
			`^[0-9]$`,
		)
	}
}

func TestRandFloat(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.Float},
			`^[0-9]\.[0-9]+$`,
		)
	}
}

func TestRandString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.String},
			`^[\x20-\x7E]{20}$`,
		)
	}
}

func TestRandStringMax(t *testing.T) {
	for i := 0; i < 1000; i++ {
		test.RunMethodRegexTest(t, cmdRand, funcName,
			"", types.Null,
			[]string{types.String, "1"},
			`^[\x20-\x7E]$`,
		)
	}
}
