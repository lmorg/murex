package escape

import (
	"testing"
)

// TestCommandLine checks the escape.CommandLine function escapes all values as expected
func TestCommandLine(t *testing.T) {
	s := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQWERTVWXYZ",
		"1234567890",
		"!\"£$%^&*()_+[]{};'#:@~\\|,./<>?",
		" \t\r\n",
	}

	CommandLine(s)

	if s[0] != "abcdefghijklmnopqrstuvwxyz" {
		t.Errorf("Lower case alphabet has been modified to: %s", s[0])
	}

	if s[1] != "ABCDEFGHIJKLMNOPQWERTVWXYZ" {
		t.Errorf("Upper case alphabet has been modified to: %s", s[1])
	}

	if s[2] != "1234567890" {
		t.Errorf("Numeric characters have been modified to: %s", s[2])
	}

	if s[3] != "!\\\"£\\$%^&*\\(\\)_+[]{};\\'\\#:\\@~\\\\\\|,./\\<\\>\\?" {
		t.Errorf("Extended characters have not been modified correctly: %s", s[3])
	}

	if s[4] != `\ \t\r\n` {
		t.Errorf("White space has not been modified correctly")
	}
}
