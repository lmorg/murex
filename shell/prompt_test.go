package shell

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils"
)

func TestLeftMost(t *testing.T) {
	count.Tests(t, 1)

	b := leftMost()

	if string(b) == utils.NewLineString {
		return
	}

	w := strconv.Itoa(len(b) - 1)
	s := fmt.Sprintf("%"+w+"s\r", " ")

	if s != string(b) {
		t.Error("Unexpected return")
		t.Log("  Expected: ", []byte(s))
		t.Log("  Actual:   ", b)
	}
}
