//go:build plan9 || windows || js
// +build plan9 windows js

package shell

import (
	"github.com/lmorg/readline/v4"
)

func previewFile(filename string) []byte {
	return nil
}

func manPage(_ string, _ *readline.PreviewSizeT) []byte {
	return nil
}
