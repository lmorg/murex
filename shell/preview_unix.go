//go:build !plan9 && !windows && !js
// +build !plan9,!windows,!js

package shell

import (
	"bytes"
	"os/exec"

	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/readline"
	"github.com/lmorg/murex/utils/rmbs"
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
	b, err := man.GetManPage(exe, size.Width).ReadAll()
	if err != nil {
		return []byte{}
	}

	s := rmbs.Remove(string(b))
	return []byte(s)
}
