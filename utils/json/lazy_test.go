package json_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func TestLazyLogging(t *testing.T) {
	count.Tests(t, 3)

	act := json.LazyLogging(nil)
	if act != "null" {
		t.Errorf("Actual != 'null': %s", act)
	}

	act = json.LazyLogging(3)
	if act != "3" {
		t.Errorf("Actual != 3 : %s", act)
	}

	act = json.LazyLogging("foobar")
	if act != `"foobar"` {
		t.Errorf(`Actual != '"foobar"' : %s`, act)
	}
}
