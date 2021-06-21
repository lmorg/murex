// +build js

package readline

import "github.com/lmorg/murex/utils/virtualterm"

// VTern is a virtual terminal
var VTerm = virtualterm.NewTerminal(120, 40)

type State struct {
	state virtualterm.PtyState
}

func MakeRaw(_ int) (*State, error) {
	state := State{state: VTerm.MakeRaw()}
	return &state, nil
}

func Restore(_ int, state *State) error {
	VTerm.Restore(state.state)
	return nil
}

// GetSize the default terminal size in the webpage
func GetSize(_ int) (width, height int, err error) {
	return VTerm.GetSize()
}
