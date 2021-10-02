package process

import "sync/atomic"

type Background struct {
	i int32
}

func (bg *Background) Get() bool {
	return atomic.LoadInt32(&bg.i) != 0
}

func (bg *Background) Set(v bool) {
	if v {
		atomic.StoreInt32(&bg.i, 1)
		return
	}

	atomic.StoreInt32(&bg.i, 0)
}

func (bg *Background) String() string {
	if bg.Get() {
		return "yes"
	}

	return "no"
}
