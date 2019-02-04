package counter

import "testing"

func TestCounter(t *testing.T) {
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
