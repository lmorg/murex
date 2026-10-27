//go:build !js && !windows && !plan9
// +build !js,!windows,!plan9

package terminal

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestTitleSanitization(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"punctuation", `title ; "quoted" C:\path`, `title ; "quoted" C:\path`},
		{"unicode", "世界 ü 🚀", "世界 ü 🚀"},
		{"utf8 continuation bytes", "Üĝ", "Üĝ"},
		{"ascii controls", "a\x07\x1b\\\n\t\x00\x7fb\rc", "a     bc"},
		{"raw C1", "a\x9c\x9d\x9bb", "a   b"},
		{"encoded C1", "a\u009c\u009d\u0085b", "a   b"},
		{"invalid utf8", "a\xff\xc0\xafb", "a   b"},
		{"color", "\x1b[31mred\x1b[0m", "red"},
		{"osc", "a\x1b]2;injected\x07b", "ab"},
		{"apc", "a\x1b_injected\x1b\\b", "ab"},
		{"tmux terminator", "a\x1b\\b", "ab"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			count.Tests(t, 3)
			for _, formatter := range []struct {
				format func([]byte) []byte
				prefix string
			}{
				{formatTitle, "\x1b]2;"},
				{formatIcon, "\x1b]1;"},
				{formatTmux, "\x1bk"},
			} {
				expected := formatter.prefix + testCase.want + "\x1b\\"
				if actual := formatter.format([]byte(testCase.input)); string(actual) != expected {
					t.Errorf("got %q, want %q", actual, expected)
				}
			}
		})
	}
}

func TestFormatTitle(t *testing.T) {
	type testT struct {
		Title  string
		Format []byte
	}

	tests := []testT{
		{
			Title:  "",
			Format: nil,
		},
		{
			Title:  "1",
			Format: []byte{27, ']', '2', ';', '1', 27, '\\'},
		},
		{
			Title:  "12",
			Format: []byte{27, ']', '2', ';', '1', '2', 27, '\\'},
		},
		{
			Title:  "123",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', 27, '\\'},
		},
		{
			Title:  "1234",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', 27, '\\'},
		},
		{
			Title:  "12345",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', 27, '\\'},
		},
		{
			Title:  "123456",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', '6', 27, '\\'},
		},
		{
			Title:  "1234567",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', '6', '7', 27, '\\'},
		},
		{
			Title:  "12345678",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', '6', '7', '8', 27, '\\'},
		},
		{
			Title:  "123456789",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', '6', '7', '8', '9', 27, '\\'},
		},
		{
			Title:  "1234567890",
			Format: []byte{27, ']', '2', ';', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 27, '\\'},
		},
		// unicode
		{
			Title:  "Hello, 世界",
			Format: append([]byte{27, ']', '2', ';'}, []byte("Hello, 世界\x1b\\")...),
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {

		actual := formatTitle([]byte(test.Title))

		if string(actual) != string(test.Format) {
			t.Errorf("Format error in %s test %d", t.Name(), i)
			t.Logf("  Title:    '%s'", test.Title)
			t.Logf("  Expected: '%s'", string(test.Format))
			t.Logf("  Actual:   '%s'", string(actual))
			t.Logf("  exp byte:  %v", test.Format)
			t.Logf("  act byte:  %v", actual)
		}
	}
}

func TestFormatIcon(t *testing.T) {
	type testT struct {
		Title  string
		Format []byte
	}

	tests := []testT{
		{
			Title:  "",
			Format: nil,
		},
		{
			Title:  "1",
			Format: []byte{27, ']', '1', ';', '1', 27, '\\'},
		},
		{
			Title:  "12",
			Format: []byte{27, ']', '1', ';', '1', '2', 27, '\\'},
		},
		{
			Title:  "123",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', 27, '\\'},
		},
		{
			Title:  "1234",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', 27, '\\'},
		},
		{
			Title:  "12345",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', 27, '\\'},
		},
		{
			Title:  "123456",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', '6', 27, '\\'},
		},
		{
			Title:  "1234567",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', '6', '7', 27, '\\'},
		},
		{
			Title:  "12345678",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', '6', '7', '8', 27, '\\'},
		},
		{
			Title:  "123456789",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', '6', '7', '8', '9', 27, '\\'},
		},
		{
			Title:  "1234567890",
			Format: []byte{27, ']', '1', ';', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 27, '\\'},
		},
		// unicode
		{
			Title:  "Hello, 世界",
			Format: append([]byte{27, ']', '1', ';'}, []byte("Hello, 世界\x1b\\")...),
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {

		actual := formatIcon([]byte(test.Title))

		if string(actual) != string(test.Format) {
			t.Errorf("Format error in %s test %d", t.Name(), i)
			t.Logf("  Title:    '%s'", test.Title)
			t.Logf("  Expected: '%s'", string(test.Format))
			t.Logf("  Actual:   '%s'", string(actual))
			t.Logf("  exp byte:  %v", test.Format)
			t.Logf("  act byte:  %v", actual)
		}
	}
}

func TestFormatTmux(t *testing.T) {
	type testT struct {
		Title  string
		Format []byte
	}

	tests := []testT{
		{
			Title:  "",
			Format: nil,
		},
		{
			Title:  "1",
			Format: []byte{27, 'k', '1', 27, '\\'},
		},
		{
			Title:  "12",
			Format: []byte{27, 'k', '1', '2', 27, '\\'},
		},
		{
			Title:  "123",
			Format: []byte{27, 'k', '1', '2', '3', 27, '\\'},
		},
		{
			Title:  "1234",
			Format: []byte{27, 'k', '1', '2', '3', '4', 27, '\\'},
		},
		{
			Title:  "12345",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', 27, '\\'},
		},
		{
			Title:  "123456",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', '6', 27, '\\'},
		},
		{
			Title:  "1234567",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', '6', '7', 27, '\\'},
		},
		{
			Title:  "12345678",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', '6', '7', '8', 27, '\\'},
		},
		{
			Title:  "123456789",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', '6', '7', '8', '9', 27, '\\'},
		},
		{
			Title:  "1234567890",
			Format: []byte{27, 'k', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 27, '\\'},
		},
		// unicode
		{
			Title:  "Hello, 世界",
			Format: []byte("\x1bkHello, 世界\x1b\\"),
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {

		actual := formatTmux([]byte(test.Title))

		if string(actual) != string(test.Format) {
			t.Errorf("Format error in %s test %d", t.Name(), i)
			t.Logf("  Title:    '%s'", test.Title)
			t.Logf("  Expected: '%s'", string(test.Format))
			t.Logf("  Actual:   '%s'", string(actual))
			t.Logf("  exp byte:  %v", test.Format)
			t.Logf("  act byte:  %v", actual)
		}
	}
}
