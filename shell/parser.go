package shell

import (
	"github.com/chzyer/readline"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type murexCompleterIface struct{}

var (
	murexCompleter    *murexCompleterIface = new(murexCompleterIface)
	rxAllowedVarChars *regexp.Regexp       = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
	rxVars            *regexp.Regexp       = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	rxHistory         *regexp.Regexp       = regexp.MustCompile(`(\^[0-9]+)`)
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

func (mc murexCompleterIface) Do(line []rune, pos int) (suggest [][]rune, retPos int) {
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

	case pt.expectFunc:
		var s string
		if pt.loc < len(line) {
			s = strings.TrimSpace(string(line[pt.loc:]))
		}
		retPos = len(s)
		switch {
		case isLocal(s):
			items = matchLocal(s, true)
			items = append(items, matchDirs(s)...)
		default:
			exes := allExecutables(true)
			items = matchExes(s, &exes, true)
		}

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
		if len(rxVars.FindAllString(string(line), -1))+len(rxHistory.FindAllString(string(line), -1)) > 0 || strings.Contains(string(line), "^!!") {
			os.Stderr.WriteString(utils.NewLineString + "Tap forward again to expand $VARS and ^HISTORY." + utils.NewLineString)
		} else {
			forward = 0
		}
		newPos = pos
		newLine = line

	case key == '{' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && !pt.qDouble {
			//newLine = append(line, '}')
			//newPos = len(newLine) - 1
			newLine = smooshLines(line, pos, '}')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '[' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && !pt.qDouble {
			newLine = smooshLines(line, pos, ']')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '\'' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && pt.qSingle && !pt.qDouble {
			newLine = smooshLines(line, pos, '\'')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == '"' && typed:
		pt := parse(line)
		forward = 0
		if !pt.escaped && !pt.qSingle && pt.qDouble {
			newLine = smooshLines(line, pos, '"')
			newPos = pos
		} else {
			newPos = pos
			newLine = line
		}

	case key == readline.CharBackspace, key == readline.CharDelete:
		newLine = line
		newPos = pos
		forward = 0

		pt := parse(line)
		switch {
		case pt.bracket < 0:
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
		case pt.qSingle:
			newLine, newPos = unsmooshLines(line, pos, '\'')
		case pt.qDouble:
			newLine, newPos = unsmooshLines(line, pos, '"')
		}

	default:
		forward = 0
		newPos = pos
		newLine = line

	}

	if newPos > len(newLine) {
		newPos = len(line) - 1
	} else if newPos < 0 {
		newPos = 0
	}

	return newLine, newPos, true
}

func smooshLines(line []rune, pos int, injectedChar rune) []rune {
	if pos == len(line) {
		return append(line, injectedChar)
	}

	// How the did this happen?
	if pos > len(line) {
		return line
	}

	var i int
	for i = pos; i < len(line); i++ {
		if line[i] == ' ' || line[i] == '\t' {
			continue
		} else {
			break
		}
	}
	// It might seem odd converting this into a slice only to convert back to []rune but Go does some pretty fucked
	// up stuff with slices sometimes due to them literally just being pointers. I found this caused all kinds of
	// annoying little glitches in this routine, as simple as it seems.
	if i >= len(line) || line[i] == '}' || line[i] == '|' || line[i] == ';' {
		before := string(line[:pos])
		after := string(line[pos:])
		newLine := before + string(injectedChar) + after
		return []rune(newLine)
	}

	return line
}

func unsmooshLines(line []rune, pos int, injectedChar rune) ([]rune, int) {
	if pos > len(line) {
		return line, pos
	}

	if pos < 2 {
		if line[0] == injectedChar {
			if len(line) > 1 {
				return line[1:], pos
			} else {
				return []rune{}, 0
			}
		} else {
			return line, pos
		}
	}

	if pos == len(line) {
		if line[pos-1] == injectedChar {
			if len(line) > 1 {
				return line[:pos], pos
			} else {
				return []rune{}, 0
			}
		} else {
			return line, pos
		}
	}

	if line[pos] == injectedChar {
		before := string(line[:pos])
		after := string(line[pos+1:])
		return []rune(before + after), pos
	}

	return line, pos
}

func expandVars(line []rune) []rune {
	s := string(line)
	match := rxVars.FindAllString(s, -1)
	for i := range match {
		s = rxVars.ReplaceAllString(s, proc.GlobalVars.GetString(match[i][1:]))
	}

	match = rxHistory.FindAllString(s, -1)
	for i := range match {
		val, _ := strconv.Atoi(match[i][1:])
		if val > len(History.List) {
			continue
		}
		s = rxHistory.ReplaceAllString(s, History.List[val].Block)
	}

	s = strings.Replace(s, "^!!", History.Last, -1)

	return []rune(s)
}
