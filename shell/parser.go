package shell

import (
	"strings"
)

var murexCompleter *MurexCompleter = new(MurexCompleter)

type MurexCompleter struct{}

func (fz MurexCompleter) Do(line []rune, pos int) (suggest [][]rune, retPos int) {
	var (
		loc        int = -1
		escaped    bool
		qSingle    bool
		qDouble    bool
		bracket    int
		expectFunc bool = true
		readFunc   bool
		funcName   string
	)

	for i := range line {
		switch line[i] {
		case '#':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `#`
				}
				escaped = false
			case qSingle, qDouble:
				if readFunc {
					funcName += `#`
				}
				escaped = false
			default:
				return
			}

		case '\\':
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `\`
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += `\`
				}
			default:
				escaped = true
			}

		case '\'':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `'`
				}
			case qDouble:
				if readFunc {
					funcName += `'`
				}
			case qSingle:
				qSingle = false
			default:
				qSingle = true
			}

		case '"':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `"`
				}
			case qSingle:
				if readFunc {
					funcName += `"`
				}
			case qDouble:
				qDouble = false
			default:
				qDouble = true
			}

		case ' ':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += ` `
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += ` `
				}
			case expectFunc && readFunc:
				expectFunc = false
				readFunc = false
			}

		/*case "-":
		switch {
		case escaped:
			escaped = false
			if readFunc {
				funcName += `-`
			}
		case qSingle, qDouble:
			if readFunc {
				funcName += `-`
			}
		case !readFunc:

		}*/

		case '>':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `>`
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += `>`
				}
			case i > 0 && line[i-1] == '-':
				expectFunc = true

			}

		case ';', '|':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += string(line[i])
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += string(line[i])
				}
			default:
				expectFunc = true
			}

		case '?':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `?`
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += `?`
				}
			case i > 0 && line[i-1] == ' ':
				expectFunc = true
			default:
				if readFunc {
					funcName += `?`
				}
			}

		case '{':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `{`
				}
			case qSingle, qDouble:
				if readFunc {
					funcName += `{`
				}
			default:
				bracket++
				expectFunc = true
			}

		case '}':
			//loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `}`
				}
			case escaped, qSingle, qDouble:
				if readFunc {
					funcName += `}`
				}
			default:
				bracket--
			}

		case ':':
			switch {
			case escaped:
				escaped = false
			case qSingle, qDouble:
				if readFunc {
					funcName += `:`
				}
			}

		default:
			switch {
			case escaped:
				escaped = false
				fallthrough
			case readFunc:
				funcName += string(line[i])
			case expectFunc:
				readFunc = true
			}
		}
	}

	loc++
	var items []string

	switch {
	case qSingle:
		items = []string{"'"}
	case qDouble:
		items = []string{"\""}
	case expectFunc:
		var s string
		if loc < len(line) {
			s = strings.TrimSpace(string(line[loc:]))
		}
		retPos = len(s)
		items = matchExes(s)

	case bracket > 0:
		items = append([]string{" } "})
	case len(line) > loc && line[loc] == '-':
		items = []string{"> "}
	default:
		items = []string{"{ ", "-> ", "| ", " ? ", "; "}
	}

	suggest = make([][]rune, len(items))
	for i := range items {
		suggest[i] = []rune(items[i])
	}

	return
}
