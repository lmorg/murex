package process

import "sync/atomic"

type AtomicBool struct {
	i int32
}

func (ab *AtomicBool) Get() bool {
	return atomic.LoadInt32(&ab.i) != 0
}

func (ab *AtomicBool) Set(v bool) {
	if v {
		atomic.StoreInt32(&ab.i, 1)
		return
	}

	atomic.StoreInt32(&ab.i, 0)
}

func (ab *AtomicBool) String() string {
	if ab.Get() {
		return "yes"
	}

	return "no"
}
