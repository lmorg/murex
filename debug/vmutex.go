package debug

import (
	"fmt"
	"runtime"
	"sync"
)

// VMutex is a verbose mutex used to test deadlocks and shouldn't be used in release code.
type VMutex struct {
	m sync.Mutex
}

// Lock is a wrapper around a mutex Lock() to help locate deadlocks
func (v *VMutex) Lock() {
	_, file1, line1, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	Log(fmt.Sprintf("( LOCK ) %s:%d, %s:%d, ", file1, line1, file2, line2))
	v.m.Lock()
}

// Unlock is a wrapper around a mutex Unlock() to help locate deadlocks
func (v *VMutex) Unlock() {
	_, file1, line1, _ := runtime.Caller(1)
	_, file2, line2, _ := runtime.Caller(2)
	Log(fmt.Sprintf("(UNLOCK) %s:%d, %s:%d, ", file1, line1, file2, line2))
	v.m.Unlock()
}
