//go:build js
// +build js

package shell

import (
	"github.com/lmorg/readline/v4"
)

func definePromptHistory() {
	// We don't want persistent history when running this from WebAssembly
	Prompt.History = &readline.ExampleHistory{}
}

func setPromptHistory() {
	// We don't want persistent history when running this from WebAssembly
	Prompt.History = &readline.ExampleHistory{}
}
