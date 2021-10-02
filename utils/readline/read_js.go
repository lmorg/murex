// +build js

package readline

import (
	"errors"
)

var Stdin = make(chan string, 0)

func read(b []byte) (int, error) {
	stdin := <-Stdin

	if len(stdin) > len(b) {
		return 0, errors.New("wasm keystrokes > b (this is a bug)")
	}

	for i := range stdin {
		b[i] = stdin[i]
	}

	return len(stdin), nil
}
