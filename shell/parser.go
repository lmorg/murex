package shell

import (
	"github.com/chzyer/readline"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"os"
	"regexp"
	"strings"
	"time"
)

type MurexCompleter struct{}

var (
	murexCompleter    *MurexCompleter = new(MurexCompleter)
	rxAllowedVarChars *regexp.Regexp  = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
	rxVars            *regexp.Regexp  = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	keyPressTimer     time.Time
)

type parseTokens struct {
	loc        int
	escaped    bool
	qSingle    bool
	qDouble    bool
	bracket    int
	expectFunc bool
	funcName   string
	variable   string
}

func parse(line []rune) (pt parseTokens) {
	var readFunc bool
	pt.loc = -1
	pt.expectFunc = true

	for i := range line {
		if pt.variable != "" && !rxAllowedVarChars.MatchString(string(line[i])) {
			pt.variable = ""
		}

		/*if i > pt.pos-1 {
			//remainder = string(line[i:])
			line = line[:i]
			break
		}*/

		switch line[i] {
		case '#':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `#`
				}
				pt.escaped = false
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `#`
				}
				pt.escaped = false
			default:
				return
			}

		case '\\':
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `\`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `\`
				}
			default:
				pt.escaped = true
			}

		case '\'':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `'`
				}
			case pt.qDouble:
				if readFunc {
					pt.funcName += `'`
				}
			case pt.qSingle:
				pt.qSingle = false
			default:
				pt.qSingle = true
			}

		case '"':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `"`
				}
			case pt.qSingle:
				if readFunc {
					pt.funcName += `"`
				}
			case pt.qDouble:
				pt.qDouble = false
			default:
				pt.qDouble = true
			}

		case ' ':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += ` `
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += ` `
				}
			case pt.expectFunc && readFunc:
				pt.expectFunc = false
				readFunc = false
			}

		case '>':
			switch {
			case i > 0 && line[i-1] == '-':
				pt.loc = i
				pt.expectFunc = true
			case pt.expectFunc, readFunc:
				readFunc = true
				pt.funcName += `>`
				fallthrough
			case pt.escaped:
				pt.escaped = false
			default:
				pt.loc = i
			}

		case ';', '|':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += string(line[i])
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += string(line[i])
				}
			default:
				pt.expectFunc = true
			}

		case '?':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `?`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `?`
				}
			case i > 0 && line[i-1] == ' ':
				pt.expectFunc = true
			default:
				if readFunc {
					pt.funcName += `?`
				}
			}

		case '{':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `{`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `{`
				}
			default:
				pt.bracket++
				pt.expectFunc = true
			}

		case '}':
			//loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `}`
				}
			case pt.escaped, pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `}`
				}
			default:
				pt.bracket--
			}

		case '$', '@':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += string(line[i])
				}
			case pt.qSingle:
				if readFunc {
					pt.funcName += string(line[i])
				}
			default:
				pt.variable = string(line[i])
			}

		case ':':
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					pt.funcName += `:`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					pt.funcName += `:`
				}
			}

		default:
			switch {
			case pt.escaped:
				pt.escaped = false
				fallthrough
			case readFunc:
				pt.funcName += string(line[i])
			case pt.expectFunc:
				pt.funcName = string(line[i])
				readFunc = true
			}
		}
	}
	pt.loc++
	return
}

func (mc MurexCompleter) Do(line []rune, pos int) (suggest [][]rune, retPos int) {
	var items []string
	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt := parse(line)

	switch {
	case pt.variable != "":
		var s string
		if pt.loc < len(line) {
			s = strings.TrimSpace(string(line[pt.loc:]))
		}
		s = pt.variable + s
		retPos = len(s)
		items = matchVars(s)

	//case pt.qSingle:
	//	items = []string{"'"}

	//case pt.qDouble:
	//	items = []string{"\""}

	case pt.expectFunc:
		var s string
		if pt.loc < len(line) {
			s = strings.TrimSpace(string(line[pt.loc:]))
		}
		retPos = len(s)
		switch {
		case isLocal(s):
			items = matchLocal(s, true)
		default:
			exes := allExecutables(true)
			items = matchExes(s, &exes, true)
		}

	//case pt.bracket > 0:
	//	items = []string{" } "}

	//case len(line) > loc && line[loc] == '-':
	//	items = []string{"> "}

	default:
		items = []string{"{ ", "-> ", "| ", " ? ", "; "}
		var s string
		if pt.loc < len(line) {
			s = strings.TrimSpace(string(line[pt.loc:]))
		}
		retPos = len(s)
		switch pt.funcName {
		case "cd", "mkdir", "rmdir":
			items = matchDirs(s)
		case "man", "which", "whereis", "sudo":
			exes := allExecutables(false)
			items = matchExes(s, &exes, false)
		default:
			items = matchFlags(s, pt.funcName)
			switch {
			case !ExesFlags[pt.funcName].NoFiles:
				items = append(items, matchFileAndDirs(s)...)
			case !ExesFlags[pt.funcName].NoDirs:
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
		if len(items[i]) > 1 && strings.Contains(items[i][:len(items[i])-1], " ") {
			items[i] = strings.Replace(items[i][:len(items[i])-1], " ", `\ `, -1) + items[i][len(items[i])-1:]
		}
		suggest[i] = []rune(items[i])
	}

	return
}

func listener(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
	typed := time.Now().After(keyPressTimer)
	keyPressTimer = time.Now().Add(5 * time.Millisecond)

	switch {
	case key == readline.CharEnter:
		return nil, 0, ok

	case forward == 2 && pos == len(line):
		newLine = expandVars(line)
		newPos = len(newLine)
		forward = 0

	case forward == 1 && pos == len(line):
		if len(rxVars.FindAllString(string(line), -1)) > 0 {
			os.Stderr.WriteString(utils.NewLineString + "Tap forward again to expand $VARS." + utils.NewLineString)
		} else {
			forward = 0
		}
		newPos = pos
		newLine = line

	case key == '{' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && !pt.qDouble {
			newLine = append(line, ' ', '}')
			newPos = len(newLine) - 1
		} else {
			newPos = pos
			newLine = line
		}

	case key == '\'' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && pt.qSingle && !pt.qDouble {
			newLine = append(line, '\'')
			newPos = len(newLine) - 1
		} else {
			newPos = pos
			newLine = line
		}

	case key == '"' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && pt.qDouble {
			newLine = append(line, '"')
			newPos = len(newLine) - 1
		} else {
			newPos = pos
			newLine = line
		}

	case key == readline.CharBackspace, key == readline.CharDelete:
		newLine = line
		newPos = pos
		pt := parse(line)
		if pt.bracket < 0 {

			for i := pos; i < len(line); i++ {
				if line[i] == '}' {
					newLine = line[:i]
					if i < len(line) {
						newLine = append(newLine, line[i+1:]...)
					} else {
						newPos = newPos - 1
					}
				}
			}
		}
		forward = 0

	default:
		forward = 0
		newPos = pos
		newLine = line

	}

	return newLine, newPos, true
}

func expandVars(line []rune) []rune {
	s := string(line)
	match := rxVars.FindAllString(s, -1)
	for i := range match {
		s = rxVars.ReplaceAllString(s, proc.GlobalVars.GetString(match[i][1:]))
	}
	return []rune(s)
}
