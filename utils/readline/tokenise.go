package readline

import "strings"

func tokeniseLine(rl *Instance) ([]string, int, int) {
	if len(rl.line) == 0 {
		return nil, 0, 0
	}

	var index, pos int
	var punc bool

	split := make([]string, 1)

	for i, r := range rl.line {
		switch {
		case (r >= 33 && 47 >= r) ||
			(r >= 58 && 64 >= r) ||
			(r >= 91 && 94 >= r) ||
			r == 96 ||
			(r >= 123 && 126 >= r):

			if i > 0 && rl.line[i-1] != r {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
			punc = true

		case r == ' ' || r == '\t':
			split[len(split)-1] += string(r)
			punc = true

		default:
			if punc {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
			punc = false
		}

		if i == rl.pos {
			index = len(split) - 1
			pos = len(split[index]) - 1
		}
	}

	return split, index, pos
}

func tokeniseSplitSpaces(rl *Instance) ([]string, int, int) {
	if len(rl.line) == 0 {
		return nil, 0, 0
	}

	var index, pos int
	split := make([]string, 1)

	for i, r := range rl.line {
		switch {
		case r == ' ' || r == '\t':
			split[len(split)-1] += string(r)

		default:
			if i > 0 && (rl.line[i-1] == ' ' || rl.line[i-1] == '\t') {
				split = append(split, "")
			}
			split[len(split)-1] += string(r)
		}

		if i == rl.pos {
			index = len(split) - 1
			pos = len(split[index]) - 1
		}
	}

	return split, index, pos
}

func tokeniseBrackets(rl *Instance) ([]string, int, int) {
	var (
		open, close    rune
		split          []string
		count          int
		pos            map[int]int = make(map[int]int)
		match          int
		single, double bool
	)

	switch rl.line[rl.pos] {
	case '(', ')':
		open = '('
		close = ')'

	case '{', '[':
		open = rl.line[rl.pos]
		close = rl.line[rl.pos] + 2

	case '}', ']':
		open = rl.line[rl.pos] - 2
		close = rl.line[rl.pos]

	default:
		return nil, 0, 0
	}

	for i := range rl.line {
		switch rl.line[i] {
		case '\'':
			if !single {
				double = !double
			}

		case '"':
			if !double {
				single = !single
			}

		case open:
			if !single && !double {
				count++
				pos[count] = i
				if i == rl.pos {
					match = count
					split = []string{string(rl.line[:i-1])}
				}

			} else if i == rl.pos {
				return nil, 0, 0
			}

		case close:
			if !single && !double {
				if match == count {
					split = append(split, string(rl.line[pos[count]:i]))
					return split, 1, 0
				}
				if i == rl.pos {
					split = []string{
						string(rl.line[:pos[count]-1]),
						string(rl.line[pos[count]:i]),
					}
					return split, 1, len(split[1])
				}
				count--

			} else if i == rl.pos {
				return nil, 0, 0
			}
		}
	}

	return nil, 0, 0
}

func rTrimWhiteSpace(oldString string) (newString string) {
	return strings.TrimRight(oldString, " ")
	// TODO: support tab chars
	/*defer fmt.Println(">" + oldString + "<" + newString + ">")
	newString = oldString
	for len(oldString) > 0 {
		if newString[len(newString)-1] == ' ' || newString[len(newString)-1] == '\t' {
			newString = newString[:len(newString)-1]
		} else {
			break
		}
	}
	return*/
}
