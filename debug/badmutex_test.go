package debug

import (
	"sync"
	"testing"
	"time"

	"github.com/lmorg/murex/test/count"
)

// TestBadMutex proves our test bad mutex (used to help diagnose locking faults)
// does not lock
func TestBadMutex(t *testing.T) {
	count.Tests(t, 1, "TestBadMutex")

	var (
		m      BadMutex // if we swap this for sync.Mutex the error should be raised
		exited bool
	)

	m.Lock()

	go func() {
		time.Sleep(500 * time.Millisecond)
		m.Unlock()
		if !exited {
			t.Error("BadMutex caused a locking condition. This should not happen")
		}
	}()

	m.Lock()
	exited = true
}

// TestGoodMutex proves our bad mutex test works
func TestGoodMutex(t *testing.T) {
	count.Tests(t, 1, "TestGoodMutex")

	var (
		m      sync.Mutex
		exited bool
	)

	m.Lock()

	go func() {
		time.Sleep(500 * time.Millisecond)
		m.Unlock()
		if exited {
			t.Error("Mutex did not cause a locking condition. The test logic has failed")
		}
	}()

	m.Lock()
	exited = true
}
