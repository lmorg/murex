package mkarray

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestConsts(t *testing.T) {
	for i, m := range mapRanges {
		for element := range m {

			count.Tests(t, len(m))

			if element != strings.ToLower(element) {
				t.Errorf("mapRange contains a non-lowercase element")
				t.Log("  mapRange:", i)
				t.Log("  map:     ", m)
				t.Log("  element: ", element)
				t.Log("  All elements in a mapRange should be lower case for performance reasons")
			}
		}
	}
}
