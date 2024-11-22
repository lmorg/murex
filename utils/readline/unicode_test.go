package readline

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestSetRunePos(t *testing.T) {
	tests := []struct {
		Value        string
		Start        int
		RunePos      int
		ExpectedRPos int
		ExpectedCPos int
	}{
		{
			Value:        "hello world!",
			Start:        5,
			RunePos:      7,
			ExpectedRPos: 7,
			ExpectedCPos: 7,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      13,
			ExpectedRPos: 13,
			ExpectedCPos: 25,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      14,
			ExpectedRPos: 14,
			ExpectedCPos: 27,
		},
		{
			Value:        "foo举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      6,
			ExpectedRPos: 6,
			ExpectedCPos: 8,
		},
		{
			Value:        "foo举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      7,
			ExpectedRPos: 7,
			ExpectedCPos: 10,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      3,
			ExpectedRPos: 3,
			ExpectedCPos: 5,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			RunePos:      4,
			ExpectedRPos: 4,
			ExpectedCPos: 7,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		u := new(UnicodeT)
		u.Set(new(Instance), []rune(test.Value))
		u.SetRunePos(test.Start)
		u.SetRunePos(test.RunePos)
		rPos := u.RunePos()
		cPos := u.CellPos()

		if rPos != test.ExpectedRPos || cPos != test.ExpectedCPos {
			t.Errorf("Unexpected position in test %d", i)
			t.Logf("  Value:      '%s'", test.Value)
			t.Logf("  Start:      %d", test.Start)
			t.Logf("  SetRunePos: %d", test.RunePos)
			t.Logf("  exp rPos:   %d", test.ExpectedRPos)
			t.Logf("  act rPos:   %d", rPos)
			t.Logf("  exp cPos:   %d", test.ExpectedCPos)
			t.Logf("  act cPos:   %d", cPos)
		}
	}
}

func TestSetCellPos(t *testing.T) {
	tests := []struct {
		Value        string
		Start        int
		CellPos      int
		ExpectedRPos int
		ExpectedCPos int
	}{
		{
			Value:        "hello world!",
			Start:        5,
			CellPos:      7,
			ExpectedRPos: 7,
			ExpectedCPos: 7,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      25,
			ExpectedRPos: 13,
			ExpectedCPos: 25,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      26,
			ExpectedRPos: 13,
			ExpectedCPos: 25,
		},
		{
			Value:        "foo举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      8,
			ExpectedRPos: 6,
			ExpectedCPos: 8,
		},
		{
			Value:        "foo举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      9,
			ExpectedRPos: 6,
			ExpectedCPos: 8,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      5,
			ExpectedRPos: 3,
			ExpectedCPos: 5,
		},
		{
			Value:        "举手之劳就可以使办公室更加环保，比如，使用再生纸。",
			Start:        2,
			CellPos:      6,
			ExpectedRPos: 3,
			ExpectedCPos: 5,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		u := new(UnicodeT)
		u.Set(new(Instance), []rune(test.Value))
		u.SetRunePos(test.Start)
		u.SetCellPos(test.CellPos)
		rPos := u.RunePos()
		cPos := u.CellPos()

		if rPos != test.ExpectedRPos || cPos != test.ExpectedCPos {
			t.Errorf("Unexpected position in test %d", i)
			t.Logf("  Value:      '%s'", test.Value)
			t.Logf("  Start:      %d", test.Start)
			t.Logf("  SetCellPos: %d", test.CellPos)
			t.Logf("  exp rPos:   %d", test.ExpectedRPos)
			t.Logf("  act rPos:   %d", rPos)
			t.Logf("  exp cPos:   %d", test.ExpectedCPos)
			t.Logf("  act cPos:   %d", cPos)
		}
	}
}
