package counter

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestCounter(t *testing.T) {
	count.Tests(t, 2)

	mc := new(MutexCounter)
	i := mc.Add()
	if i != 1 {
		t.Errorf("MutexCounter should eq 1: %d", i)
	}

	mc.Set(2)
	if !mc.NotEqual(i) {
		t.Errorf("MutexCounter should eq 2: %d != 2", i)
	}
}
