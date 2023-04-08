package readline

import (
	"fmt"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
	"github.com/mattn/go-runewidth"
)

func TestLineWrap(t *testing.T) {
	type TestLineWrapT struct {
		Prompt    string
		Line      string
		TermWidth int
		Expected  []string
	}

	tests := []TestLineWrapT{
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 80,
			Expected:  []string{"1234567890 "},
		},
		{
			Prompt:    "foobar",
			Line:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			TermWidth: 86,
			Expected:  []string{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", "       "},
		},
		{
			Prompt:    "foobar",
			Line:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			TermWidth: 87,
			Expected:  []string{"12345678901234567890123456789012345678901234567890123456789012345678901234567890 "},
		},
		{
			Prompt:    "foobar",
			Line:      "123456789012345678901234567890",
			TermWidth: 20,
			Expected:  []string{"12345678901234", "      56789012345678", "      90 "},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 4,
			Expected:  []string{"1234", "5678", "90 "},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 5,
			Expected:  []string{"12345", "67890", " "},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 6,
			Expected:  []string{"123456", "7890 "},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 7,
			Expected:  []string{"1", "      2", "      3", "      4", "      5", "      6", "      7", "      8", "      9", "      0", "       "},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 8,
			Expected:  []string{"12", "      34", "      56", "      78", "      90", "       "},
		},
		{
			Prompt:    "foobar",
			Line:      "使用再生纸",
			TermWidth: 8,
			Expected:  []string{"使", "      用", "      再", "      生", "      纸", "       "},
		},
		{
			Prompt:    "foobar",
			Line:      "使用再生纸",
			TermWidth: 9,
			Expected:  []string{"使", "      用", "      再", "      生", "      纸 "},
		},
	}

	count.Tests(t, len(tests))
	for i, test := range tests {
		rl := NewInstance()
		rl.SetPrompt(test.Prompt)
		rl.line.Set([]rune(test.Line))

		wrap := lineWrap(rl, test.TermWidth)
		if len(wrap) != len(test.Expected) {
			t.Error("Slice lens do not match:")
			t.Logf("  Test:         %d (%s)", i, t.Name())
			t.Logf("  Prompt:      '%s'", test.Prompt)
			t.Logf("  Line:        '%s'", test.Line)
			t.Logf("  Width:        %d", test.TermWidth)
			t.Logf("  len(exp):     %d", len(test.Expected))
			t.Logf("  len(act):     %d", len(wrap))
			t.Logf("  Slice exp:   '%s'", fmt.Sprint(test.Expected))
			t.Logf("  Slice act:   '%s'", fmt.Sprint(wrap))
			t.Logf("  Slice json e: %s", json.LazyLogging(test.Expected))
			t.Logf("  Slice json a: %s", json.LazyLogging(wrap))
			t.Logf("  rl.promptLen: %d'", rl.promptLen)
			t.Logf("  rl.line:     '%s'", rl.line.String())
			continue
		}

		for j := range wrap {
			if wrap[j] != test.Expected[j] {
				t.Error("Slice element does not match:")
				t.Logf("  Test:      %d (%s)", i, t.Name())
				t.Logf("  Prompt:   '%s'", test.Prompt)
				t.Logf("  Line:     '%s'", test.Line)
				t.Logf("  Width:     %d", test.TermWidth)
				t.Logf("  Expected:  %s", test.Expected[j])
				t.Logf("  Actual:    %s", wrap[j])
				t.Logf("  len(exp):  %d", len(test.Expected))
				t.Logf("  len(act):  %d", len(wrap))
				t.Logf("  Slice exp:'%s'", fmt.Sprint(test.Expected))
				t.Logf("  Slice act:'%s'", fmt.Sprint(wrap))
				t.Logf("  Slice j e: %s", json.LazyLogging(test.Expected))
				t.Logf("  Slice j a: %s", json.LazyLogging(wrap))
			}
		}
	}
}

func TestLineWrapCell(t *testing.T) {
	type ExpectedT struct {
		X, Y int
	}

	type TestLineWrapPosT struct {
		Prompt    string
		Line      string
		TermWidth int
		Expected  ExpectedT
	}

	tests := []TestLineWrapPosT{
		{
			Prompt:    "12345",
			Line:      "",
			TermWidth: 10,
			Expected:  ExpectedT{5 + 0, 0},
		},
		/////
		{
			Prompt:    "12345",
			Line:      "123",
			TermWidth: 10,
			Expected:  ExpectedT{5 + 3, 0},
		},
		{
			Prompt:    "12345",
			Line:      "1234",
			TermWidth: 10,
			Expected:  ExpectedT{5 + 4, 0},
		},
		{
			Prompt:    "12345",
			Line:      "12345",
			TermWidth: 10,
			Expected:  ExpectedT{10, 0},
		},
		{
			Prompt:    "12345",
			Line:      "123456",
			TermWidth: 10,
			Expected:  ExpectedT{5 + 1, 1},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 80,
			Expected:  ExpectedT{6 + 10, 0},
		},
		{
			Prompt:    "foobar",
			Line:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			TermWidth: 85,
			Expected:  ExpectedT{6 + 1, 1},
		},
		{
			Prompt:    "foobar",
			Line:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			TermWidth: 86,
			Expected:  ExpectedT{86, 0},
		},
		{
			Prompt:    "foobar",
			Line:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890",
			TermWidth: 87,
			Expected:  ExpectedT{86, 0},
		},
		{
			Prompt:    "foobar",
			Line:      "123456789012345678901234567890",
			TermWidth: 20,
			//Expected:  []string{"12345678901234", "56789012345678", "90"},
			Expected: ExpectedT{6 + 2, 2},
		},
		{ // 10
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 4,
			//Expected:  []string{"1234", "5678", "90"},
			Expected: ExpectedT{0 + 2, 2},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 5,
			//Expected:  []string{"12345", "67890"},
			Expected: ExpectedT{5, 1},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 6,
			//Expected:  []string{"123456", "7890"},
			Expected: ExpectedT{0 + 4, 1},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 7,
			//Expected:  []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
			Expected: ExpectedT{7, 9},
		},
		{
			Prompt:    "foobar",
			Line:      "1234567890",
			TermWidth: 8,
			//Expected:  []string{"12", "34", "56", "78", "90"},
			Expected: ExpectedT{6 + 2, 4},
		},
		/////
		{
			Prompt:    "foobar",
			Line:      "使用再生纸",
			TermWidth: 8,
			//Expected:  []string{"12", "34", "56", "78", "90"},
			Expected: ExpectedT{6 + 2, 4},
		},
		{
			Prompt:    "foo",
			Line:      "使用 再生纸",
			TermWidth: 8,
			//Expected:  []string{"12", "34", "56", "78", "90"},
			Expected: ExpectedT{X: 5, Y: 2},
		},
		{
			Prompt:    "foo",
			Line:      "使用 再生纸 使用 再生",
			TermWidth: 8,
			//Expected:  []string{"12", "34", "56", "78", "90"},
			Expected: ExpectedT{X: 5, Y: 4},
		},
		{
			Prompt:    "使用",
			Line:      "使用再生纸使用再生",
			TermWidth: 8,
			//Expected:  []string{"12", "34", "56", "78", "90"},
			Expected: ExpectedT{X: 6, Y: 4},
		},
	}

	count.Tests(t, len(tests))
	for i, test := range tests {
		promptLen := runewidth.StringWidth(test.Prompt)
		x, y := lineWrapCell(promptLen, []rune(test.Line), test.TermWidth)

		if (test.Expected.X != x) || (test.Expected.Y != y) {
			t.Error("X or Y does not match:")
			t.Logf("  Test:      %d (%s)", i, t.Name())
			t.Logf("  Prompt:   '%s'", test.Prompt)
			t.Logf("  Prompt len %d", promptLen)
			t.Logf("  Line:     '%s'", test.Line)
			t.Logf("  Width:     %d", test.TermWidth)
			t.Logf("  Expected X:%d", test.Expected.X)
			t.Logf("  Actual   X:%d", x)
			t.Logf("  Expected Y:%d", test.Expected.Y)
			t.Logf("  Actual   Y:%d", y)
		}

	}
}
