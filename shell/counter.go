package shell

import "sync"

type mutexCounter struct {
	i int
	m sync.Mutex
}

func (mc *mutexCounter) Add() int {
	mc.m.Lock()
	defer mc.m.Unlock()

	mc.i++
	return mc.i
}

func (mc *mutexCounter) Set(i int) {
	mc.m.Lock()
	mc.i = i
	mc.m.Unlock()
}

func (mc *mutexCounter) NotEqual(i int) bool {
	mc.m.Lock()
	defer mc.m.Unlock()

	//debug.Log(mc.i, i)

	if mc.i != i {
		return true
	}

	return false
}
