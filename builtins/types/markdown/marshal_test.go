//go:build ignore
// +build ignore

package markdown

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

// Bugfix: https://github.com/lmorg/murex/issues/801
func TestMarshal(t *testing.T) {
	var (
		src      = `[{"Name":"Jake","Department":"Sales"},{"Name":"Carl","Department":"Accounting"},{"Name":"Abigail","Department":"IT"}]`
		expected = "Department,Name\nSales,Jake\nAccounting,Carl\nIT,Abigail\n"
		v        any
	)

	count.Tests(t, 1)

	err := json.Unmarshal([]byte(src), &v)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}

	b, err := marshal(lang.NewTestProcess(), v)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
		return
	}

	if string(b) != expected {
		t.Errorf("Expected != actual:")
		t.Logf("  Expected: %s", strings.ReplaceAll(expected, "\n", `\n`))
		t.Logf("  Actual:   %s", strings.ReplaceAll(string(b), "\n", `\n`))
	}
}
