package lists

func Match(a []string, s string) bool {
	for i := range a {
		if a[i] == s {
			return true
		}
	}

	return false
}

func MatchIndexString(a []string, s string) int {
	for i := range a {
		if a[i] == s {
			return i
		}
	}

	return -1
}
