package autocomplete_test

import (
	"testing"

	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/test/count"
)

func TestFormatSuggestions(t *testing.T) {
	tests := []string{
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
	}

	expected := []string{
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
	}

	count.Tests(t, len(tests)*3)

	original := make([]string, len(tests))
	copy(original, tests)

	autocomplete.FormatSuggestions(&autocomplete.AutoCompleteT{
		Items:       tests,
		Definitions: make(map[string]string),
	})

	for i := range tests {
		if tests[i] != expected[i] {
			t.Error("formatSuggestionsArray mismatch:")
			t.Logf("  Original: '%s'", original[i])
			t.Logf("  Expected: '%s'", expected[i])
			t.Logf("  Actual:   '%s'", tests[i])
			t.Logf("  Index:     %d", i)
		}

		act := autocomplete.AutoCompleteT{
			Definitions: map[string]string{original[i]: "some data"},
		}

		autocomplete.FormatSuggestions(&act)

		if len(act.Definitions) != 1 {
			t.Errorf("Invalid test length in formatSuggestionsMap:")
			t.Logf("  Expected:  %d", 1)
			t.Logf("  Actual:    %d", len(act.Definitions))
			t.Logf("  Index:     %d", i)
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
