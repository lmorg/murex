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

	a.Add("test", []string{"foo", "bar"})

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
