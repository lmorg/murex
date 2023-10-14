package debug

import (
	"fmt"
	"runtime"
	"sync"
)

// BadMutex is only used to test deadlocks and shouldn't be used in release code.
type VLock struct {
	m sync.Mutex
}

// Lock is a wrapper around a mutex Lock() to help locate deadlocks
func (v *VLock) Lock() {
	_, file1, line1, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	Log(fmt.Sprintf("( LOCK ) %s:%d, %s:%d, ", file1, line1, file2, line2))
	v.m.Lock()
}

// Lock is a wrapper around a mutex Unlock() to help locate deadlocks
func (v *VLock) Unlock() {
	_, file1, line1, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	Log(fmt.Sprintf("(UNLOCK) %s:%d, %s:%d, ", file1, line1, file2, line2))
	v.m.Unlock()
}
