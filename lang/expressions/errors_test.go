package expressions

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestTrimCodeInErrMsg(t *testing.T) {
	tests := []struct {
		Code  string
		Pos   int
		Width int
	}{
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   62,
			Width: 100,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   1,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   5,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   10,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   14,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   15,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   16,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   20,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   25,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   30,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   35,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   40,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   45,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   50,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   55,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   60,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   61,
			Width: 10,
		},
		{
			Code:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
			Pos:   62,
			Width: 10,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		code := []rune(test.Code)
		pos := test.Pos - 1
		rr, ri := _cropCodeInErrMsg(code, test.Pos, test.Width)
		expected := code[pos]
		actual := rr[ri-1]
		if expected != actual {
			t.Errorf("error message incorrect in test %d", i)
			t.Logf("  Code:     '%s'", test.Code)
			t.Logf("  Pos:       %d, '%s' (%d)", test.Pos, string(test.Code[pos]), test.Code[pos])
			t.Logf("  Return:   '%s' (%d)", string(rr), ri)
			t.Logf("  expected: '%s' (%d)", string([]rune{expected}), expected)
			t.Logf("  actual:   '%s' (%d)", string([]rune{actual}), actual)
		}
	}
}
