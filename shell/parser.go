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
			switch {
			case i > 0 && line[i-1] == '-':
				loc = i
				expectFunc = true
			case expectFunc, readFunc:
				readFunc = true
				funcName += `>`
				fallthrough
			case escaped:
				escaped = false
			default:
				loc = i
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
		items = matchVars(s)

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
			items = matchLocal(s, true)
		default:
			exes := allExecutables(true)
			items = matchExes(s, &exes, true)
		}

	case bracket > 0:
		items = []string{" } "}

	//case len(line) > loc && line[loc] == '-':
	//	items = []string{"> "}

	default:
		items = []string{"{ ", "-> ", "| ", " ? ", "; "}
		var s string
		if loc < len(line) {
			s = strings.TrimSpace(string(line[loc:]))
		}
		retPos = len(s)
		switch funcName {
		case "cd", "mkdir", "rmdir":
			items = matchDirs(s)
		//case "vi", "vim", "cat", "zcat", "text", "open":
		case "man":
			exes := allExecutables(false)
			items = matchExes(s, &exes, false)
		default:
			items = matchFlags(s, funcName)
			switch {
			case !ExesFlags[funcName].NoFiles:
				items = append(items, matchFileAndDirs(s)...)
			case !ExesFlags[funcName].NoDirs:
				items = append(items, matchDirs(s)...)
			}
		}
	}

	maxItems := 30
	if len(items) < maxItems {
		maxItems = len(items)
	}

	suggest = make([][]rune, len(items[:maxItems]))
	for i := range items[:maxItems] {
		suggest[i] = []rune(items[i])
	}
	return
}
