//go:build go1.18
// +build go1.18

package inject_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/inject"
)

func FuzzInjectString(f *testing.F) {
	pos := []int{-100, -5, -2, -1, 0, 1, 2, 5, 10, 20, 100}

	for _, i := range pos {
		f.Add("", "", i)
		f.Add("", ".", i)
		f.Add(".", "", i)
		f.Add(`!"£$%^&*(`, `!"£$%^&*()`, i)
		f.Add("12345!", "12345!", i)
		f.Add("foobar", "foobar", i)
		f.Add("世", "界", i)
		f.Add("世界", "世界", i)
	}

	f.Fuzz(func(t *testing.T, old, insert string, pos int) {
		count.Tests(t, 1)
		inject.String(old, insert, pos)
		// we are just testing we can't cause an unhandled panic
	})
}

func FuzzInjectRune(f *testing.F) {
	pos := []int{-100, -5, -2, -1, 0, 1, 2, 5, 10, 20, 100}

	for _, i := range pos {
		f.Add("", "", i)
		f.Add("", ".", i)
		f.Add(".", "", i)
		f.Add(`!"£$%^&*(`, `!"£$%^&*()`, i)
		f.Add("12345!", "12345!", i)
		f.Add("foobar", "foobar", i)
		f.Add("世", "界", i)
		f.Add("世界", "世界", i)
	}

	f.Fuzz(func(t *testing.T, old, insert string, pos int) {
		count.Tests(t, 1)
		inject.Rune([]rune(old), []rune(insert), pos)
		// we are just testing we can't cause an unhandled panic
	})
}
