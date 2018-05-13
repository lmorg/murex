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
			(r >= 91 && 96 >= r) ||
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
