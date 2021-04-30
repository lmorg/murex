package autocomplete

// isSpecialBuiltin identifies special builtins
func isSpecialBuiltin(s string) bool {
	switch s {
	case ">", ">>", "[", "![", "[[", "@[", "=", "(", "!", ".", "@g":
		return true
	default:
		return false
	}
}

/*func sortColon(items []string, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := items[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if noColon(items[i]) < noColon(pivot) {
			temp := items[splitIndex]

			items[splitIndex] = items[i]
			items[i] = temp

			splitIndex++
		}
	}

	items[end] = items[splitIndex]
	items[splitIndex] = pivot

	sortColon(items, start, splitIndex-1)
	sortColon(items, splitIndex+1, end)
}

func noColon(s string) string {
	if len(s) > 1 && s[len(s)-1] == ':' {
		s = s[:len(s)-1]
	}
	return s
}*/
