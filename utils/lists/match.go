package lists

func Match(a []string, s string) bool {
	for i := range a {
		if a[i] == s {
			return true
		}
	}

	return false
}
