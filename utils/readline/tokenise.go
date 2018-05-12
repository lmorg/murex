package readline

func (rl *Instance) tokeniseLine() ([]string, int, int) {
	if len(rl.line) == 0 {
		return nil, 0, 0
	}

	var index, pos int
	split := make([]string, 1)

	for i, r := range rl.line {
		switch {
		case r == '_' || r == '-' || r == ':' ||
			(r >= 'a' && 'z' >= r) ||
			(r >= 'A' && 'Z' >= r) ||
			(r >= '0' && '9' >= r):

			if i > 0 && (rl.line[i-1] == ' ' || rl.line[i-1] == '\t') {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)

		case r == ' ' || r == '\t':
			split[len(split)-1] += string(r)

		default:
			split = append(split, string(r))
		}

		if i == rl.pos {
			index = len(split) - 1
			pos = len(split[index]) - 1
		}
	}

	return split, index, pos
}
