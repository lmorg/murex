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

	sort.Strings(s)

	j := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			s[j] = s[i+1]
			j++
		}
	}

	return j
}
