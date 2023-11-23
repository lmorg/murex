package autocomplete

import (
	"testing"
	"time"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/cache"
)

func TestDynamicCache(t *testing.T) {
	count.Tests(t, 5)

	dc := NewDynamicCache()

	var (
		exe      = "bob"
		params   = []string{"foo", "bar", "baz"}
		block    = []rune("out: hello world")
		stdout   = []byte("hello world")
		dataType = "str"
		hash     = cache.CreateHash(exe, params, block)
	)

	dc.Set(hash, stdout, dataType, -1)
	out, _ := dc.Get(hash)
	if len(out) > 0 {
		t.Errorf("out should be empty: '%s'", string(out))
	}

	dc.Set(hash, stdout, dataType, 10)
	out, _ = dc.Get(hash)
	if string(out) != string(stdout) {
		t.Errorf("out doesn't match expected: '%s'", string(out))
	}

	dc.mutex.Lock()
	dc.gcSleep = 0
	dc.mutex.Unlock()
	go dc.garbageCollection()

	time.Sleep(1 * time.Second)
	dc.mutex.Lock()
	if dc.hash[string(hash)].DataType != "str" {
		t.Errorf("dataType should be set")
	}
	dc.mutex.Unlock()

	dc.Set(hash, stdout, dataType, -1)
	time.Sleep(1 * time.Second)
	dc.mutex.Lock()
	if dc.hash[string(hash)].DataType != "" {
		t.Errorf("dataType should be unset")
	}
	dc.mutex.Unlock()
}
