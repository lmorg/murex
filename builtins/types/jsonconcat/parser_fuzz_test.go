//go:build go1.18
// +build go1.18

package jsonconcat

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

var fuzzTests = []string{"{}{}", "[][]", `!"£$%^&*()`, "12345!", "foobar", "世", "世界"}

func FuzzParser(f *testing.F) {
	for _, tc := range fuzzTests {
		f.Add(tc)
	}

	callback := func([]byte) {
		// do nothing
	}

	f.Fuzz(func(t *testing.T, json string) {
		count.Tests(t, 1)
		parse([]byte(json), callback)
		// we are just testing we can't cause an unhandled panic
	})
}
