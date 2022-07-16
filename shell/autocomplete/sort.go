//go:build ignore
// +build ignore

package autocomplete

func sortCompletions(items []string) {
	sortCompletionsR(items, 0, len(items)-1)
}

func sortCompletionsR(items []string, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := items[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if isLt(noColon(items[i]), noColon(pivot)) {
			temp := items[splitIndex]

			items[splitIndex] = items[i]
			items[i] = temp

			splitIndex++
		}
	}

	items[end] = items[splitIndex]
	items[splitIndex] = pivot

	sortCompletionsR(items, start, splitIndex-1)
	sortCompletionsR(items, splitIndex+1, end)
}

func isLt(a, b string) bool {
	switch {
	case len(a) == 0 || len(b) == 0:
		return a < b

	case a[0] == '-':
		if b[0] == '-' {
			return a < b
		}
		return false

	case b[0] == '-':
		return true

	default:
		return a < b
	}
}

func noColon(s string) string {
	if len(s) > 1 && s[len(s)-1] == ':' {
		s = s[:len(s)-1]
	}
	return s
}
