//go:build !js
// +build !js

package readline

import (
	"fmt"
	"os"

	"github.com/lmorg/murex/debug"
)

func read(b []byte) (int, error) {
	i, err := os.Stdin.Read(b)

	if err != nil && debug.Enabled {
		s := fmt.Sprintf("!!! cannot read from stdin: %s", err.Error())
		debug.Log(s)
		panic(s)
	}

	return i, err
}
