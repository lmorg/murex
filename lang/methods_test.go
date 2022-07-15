package lang

import (
	"sort"
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestMethodExists(t *testing.T) {
	count.Tests(t, 5)

	m := newMethods()

	cmds := m.dt[t.Name()]
	if m.Exists("1", t.Name()) {
		t.Error("Method exists")
		t.Logf("  cmds: %s", json.LazyLogging(cmds))
		t.Logf("  m.dt: %s", json.LazyLogging(m.dt))
	}

	m.Define("1", t.Name())
	if len(m.dt[t.Name()]) != 1 || m.dt[t.Name()][0] != "1" {
		t.Errorf("m.dt not getting set: %v", json.LazyLogging(m.dt))
	}

	if !m.Exists("1", t.Name()) {
		t.Error("Method does not exist")
		t.Logf("  cmds: %s", json.LazyLogging(cmds))
		t.Logf("  m.dt: %s", json.LazyLogging(m.dt))
	}

	m.Define("2", t.Name())
	if len(m.dt[t.Name()]) != 2 || m.dt[t.Name()][1] != "2" {
		t.Errorf("m.dt not getting set: %v", json.LazyLogging(m.dt))
	}

	if !m.Exists("2", t.Name()) {
		t.Error("Method does not exist")
		t.Logf("  cmds: %s", json.LazyLogging(cmds))
		t.Logf("  m.dt: %s", json.LazyLogging(m.dt))
	}
}

func TestMethodDefine(t *testing.T) {
	count.Tests(t, 1)

	m := newMethods()
	m.Define("1", t.Name())
	if len(m.dt[t.Name()]) != 1 || m.dt[t.Name()][0] != "1" {
		t.Errorf("m.dt not getting set: %v", json.LazyLogging(m.dt))
	}
}

func TestMethods(t *testing.T) {
	count.Tests(t, 11)

	m := newMethods()

	dump := m.Dump()

	if len(dump) != 0 {
		t.Errorf("Init len(m.Dump) != 0: %d", len(dump))
	}

	get := m.Get("test")
	if get == nil {
		t.Errorf("m.Get() should equal []string{} not nil")
	}

	// foo

	m.Define("foo", "test")

	get = m.Get("test")
	if len(get) != 1 || get[0] != "foo" {
		t.Error(`m.Get[0] != "foo":`)
		t.Logf(`  len(get): %d`, len(get))
		t.Logf(`  get:     `+"`%s`", json.LazyLogging(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 1 || dump["test"][0] != "foo" {
		t.Error(`m.Dump["test"][0] != "foo":`)
		t.Logf(`  len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%s`", json.LazyLogging(dump["test"]))
	}

	// bar

	m.Define("bar", "test")

	get = m.Get("test")
	if len(get) != 2 || get[1] != "bar" {
		t.Error(`m.Get[1] != "bar":`)
		t.Logf(`  len(get): %d`, len(get))
		t.Logf(`  get:     `+"`%s`", json.LazyLogging(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 2 || dump["test"][1] != "bar" {
		t.Error(`m.Dump["test"][1] != "foo":`)
		t.Logf(`  len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%v`", json.LazyLogging(dump["test"]))
	}

	// foo (dedup)

	m.Define("foo", "test")

	get = m.Get("test")
	if len(get) != 2 {
		t.Errorf(`len(get) != 2: %d`, len(get))
		t.Logf(`  get:     `+"`%s`", json.LazyLogging(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 2 {
		t.Errorf(`len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%v`", json.LazyLogging(dump["test"]))
	}
}

func sortedSlice(a []string) string {
	sort.Strings(a)
	return json.LazyLogging(a)
}

func TestMethodTypes(t *testing.T) {
	m := newMethods()
	tests := []struct {
		Command  string
		Type     string
		Expected []string
	}{
		{
			Command:  "name",
			Type:     "str",
			Expected: []string{"str"},
		},
		{
			Command:  "age",
			Type:     types.Math,
			Expected: []string{"bool", "int", "num", "float"},
		},
	}

	count.Tests(t, len(tests))

	for _, test := range tests {
		m.Define(test.Command, test.Type)
	}

	m.Degroup()

	for i, test := range tests {
		actual := sortedSlice(m.Types(test.Command))
		expected := sortedSlice(test.Expected)

		if expected != actual {
			t.Errorf("Unexpected return in test %d", i)
			t.Logf("  Expected: %s", expected)
			t.Logf("  Actual:   %s", actual)
		}
	}
}
