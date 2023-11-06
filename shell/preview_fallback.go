//go:build plan9 || windows || js
// +build plan9 windows js

package shell

import (
	"github.com/lmorg/murex/utils/readline"
)

func previewFile(filename string) []byte {
	return nil
}

func manPage(_ string, _ *readline.PreviewSizeT) []byte {
	return nil
}
