package debug

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/lmorg/murex/test/count"
)

// TestBadMutex proves our test bad mutex (used to help diagnose locking faults)
// does not lock
func TestBadMutex(t *testing.T) {
	count.Tests(t, 1)

	var (
		m BadMutex // if we swap this for sync.Mutex the error should be raised
		i int32
	)

	go func() {
		m.Lock()
		time.Sleep(500 * time.Millisecond)
		atomic.AddInt32(&i, 1)
		m.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)
	m.Lock()
	m.Unlock()

	if atomic.LoadInt32(&i) != 0 {
		t.Error("BadMutex caused a locking condition. This should not happen")
	}
}

// TestGoodMutex proves our bad mutex test works
func TestGoodMutex(t *testing.T) {
	count.Tests(t, 1)

	var (
		m sync.Mutex // if we swap this for sync.Mutex the error should be raised
		i int32
	)

	go func() {
		m.Lock()
		time.Sleep(500 * time.Millisecond)
		atomic.AddInt32(&i, 1)
		m.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)
	m.Lock()
	m.Unlock()

	if atomic.LoadInt32(&i) == 0 {
		t.Error("Mutex did not cause a locking condition. The test logic has failed")
	}
}
