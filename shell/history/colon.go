package history

func noColon(line string) string {
	var escape, qSingle, qDouble, funcStart bool

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
				switch {
				case !funcStart:
					line = line[i+1:]
					//i--
					continue

				// colon mid command - must split
				case i < len(line)-1 && line[i+1] != ' ':
					return line[:i] + " " + line[i+1:]

				default:
					return line[:i] + line[i+1:]
				}
			}
		default:
			funcStart = true
		}
	}

	return line
}
