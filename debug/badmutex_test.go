package debug

import (
	"testing"
	"time"
)

// TestBadMutex proves our test bad mutex (used to help diagnose locking faults)
// does not lock
func TestBadMutex(t *testing.T) {
	var (
		m      BadMutex // if we swap this for sync.Mutex the error should be raised
		exited bool
	)

	m.Lock()

	go func() {
		time.Sleep(2 * time.Second)
		m.Unlock()
		if !exited {
			t.Error("BadMutex caused a locking condition. This should not happen")
		}
	}()

	m.Lock()
	exited = true
}
