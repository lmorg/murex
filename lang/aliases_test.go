package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestAliases(t *testing.T) {
	count.Tests(t, 9)

	a := lang.NewAliases()

	dump := a.Dump()
	if len(dump) != 0 {
		t.Errorf("Empty aliases not equal 0 in dump: %d", len(dump))
	}

	if a.Exists("test") {
		t.Error("test alias appears in empty aliases")
	}

	a.Add("test", []string{"foo", "bar"}, nil)

	dump = a.Dump()
	if len(dump) != 1 {
		t.Errorf("Alias dump not equal 1: %d", len(dump))
	}

	if !a.Exists("test") {
		t.Error("test alias does not appear in aliases")
	}

	s := a.Get("test")
	if len(s) != 2 || s[0] != "foo" || s[1] != "bar" {
		t.Errorf("Aliases not retrieved correctly: len(%v) == %d", s, len(s))
	}

	err := a.Delete("test")
	if err != nil {
		t.Error(err)
	}

	dump = a.Dump()
	if len(dump) != 0 {
		t.Errorf("Empty aliases not equal 0 in dump: %d", len(dump))
	}

	if a.Exists("test") {
		t.Error("test alias appears in empty aliases")
	}
}

// TestAliasArray tests what goes in is the same as what comes out
func TestAliasArray(t *testing.T) {
	count.Tests(t, 1)

	exp := []string{"testing", "$foo", "@bar", "foo bar"}
	a := lang.NewAliases()

	a.Add("test", exp, nil)
	act := a.Get("test")

	if len(exp) != len(act) {
		t.Error("Array len() mismatch:")
		t.Logf("  Expected: %d", len(exp))
		t.Logf("  Actual:   %d", len(act))
	}

	for i := range exp {
		if exp[i] != act[i] {
			t.Errorf("Array item mismatch: %d", i)
			t.Logf("  Expected: %s", exp[i])
			t.Logf("  Actual:   %s", act[i])
		}
	}
}
