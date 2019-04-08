package counter

import "sync"

// MutexCounter is a thread-safe counter
type MutexCounter struct {
	i int
	m sync.Mutex
}

// Add increments the counter by 1
func (mc *MutexCounter) Add() int {
	mc.m.Lock()
	mc.i++
	i := mc.i
	mc.m.Unlock()

	return i
}

// Set redefines the counter to a specified value
func (mc *MutexCounter) Set(i int) {
	mc.m.Lock()
	mc.i = i
	mc.m.Unlock()
}

// NotEqual tests if the counters value is the same as the supplied value
func (mc *MutexCounter) NotEqual(i int) bool {
	mc.m.Lock()
	ne := mc.i != i
	mc.m.Unlock()
	return ne
}
