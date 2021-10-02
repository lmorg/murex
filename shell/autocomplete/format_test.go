package autocomplete

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

var (
	testData = []string{
		"foobar",
		"foobar ",
		"foobar:",
		"foobar: ",
		"foo bar",
		"",
		"1",
		"123",
		"(",
		")",
		"{",
		"}",
		"=",
		"/",
		"'",
		`"`,
		";",
		"|",
		"?",
		"->",
		"#",

		"\\foobar",
		"\\foobar ",
		"foobar\\:",
		"foobar:\\ ",
		"foo\\ bar",
		"\\1",
		"1\\23",
		"\\(",
		"\\)",
		"\\{",
		"\\}",
		"\\=",
		"\\/",
		"\\'",
		`\"`,
		"\\;",
		"\\|",
		"\\?",
		"\\->",
		"-\\>",
		"\\-\\>",
		"\\#",

		`"foobar"`,
		`"foobar "`,
		`"foobar:"`,
		`"foobar: "`,
		`"foo bar"`,
		`"foo bar "`,
		`"1"`,
		`"1 "`,
		`"123"`,
		`"123 "`,
		`"("`,
		`")"`,
		`"{"`,
		`"}"`,
		`"="`,
		`"/"`,
		`"'"`,
		`";"`,
		`"|"`,
		`"?"`,
		`"->"`,
		`"#"`,

		`'foobar'`,
		`'foobar '`,
		`'foobar:'`,
		`'foobar: '`,
		`'foo bar'`,
		`'foo bar '`,
		`'1'`,
		`'1 '`,
		`'123'`,
		`'123 '`,
		`'('`,
		`')'`,
		`'{'`,
		`'}'`,
		`'='`,
		`'/'`,
		`'"'`,
		`';'`,
		`'|'`,
		`'?'`,
		`'->'`,
		`'#'`,

		`(foobar)`,
		`(foobar )`,
		`(foobar:)`,
		`(foobar: )`,
		`(foo bar)`,
		`(foo bar )`,
		`(1)`,
		`(1 )`,
		`(123)`,
		`(123 )`,
		`(()`,
		`())`,
		`({)`,
		`(})`,
		`(=)`,
		`(/)`,
		`(')`,
		`(")`,
		`(;)`,
		`(|)`,
		`(?)`,
		`(->)`,
		`(#)`,
	}

	expected = []string{
		"foobar ",
		"foobar\\ ",
		"foobar: ",
		"foobar:\\ ",
		"foo\\ bar ",
		" ",
		"1 ",
		"123 ",
		"\\( ",
		"\\) ",
		"\\{ ",
		"\\} ",
		"=",
		"/",
		"\\' ",
		`\" `,
		"\\; ",
		"\\| ",
		"\\? ",
		"-\\> ",
		"\\# ",

		"\\\\foobar ",
		"\\\\foobar\\ ",
		"foobar\\\\: ",
		"foobar:\\\\\\ ",
		"foo\\\\\\ bar ",
		"\\\\1 ",
		"1\\\\23 ",
		"\\\\\\( ",
		"\\\\\\) ",
		"\\\\\\{ ",
		"\\\\\\} ",
		"\\\\=",
		"\\\\/",
		"\\\\\\' ",
		`\\\" `,
		"\\\\\\; ",
		"\\\\\\| ",
		"\\\\\\? ",
		"\\\\-\\> ",
		"-\\\\> ",
		"\\\\-\\\\> ",
		"\\\\\\# ",

		`\"foobar\" `,
		`\"foobar\ \" `,
		`\"foobar:\" `,
		`\"foobar:\ \" `,
		`\"foo\ bar\" `,
		`\"foo\ bar\ \" `,
		`\"1\" `,
		`\"1\ \" `,
		`\"123\" `,
		`\"123\ \" `,
		`\"\(\" `,
		`\"\)\" `,
		`\"\{\" `,
		`\"\}\" `,
		`\"=\" `,
		`\"/\" `,
		`\"\'\" `,
		`\"\;\" `,
		`\"\|\" `,
		`\"\?\" `,
		`\"-\>\" `,
		`\"\#\" `,

		`\'foobar\' `,
		`\'foobar\ \' `,
		`\'foobar:\' `,
		`\'foobar:\ \' `,
		`\'foo\ bar\' `,
		`\'foo\ bar\ \' `,
		`\'1\' `,
		`\'1\ \' `,
		`\'123\' `,
		`\'123\ \' `,
		`\'\(\' `,
		`\'\)\' `,
		`\'\{\' `,
		`\'\}\' `,
		`\'=\' `,
		`\'/\' `,
		`\'\"\' `,
		`\'\;\' `,
		`\'\|\' `,
		`\'\?\' `,
		`\'-\>\' `,
		`\'\#\' `,

		`\(foobar\) `,
		`\(foobar\ \) `,
		`\(foobar:\) `,
		`\(foobar:\ \) `,
		`\(foo\ bar\) `,
		`\(foo\ bar\ \) `,
		`\(1\) `,
		`\(1\ \) `,
		`\(123\) `,
		`\(123\ \) `,
		`\(\(\) `,
		`\(\)\) `,
		`\(\{\) `,
		`\(\}\) `,
		`\(=\) `,
		`\(/\) `,
		`\(\'\) `,
		`\(\"\) `,
		`\(\;\) `,
		`\(\|\) `,
		`\(\?\) `,
		`\(-\>\) `,
		`\(\#\) `,
	}
)

// Tests the order is preserved and that items can be added the map without
// breaking the function
func TestFormatSuggestionsOrder(t *testing.T) {
	/*formatSuggestions := func(act *AutoCompleteT) {
		//sortCompletions(act.Items)
		formatSuggestionsArray(act.ParsedTokens, act.Items)
		formatSuggestionsMap(act.ParsedTokens, &act.Definitions)
	}*/

	count.Tests(t, len(testData)*3)

	tests := make([]string, len(testData))
	copy(tests, testData)
	original := make([]string, len(tests))
	copy(original, tests)

	act1 := AutoCompleteT{
		Items:       tests,
		Definitions: make(map[string]string),
	}

	FormatSuggestions(&act1)

	for i := range tests {
		if tests[i] != expected[i] {
			t.Error("formatSuggestionsArray mismatch:")
			t.Logf("  Original: '%s'", original[i])
			t.Logf("  Expected: '%s'", expected[i])
			t.Logf("  Actual:   '%s'", tests[i])
			t.Logf("  Index:     %d", i)
		}

		act2 := AutoCompleteT{
			Definitions: map[string]string{original[i]: "some data"},
		}

		FormatSuggestions(&act2)

		if len(act2.Definitions) != 1 {
			t.Errorf("Invalid test length in formatSuggestionsMap:")
			t.Logf("  Expected:  %d", 1)
			t.Logf("  Actual:    %d", len(act2.Definitions))
			t.Logf("  Index:     %d", i)
		}
		for key := range act2.Definitions {
			if key != expected[i] {
				t.Error("formatSuggestionsMap mismatch:")
				t.Logf("  Original: '%s'", original[i])
				t.Logf("  Expected: '%s'", expected[i])
				t.Logf("  Actual:   '%s'", key)
				t.Logf("  Index:     %d", i)
			}
		}
	}
}

// Checks the completions are correct albeit not ordered
func TestFormatSuggestionsUnordered(t *testing.T) {
	count.Tests(t, len(testData)*2)

	original := make([]string, len(testData))
	copy(original, testData)

	for i := range testData {
		act := &AutoCompleteT{
			Items:       []string{testData[i]},
			Definitions: map[string]string{testData[i]: "foobar"},
		}

		FormatSuggestions(act)

		if act.Items[0] != expected[i] {
			t.Error("formatSuggestionsArray mismatch:")
			t.Logf("  Original: '%s'", original[i])
			t.Logf("  Expected: '%s'", expected[i])
			t.Logf("  Actual:   '%s'", testData[i])
			t.Logf("  Index:     %d", i)
		}

		if len(act.Definitions) != 1 {
			t.Errorf("formatSuggestionsMap len mismatch: %d", len(act.Definitions))
			return
		}

		for key := range act.Definitions {
			if key != expected[i] {
				t.Error("formatSuggestionsMap mismatch:")
				t.Logf("  Original: '%s'", original[i])
				t.Logf("  Expected: '%s'", expected[i])
				t.Logf("  Actual:   '%s'", key)
				t.Logf("  Index:     %d", i)
			}
		}
	}
}
