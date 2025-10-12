package shell

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func FuzzHint(f *testing.F) {
	pos := []int{-100, -5, -2, -1, 0, 1, 2, 5, 10, 20, 100}

	for _, i := range pos {
		f.Add("", i)
		f.Add(".", i)
		f.Add(`!"£$%^&*(`, i)
		f.Add("12345!", i)
		f.Add("foobar", i)
		f.Add("世", i)
		f.Add("世界", i)
	}

	lang.InitEnv()

	f.Fuzz(func(t *testing.T, line string, pos int) {
		count.Tests(t, 1)
		hintText([]rune(line), pos)
		// we are just testing we can't cause an unhandled panic
	})
}
