package lang

import (
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func quickJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func TestMethodExists(t *testing.T) {
	count.Tests(t, 5)

	m := newMethods()

	cmds := m.dt[t.Name()]
	i := m.methodExists("1", t.Name())
	if i != -1 {
		t.Errorf("methodExists != -1: %d", i)
		t.Logf("  cmds: %s", quickJson(cmds))
		t.Logf("  m.dt: %s", quickJson(m.dt))
	}

	m.Define("1", t.Name())
	if len(m.dt[t.Name()]) != 1 || m.dt[t.Name()][0] != "1" {
		t.Errorf("m.dt not getting set: %v", quickJson(m.dt))
	}

	i = m.methodExists("1", t.Name())
	if i != 0 {
		t.Errorf("methodExists != 0: %d", i)
		t.Logf("  cmds: %s", quickJson(cmds))
		t.Logf("  m.dt: %s", quickJson(m.dt))
	}

	m.Define("2", t.Name())
	if len(m.dt[t.Name()]) != 2 || m.dt[t.Name()][1] != "2" {
		t.Errorf("m.dt not getting set: %v", quickJson(m.dt))
	}

	i = m.methodExists("2", t.Name())
	if i != 1 {
		t.Errorf("methodExists != 1: %d", i)
		t.Logf("  cmds: %s", quickJson(cmds))
		t.Logf("  m.dt: %s", quickJson(m.dt))
	}
}

func TestMethodDefine(t *testing.T) {
	count.Tests(t, 1)

	m := newMethods()
	m.Define("1", t.Name())
	if len(m.dt[t.Name()]) != 1 || m.dt[t.Name()][0] != "1" {
		t.Errorf("m.dt not getting set: %v", quickJson(m.dt))
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
		t.Logf(`  get:     `+"`%s`", quickJson(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 1 || dump["test"][0] != "foo" {
		t.Error(`m.Dump["test"][0] != "foo":`)
		t.Logf(`  len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%s`", quickJson(dump["test"]))
	}

	// bar

	m.Define("bar", "test")

	get = m.Get("test")
	if len(get) != 2 || get[1] != "bar" {
		t.Error(`m.Get[1] != "bar":`)
		t.Logf(`  len(get): %d`, len(get))
		t.Logf(`  get:     `+"`%s`", quickJson(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 2 || dump["test"][1] != "bar" {
		t.Error(`m.Dump["test"][1] != "foo":`)
		t.Logf(`  len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%v`", quickJson(dump["test"]))
	}

	// foo (dedup)

	m.Define("foo", "test")

	get = m.Get("test")
	if len(get) != 2 {
		t.Errorf(`len(get) != 2: %d`, len(get))
		t.Logf(`  get:     `+"`%s`", quickJson(get))
	}

	dump = m.Dump()
	if len(dump["test"]) != 2 {
		t.Errorf(`len(dump["test"]): %d`, len(dump["test"]))
		t.Logf(`  dump["test"]:     `+"`%v`", quickJson(dump["test"]))
	}
}
