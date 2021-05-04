// +build js

package readline

// State contains the state of a terminal.
type State struct{}

// MakeRaw is an empty function for web assembly since the terminal is just a
// webpage
func MakeRaw(_ int) (*State, error) {
	return new(State), nil
}

// Restore is an empty function for web assembly since the terminal is just a
// webpage
func Restore(_ int, _ *State) error {
	return nil
}

// GetSize the default terminal size in the webpage
func GetSize(_ int) (width, height int, err error) {
	return 120, 40, nil
}
