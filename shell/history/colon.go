package history

func noColon(line string) string {
	var escape, qSingle, qDouble bool

	for i := range line {
		switch line[i] {
		case '#':
			return line
		case '\\':
			switch {
			case escape:
				escape = false
			case qSingle:
				// do nothing
			default:
				escape = true
			}
		case '\'':
			switch {
			case qDouble, escape:
				escape = false
			default:
				qSingle = !qSingle
			}
		case '"':
			switch {
			case qSingle, escape:
				escape = false
			default:
				qDouble = !qDouble
			}
		case '{':
			if !escape && !qSingle && !qDouble {
				return line
			}
		case '\r', '\n', '\t', ' ':
			if !escape && !qSingle && !qDouble {
				return line
			}
		case ':':
			if !escape && !qSingle && !qDouble {
				return line[:i] + line[i+1:]
			}
		}
	}

	return line
}
