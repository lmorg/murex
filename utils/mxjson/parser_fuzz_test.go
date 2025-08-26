package mxjson_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/mxjson"
)

func FuzzParser(f *testing.F) {
	tests := []string{``, "[\n  1,\n  2,\n  3\n]", `{ "key: "#value" }`, `{ "key": ({ value }) }`}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		count.Tests(t, 1)
		mxjson.Parse([]byte(orig))
		// we are just testing we can't cause an unhandled panic
	})
}
