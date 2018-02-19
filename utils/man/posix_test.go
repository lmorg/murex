// +build !windows

package man

import (
	"testing"
)

// TestMan tests the builtins package
func TestMan(t *testing.T) {
	flags := ScanManPages("cat")
	if len(flags) == 0 {
		t.Error("No flags returned for `cat`")
	}
}
