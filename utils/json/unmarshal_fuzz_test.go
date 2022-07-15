//go:build go1.18
// +build go1.18

package json

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func FuzzParser(f *testing.F) {
	tests := []string{``, "[\n  1,\n  2,\n  3\n]", `{ "key: "#value" }`, `{ "key": ({ value }) }`}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, data string) {
		count.Tests(t, 1)
		unmarshalMurex([]byte(data), nil)
		// we are just testing we can't cause an unhandled panic
	})
}
