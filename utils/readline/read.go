package readline

func removeNonPrintableChars(s []byte) int {
	var (
		i    int
		next int
	)

	for next = 0; next < len(s); next++ {
		if s[next] < ' ' && s[next] != charEOF && s[next] != charEscape &&
			s[next] != charTab && s[next] != charBackspace {

			continue

		} else {
			s[i] = s[next]
			i++
		}
	}

	return i
}
