package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
)

func TestMimeToMurex_appPlusJson(t *testing.T) {
	type testT struct {
		Mime     string
		DataType string
	}

	tests := []testT{
		{
			Mime:     "application/json",
			DataType: types.Json,
		},
		{
			Mime:     "application/vnd.contentful.management.v1+json",
			DataType: types.Json,
		},
		{
			Mime:     "text/bob",
			DataType: types.String,
		},
		{
			Mime:     "foo/bar",
			DataType: types.Generic,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		dt := lang.MimeToMurex(test.Mime)

		if dt != test.DataType {
			t.Errorf("Mime convertion failed in test %d", i)
			t.Logf("Mime:     '%s'", test.Mime)
			t.Logf("Expected: '%s'", test.DataType)
			t.Logf("Actual:   '%s'", dt)
		}
	}
}
