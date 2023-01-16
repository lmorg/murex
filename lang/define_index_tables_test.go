package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
)

func TestCharToIndex(t *testing.T) {
	lower := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z'}
	upper := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}

	count.Tests(t, 2)

	test := func(chars []byte) {
		for i, v := range chars {
			if i != charToIndex(v) {
				t.Errorf("'%s' shouldn't equal %d, expecting %d", string([]byte{v}), charToIndex(v), i)
			}
		}
	}

	test(lower)
	test(upper)
}
