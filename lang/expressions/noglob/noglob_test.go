package noglob

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestCanGlobCmd(t *testing.T) {
	falses := []string{"foo", "bar"}
	trues := append(noGlobCmds, "cast", "format", "select", "find", "regexp")

	count.Tests(t, len(trues)+len(falses))

	for i := range trues {
		v := canGlobCmd(trues[i])
		if v != true {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "false", "true", trues[i])
		}
	}

	for i := range falses {
		v := canGlobCmd(falses[i])
		if v != false {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "true", "false", falses[i])
		}
	}
}

func TestReadNoGlobCmds(t *testing.T) {
	count.Tests(t, 1)

	v, err := ReadNoGlobCmds()
	if err != nil {
		t.Error(err.Error())
	}

	switch v.(type) {
	case []string:
		// success
	default:
		t.Errorf("incorrect type. Should be a []string")
	}
}

func TestWriteNoGlobCmds(t *testing.T) {
	count.Tests(t, 5)

	noGlobCmdsDefault := make([]string, len(noGlobCmds))
	copy(noGlobCmdsDefault, noGlobCmds)
	defer copy(noGlobCmds, noGlobCmdsDefault)

	var v any

	v = 13
	err := WriteNoGlobCmds(v)
	if err == nil || !strings.Contains(err.Error(), "invalid data-type") {
		t.Error("invalid data-type accepted:", err)
	}

	v = []string{"foo", "bar"}
	err = WriteNoGlobCmds(v)
	if err == nil || !strings.Contains(err.Error(), "invalid data-type") {
		t.Error("invalid data-type accepted:", err)
	}

	v = "foobar"
	err = WriteNoGlobCmds(v)
	if err == nil {
		t.Error("value not marshalled as JSON:", err)
	}

	v = `{"foo": "bar"}`
	err = WriteNoGlobCmds(v)
	if err == nil {
		t.Error("invalid JSON object accepted:", err)
	}

	v = `["foo"]`
	err = WriteNoGlobCmds(v)
	if err != nil {
		t.Error("JSON object should have be valid:", err)
	}
	if len(noGlobCmds) != 1 || noGlobCmds[0] != "foo" {
		t.Errorf("noGlobCmds incorrectly set: %s", json.LazyLogging(noGlobCmds))
	}
}
