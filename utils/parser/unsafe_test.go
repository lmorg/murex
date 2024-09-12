package parser

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestIsCmdUnsafe(t *testing.T) {
	trues := []string{">", ">>", "$var", "rm"}
	falses := append(safeCmds,
		"open", "regexp", "match",
		"cast", "format", "[", "[[",
		"runtime",
	)

	count.Tests(t, len(trues)+len(falses))

	for i := range trues {
		v := isCmdUnsafe(trues[i])
		if v != true {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "false", "true", trues[i])
		}
	}

	for i := range falses {
		v := isCmdUnsafe(falses[i])
		if v != false {
			t.Errorf("Returned `%s` expected `%s`: '%s'", "true", "false", falses[i])
		}
	}
}

func TestReadSafeCmds(t *testing.T) {
	count.Tests(t, 1)

	v, err := ReadSafeCmds()
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

func TestWriteSafeCmds(t *testing.T) {
	count.Tests(t, 5)

	safeCmdsDefault := make([]string, len(safeCmds))
	copy(safeCmdsDefault, safeCmds)
	defer copy(safeCmds, safeCmdsDefault)

	var v interface{}

	v = 13
	err := WriteSafeCmds(v)
	if err == nil || !strings.Contains(err.Error(), "invalid data-type") {
		t.Error("invalid data-type accepted:", err)
	}

	v = []string{"foo", "bar"}
	err = WriteSafeCmds(v)
	if err == nil || !strings.Contains(err.Error(), "invalid data-type") {
		t.Error("invalid data-type accepted:", err)
	}

	v = "foobar"
	err = WriteSafeCmds(v)
	if err == nil {
		t.Error("value not marshalled as JSON:", err)
	}

	v = `{"foo": "bar"}`
	err = WriteSafeCmds(v)
	if err == nil {
		t.Error("invalid JSON object accepted:", err)
	}

	v = `["foo"]`
	err = WriteSafeCmds(v)
	if err != nil {
		t.Error("JSON object should have be valid:", err)
	}
	if len(safeCmds) != 1 || safeCmds[0] != "foo" {
		t.Errorf("safeCmds incorrectly set: %s", json.LazyLogging(safeCmds))
	}
}
