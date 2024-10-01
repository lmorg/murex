package readline

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestCropCaption(t *testing.T) {
	count.Tests(t, 1)
	// We aren't really bothered about the quality of the output here, just
	// testing that the function doesn't generate any slice out of bounds
	// exceptions

	var caption, maxLen, cellWidth int

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic raised on iteration %d,%d,%d: %s", caption, maxLen, cellWidth, r)
		}
	}()

	for caption = 0; caption < 10; caption++ {
		for maxLen = 0; maxLen < 10; maxLen++ {
			for cellWidth = 0; cellWidth < 10; cellWidth++ {
				_ = cropCaption(strings.Repeat("s", caption), maxLen, cellWidth)
			}
		}
	}
}
