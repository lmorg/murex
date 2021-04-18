package inject

import "fmt"

// String injects one string into another at a given postion
func String(old, insert string, pos int) (string, error) {
	switch {
	case len(old) == 0:
		if pos == 0 {
			return insert, nil
		} else {
			return "", fmt.Errorf("pos cannot be non-zero when old is empty")
		}

	case pos < 0:
		return "", fmt.Errorf("pos cannot be less than zero")

	case len(old) < pos:
		return "", fmt.Errorf("Len of old is less than pos")

	case pos == 0:
		return insert + old, nil

	case pos == len(old):
		return old + insert, nil

	default:
		return old[:pos] + insert + old[pos:], nil
	}
}

// Rune injects one []rune into another at a given postion
func Rune(old, insert []rune, pos int) ([]rune, error) {
	switch {
	case len(old) == 0:
		if pos == 0 {
			return insert, nil
		} else {
			return []rune{}, fmt.Errorf("pos cannot be non-zero when old is empty")
		}

	case pos < 0:
		return []rune{}, fmt.Errorf("pos cannot be less than zero")

	case len(old) < pos:
		return []rune{}, fmt.Errorf("Len of old is less than pos")

	case pos == 0:
		return append(insert, old...), nil

	case pos == len(old):
		return append(old, insert...), nil

	default:
		new := make([]rune, len(old)+len(insert))
		for i := 0; i < pos; i++ {
			new[i] = old[i]
		}
		for i := 0; i < len(insert); i++ {
			new[pos+i] = insert[i]
		}
		l := len(insert)
		for i := pos; i < len(old); i++ {
			new[l+i] = old[i]
		}

		return new, nil
	}
}
