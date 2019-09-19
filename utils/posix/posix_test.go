package posix

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

// TestPosix checks isPosix logic is correct
func TestPosix(t *testing.T) {
	platforms := map[string]bool{
		"linux":     true,
		"freebsd":   true,
		"openbsd":   true,
		"netbsd":    true,
		"dragonfly": true,
		"darwin":    true,
		"solaris":   true,
		"windows":   false,
		"plan9":     false,
	}

	count.Tests(t, len(platforms))

	for os, val := range platforms {
		if val {
			if !isPosix(os) {
				t.Errorf("%s not returning as posix. Expecting it to be posix", os)
			}
		} else {
			if isPosix(os) {
				t.Errorf("%s is returning as posix. Expecting it not to", os)
			}
		}
	}
}
