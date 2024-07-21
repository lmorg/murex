//go:build !js && !windows
// +build !js,!windows

package ansititle

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

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
		// TODO: test unicode
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
		// TODO: test unicode
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
		// TODO: test unicode
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
