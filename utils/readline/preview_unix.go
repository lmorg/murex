//go:build !plan9 && !windows && !js
// +build !plan9,!windows,!js

package readline

import (
	"bytes"
	"os/exec"
)

func previewFile(filename string) []byte {
	cmd := exec.Command("file", filename)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil
	}

	return out.Bytes()
}
