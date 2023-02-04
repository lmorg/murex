//go:build plan9 && windows && js
// +build plan9,windows,js

package readline

import (
	"bytes"
	"os/exec"
)

func previewFile(filename string) []byte {
	return nil
}
