package utils

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestCrLfTrim tests the CrLfTrim function
func TestCrLfTrim(t *testing.T) {
	count.Tests(t, 5, "TestCrLfTrim")

	b := []byte("test")

	b = CrLfTrim(b)
	if string(b) != "test" {
		t.Error("CrLfTrim test 1 didn't return the correct []byte:", b)
	}

	b = append(b, '\r', '\n')
	b = CrLfTrim(b)
	if string(b) != "test" {
		t.Error("CrLfTrim test 2 didn't return the correct []byte:", b)
	}

	b = append(b, '\n')
	b = CrLfTrim(b)
	if string(b) != "test" {
		t.Error("CrLfTrim test 3 didn't return the correct []byte:", b)
	}

	b = append(b, '\r', '\n', '\r', '\n')
	b = CrLfTrim(b)
	if string(b) != "test\r\n" {
		t.Error("CrLfTrim test 4 didn't return the correct []byte:", b)
	}

	b = []byte("test\n\n")
	b = CrLfTrim(b)
	if string(b) != "test\n" {
		t.Error("CrLfTrim test 5 didn't return the correct []byte:", b)
	}
}

// TestCrLfTrimRune tests the CrLfTrimRune function
func TestCrLfTrimRune(t *testing.T) {
	count.Tests(t, 5, "TestCrLfTrimRune")

	r := []rune("test")

	r = CrLfTrimRune(r)
	if string(r) != "test" {
		t.Error("CrLfTrimRune test 1 didn't return the correct []rune:", r)
	}

	r = append(r, '\r', '\n')
	r = CrLfTrimRune(r)
	if string(r) != "test" {
		t.Error("CrLfTrimRune test 2 didn't return the correct []rune:", r)
	}

	r = append(r, '\n')
	r = CrLfTrimRune(r)
	if string(r) != "test" {
		t.Error("CrLfTrimRune test 3 didn't return the correct []rune:", r)
	}

	r = append(r, '\r', '\n', '\r', '\n')
	r = CrLfTrimRune(r)
	if string(r) != "test\r\n" {
		t.Error("CrLfTrimRune test 4 didn't return the correct []rune:", r)
	}

	r = []rune("test\n\n")
	r = CrLfTrimRune(r)
	if string(r) != "test\n" {
		t.Error("CrLfTrimRune test 5 didn't return the correct []rune:", r)
	}
}

// TestCrLfTrimString tests the CrLfTrimString function
func TestCrLfTrimString(t *testing.T) {
	count.Tests(t, 5, "TestCrLfTrimString")

	s := CrLfTrimString("test")

	if s != "test" {
		t.Error("CrLfTrimString test 1 didn't return the correct string:", []byte(s))
	}

	s = CrLfTrimString("test\r\n")
	if s != "test" {
		t.Error("CrLfTrimString test 2 didn't return the correct string:", []byte(s))
	}

	s = CrLfTrimString("test\n")
	if s != "test" {
		t.Error("CrLfTrimString test 3 didn't return the correct string:", []byte(s))
	}

	s = CrLfTrimString("test\r\n\r\n")
	if s != "test\r\n" {
		t.Error("CrLfTrimString test 4 didn't return the correct string:", []byte(s))
	}

	s = CrLfTrimString("test\n\n")
	if s != "test\n" {
		t.Error("CrLfTrimString test 5 didn't return the correct string:", []byte(s))
	}
}
