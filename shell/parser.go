package shell

import (
	"regexp"
	"strings"
)

type MurexCompleter struct{}

var (
	murexCompleter    *MurexCompleter = new(MurexCompleter)
	rxAllowedVarChars *regexp.Regexp  = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
)

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
		variable   string
	)

	for i := range line {
		if variable != "" && !rxAllowedVarChars.MatchString(string(line[i])) {
			variable = ""
		}

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

		case '$', '@':
			loc = i
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += string(line[i])
				}
			case qSingle:
				if readFunc {
					funcName += string(line[i])
				}
			default:
				variable = string(line[i])
			}

		case ':':
			switch {
			case escaped:
				escaped = false
				if readFunc {
					funcName += `:`
				}
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
				funcName = string(line[i])
				readFunc = true
			}
		}
	}

	loc++
	var items []string

	switch {
	case variable != "":
		var s string
		if loc < len(line) {
			s = strings.TrimSpace(string(line[loc:]))
		}
		s = variable + s
		retPos = len(s)
		items = getVars(s)

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
		switch {
		case isLocal(s):
			items = matchLocal(s)
		default:
			exes := allExecutables()
			items = matchExes(s, &exes)
		}

	case bracket > 0:
		items = []string{" } "}

	case len(line) > loc && line[loc] == '-':
		items = []string{"> "}

	default:
		items = []string{"{ ", "-> ", "| ", " ? ", "; "}
		switch funcName {
		case "cd":
			var s string
			if loc < len(line) {
				s = strings.TrimSpace(string(line[loc:]))
			}
			retPos = len(s)
			items = append(matchDirs(s))
		default:
			items = append(items, getExeFlags(funcName)...)
		}
	}

	suggest = make([][]rune, len(items))
	for i := range items {
		suggest[i] = []rune(items[i])
	}

	return
}
