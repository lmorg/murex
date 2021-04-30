package humannumbers

import "strconv"

// Ordinal returns the number + ordinal suffix
func Ordinal(i int) string {
	s := strconv.Itoa(i)

	switch i % 100 {
	case 11, -11, 12, -12:
		return s + "th"
	}

	switch i % 10 {
	case 1, -1:
		return s + "st"
	case 2, -2:
		return s + "nd"
	case 3, -3:
		return s + "rd"
	default:
		return s + "th"
	}
}
