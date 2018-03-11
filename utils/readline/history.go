package readline

// LineHistory is an interface to allow you to write your own history logging
// tools. eg sqlite backend instead of a file system.
// By default readline will just use the dummyLineHistory interface which only
// logs the history to memory ([]string to be precise).
type LineHistory interface {
	// Append takes the line and returns an updated number of lines or an error
	Append(string) (int, error)
	// GetLine takes the historic line number and returns the line or an error
	GetLine(int) (string, error)
	// Len returns the number of history lines
	Len() int
}

// Example LineHistory interface.

type dummyLineHistory struct {
	items []string
}

func (h *dummyLineHistory) Append(s string) (int, error) {
	h.items = append(h.items, s)
	return len(h.items), nil
}

func (h *dummyLineHistory) GetLine(i int) (string, error) {
	return h.items[i], nil
}

func (h *dummyLineHistory) Len() int {
	return len(h.items)
}
