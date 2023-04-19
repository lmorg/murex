//go:build !plan9 && !windows && !js
// +build !plan9,!windows,!js

package preview

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/lmorg/murex/utils/readline"
)

func previewFile(filename string) []byte {
	cmd := exec.Command("file", filename)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil
	}

	return bytes.ReplaceAll(out.Bytes(), []byte{',', ' '}, []byte{',', '\n', '>', ' '})
}

func manPage(exe string, size *readline.PreviewSizeT) []byte {
	cmd := exec.Command("man", exe)
	cmd.Env = []string{fmt.Sprintf("MANWIDTH=%d", size.Width)}

	var out, err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err

	if e := cmd.Run(); e != nil {
		return err.Bytes()
	}

	if out.Len() == 0 {
		return err.Bytes()
	}

	return out.Bytes()
}
