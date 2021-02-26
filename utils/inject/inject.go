package inject

import "fmt"

// String injects one string into another at a given postion
func String(old, new string, pos int) (string, error) {
	if pos < 0 {
		return "", fmt.Errorf("pos cannot be less than zero")
	}

	if len(old) < pos {
		return "", fmt.Errorf("Len of string is less than pos")
	}

	if pos == 0 {
		return new + old, nil
	}

	if pos == len(old) {
		return old + new, nil
	}

	return old[:pos] + new + old[pos:], nil
}
