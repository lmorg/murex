package history

import (
	"errors"
)

// TestHistory is a dummy history struct for testing
type TestHistory struct {
	list []string
}

func NewTestHistory() *TestHistory {
	h := new(TestHistory)
	h.list = []string{
		"out: the quick brown #fox",
		"out: jumped over",
		"out: the lazy dog",
	}
	return h
}

// Write item to history file. eg ~/.murex_history
func (h *TestHistory) Write(s string) (int, error) {
	h.list = append(h.list)
	return len(h.list), nil
}

// GetLine returns a specific line from the history file
func (h *TestHistory) GetLine(i int) (string, error) {
	if i < 0 {
		return "", errors.New("Cannot use a negative index when requesting historic commands")
	}
	if i < len(h.list) {
		return h.list[i], nil
	}
	return "", errors.New("Index requested greater than number of items in history")
}

// Len returns the number of items in the history file
func (h *TestHistory) Len() int {
	return len(h.list)
}

// Dump returns the entire history file
func (h *TestHistory) Dump() interface{} {
	return h.list
}
