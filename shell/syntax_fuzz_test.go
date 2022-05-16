//go:build go1.18
// +build go1.18

package shell

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func FuzzSyntaxCompletion(f *testing.F) {
	pos := []int{-100, -5, -2, -1, 0, 1, 2, 5, 10, 20, 100}

	for _, i := range pos {
		f.Add("", "", i)
		f.Add("", ".", i)
		f.Add(".", "", i)

		f.Add(`!"£$%^&*(`, ``, i)
		f.Add("12345!", "", i)
		f.Add("foobar", "", i)
		f.Add("世", "", i)
		f.Add("世界", "", i)

		f.Add(`!"£$%^&*(`, `{`, i)
		f.Add("12345!", "[", i)
		f.Add("foobar", "(", i)
		f.Add("世", "(", i)
		f.Add("世界", "{", i)

		f.Add(`!"£$%^&*(`, `!"£$%^&*()`, i)
		f.Add("12345!", "12345!", i)
		f.Add("foobar", "foobar", i)
		f.Add("世", "界", i)
		f.Add("世界", "世界", i)
	}

	f.Fuzz(func(t *testing.T, line, change string, pos int) {
		count.Tests(t, 1)
		syntaxCompletion([]rune(line), change, pos)
		// we are just testing we can't cause an unhandled panic
	})
}
