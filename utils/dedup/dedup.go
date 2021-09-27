package dedup

import "sort"

// SortAndDedupString takes a slice of strings, sorts it and then returns an
// integer for the new size. The existing slice is modified however neither its
// length nor capacity is altered. Hence the return value being an integer.
// For example
//
//     s := []string{"a", "f", "f", "c", "g", "d", "a", "b", "e", "a", "b", "b"}
//     i := dedup.SortAndDedupString(s)
//     fmt.Println(s)     // [a b c d e f g d e f f g]
//     fmt.Println(s[:i]) // [a b c d e f g]
func SortAndDedupString(s []string) int {
	if len(s) == 0 {
		return 0
	}

	sort.Slice(s, func(i, j int) bool {
		switch {
		case len(s[i]) == 0:
			return true

		case len(s[j]) == 0:
			return false

		case s[i][0] == '-' && s[j][0] != '-':
			return false

		case s[i][0] != '-' && s[j][0] == '-':
			return true

		case len(s[i]) < len(s[j]):
			for pos := range s[i] {
				switch {
				case s[i][pos] == ':':
					return true
				//case s[j][pos] == ':':
				//	return false
				case s[i][pos] == s[j][pos]:
					continue
				default:
					return s[i][pos] < s[j][pos]
				}
			}
			return true

		case len(s[i]) > len(s[j]):
			for pos := range s[j] {
				switch {
				//case s[i][pos] == ':':
				//	return true
				case s[j][pos] == ':':
					return false
				case s[i][pos] == s[j][pos]:
					continue
				default:
					return s[i][pos] < s[j][pos]
				}
			}
			return false

		default:
			for pos := range s[i] {
				switch {
				//case s[i][pos] == ':':
				//	return true
				//case s[j][pos] == ':':
				//	return false
				case s[i][pos] == s[j][pos]:
					continue
				default:
					return s[i][pos] < s[j][pos]
				}
			}
			return true
		}
	})

	j := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			s[j] = s[i+1]
			j++
		}
	}

	return j
}
