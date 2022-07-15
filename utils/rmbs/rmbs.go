package rmbs

// Remove back spaces from a string
func Remove(s string) string {
	r := []rune(s)
	//l := len(r)
	p := 0

	for i := range r {
		if r[i] == 8 && p > 0 {
			p--
			continue
		}
		r[p] = r[i]
		p++
	}

	return string(r[:p])
}
