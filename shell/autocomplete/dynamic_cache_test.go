package autocomplete

import (
	"testing"
	"time"

	"github.com/lmorg/murex/test/count"
)

func TestDynamicCache(t *testing.T) {
	count.Tests(t, 5)

	cache := NewDynamicCache()

	var (
		exe      = "bob"
		params   = []string{"foo", "bar", "baz"}
		block    = []rune("out: hello world")
		stdout   = []byte("hello world")
		dataType = "str"
		hash     = cache.CreateHash(exe, params, block)
	)

	cache.Set(hash, stdout, dataType, -1)
	out, _ := cache.Get(hash)
	if len(out) > 0 {
		t.Errorf("out should be empty: '%s'", string(out))
	}

	cache.Set(hash, stdout, dataType, 10)
	out, _ = cache.Get(hash)
	if string(out) != string(stdout) {
		t.Errorf("out doesn't match expected: '%s'", string(out))
	}

	cache.mutex.Lock()
	cache.gcSleep = 0
	cache.mutex.Unlock()
	go cache.garbageCollection()

	time.Sleep(1 * time.Second)
	cache.mutex.Lock()
	if cache.hash[string(hash)].dataType != "str" {
		t.Errorf("dataType should be set")
	}
	cache.mutex.Unlock()

	cache.Set(hash, stdout, dataType, -1)
	time.Sleep(1 * time.Second)
	cache.mutex.Lock()
	if cache.hash[string(hash)].dataType != "" {
		t.Errorf("dataType should be unset")
	}
	cache.mutex.Unlock()
}
