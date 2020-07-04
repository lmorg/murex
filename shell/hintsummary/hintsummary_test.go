package hintsummary

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestHintSummary tests the HintSummary structure
func TestHintSummary(t *testing.T) {
	count.Tests(t, 6)

	summary := New()

	summary.Set("cmd1", "sum1")
	if summary.Get("cmd1") != "sum1" {
		t.Error("Get (1) returns the wrong string")
	}

	summary.Set("cmd2", "sum2")
	if summary.Get("cmd2") != "sum2" {
		t.Error("Get (2) returns the wrong string")
	}

	summary.Set("cmd3", "sum3")
	if summary.Get("cmd3") != "sum3" {
		t.Error("Get (3) returns the wrong string")
	}

	if len(summary.Dump()) != 3 {
		t.Error("length of summary map is incorrect")
	}

	err := summary.Delete("cmd1")
	if err != nil {
		t.Error(err.Error())
	}

	if len(summary.Dump()) != 2 {
		t.Error("length of summary map is incorrect")
	}
}
