//go:build js
// +build js

package readline

import "errors"

func (rl *Instance) launchEditor(multiline []rune) ([]rune, error) {
	return rl.line, errors.New("Not currently supported in WebAssembly")
}
