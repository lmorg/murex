//go:build go1.18 && !js && !windows && !plan9
// +build go1.18,!js,!windows,!plan9

package ansititle

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

var fuzzTests = []string{"", ".", `!"£$%^&*()`, "12345!", "foobar", "世", "世界"}

func FuzzFormatTitle(f *testing.F) {
	for _, tc := range fuzzTests {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, title string) {
		count.Tests(t, 1)
		formatTitle([]byte(title))
		// we are just testing we can't cause an unhandled panic
	})
}

func FuzzFormatTmux(f *testing.F) {
	for _, tc := range fuzzTests {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, title string) {
		count.Tests(t, 1)
		formatTmux([]byte(title))
		// we are just testing we can't cause an unhandled panic
	})
}

func FuzzSanatise(f *testing.F) {
	for _, tc := range fuzzTests {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, title string) {
		count.Tests(t, 1)
		sanatise([]byte(title))
		// we are just testing we can't cause an unhandled panic
	})
}
