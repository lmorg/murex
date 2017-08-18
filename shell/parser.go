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

type murexCompleterIface struct{}

var (
	murexCompleter    *murexCompleterIface = new(murexCompleterIface)
	rxAllowedVarChars *regexp.Regexp       = regexp.MustCompile(`^[_a-zA-Z0-9]$`)
	rxVars            *regexp.Regexp       = regexp.MustCompile(`(\$[_a-zA-Z0-9]+)`)
	keyPressTimer     time.Time
)

type parseTokens struct {
	loc        int
	escaped    bool
	qSingle    bool
	qDouble    bool
	bracket    int
	expectFunc bool
	__pop      *string
	funcName   string
	parameters []string
	variable   string
}

func parse(line []rune) (pt parseTokens) {
	var readFunc bool
	pt.loc = -1
	pt.expectFunc = true
	pt.__pop = &pt.funcName

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
					*pt.__pop += `#`
				}
				pt.escaped = false
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += `#`
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
					*pt.__pop += `\`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += `\`
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
					*pt.__pop += `'`
				}
			case pt.qDouble:
				if readFunc {
					*pt.__pop += `'`
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
					*pt.__pop += `"`
				}
			case pt.qSingle:
				if readFunc {
					*pt.__pop += `"`
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
					*pt.__pop += ` `
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += ` `
				}
			case pt.expectFunc && readFunc:
				pt.expectFunc = false
				readFunc = false
				pt.parameters = append(pt.parameters, "")
				pt.__pop = &pt.parameters[0]
			default:
				pt.parameters = append(pt.parameters, "")
				pt.__pop = &pt.parameters[len(pt.parameters)-1]
			}

		case '>':
			switch {
			case i > 0 && line[i-1] == '-':
				pt.loc = i
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)
			case pt.expectFunc, readFunc:
				readFunc = true
				*pt.__pop += `>`
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
					*pt.__pop += string(line[i])
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += string(line[i])
				}
			default:
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)

			}

		case '?':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					*pt.__pop += `?`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += `?`
				}
			case i > 0 && line[i-1] == ' ':
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)

			default:
				if readFunc {
					*pt.__pop += `?`
				}
			}

		case '{':
			pt.loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					*pt.__pop += `{`
				}
			case pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += `{`
				}
			default:
				pt.bracket++
				pt.expectFunc = true
				pt.__pop = &pt.funcName
				pt.parameters = make([]string, 0)

			}

		case '}':
			//loc = i
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					*pt.__pop += `}`
				}
			case pt.escaped, pt.qSingle, pt.qDouble:
				if readFunc {
					*pt.__pop += `}`
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
					*pt.__pop += string(line[i])
				}
			case pt.qSingle:
				if readFunc {
					*pt.__pop += string(line[i])
				}
			default:
				pt.variable = string(line[i])
			}

		case ':':
			switch {
			case pt.escaped:
				pt.escaped = false
				if readFunc {
					*pt.__pop += `:`
				}
			case pt.qSingle, pt.qDouble:
				//if readFunc {
				*pt.__pop += `:`
				//}
			case !pt.expectFunc:
				*pt.__pop += `:`
			}

		default:
			switch {
			case pt.escaped:
				pt.escaped = false
				fallthrough
			case readFunc:
				*pt.__pop += string(line[i])
			case pt.expectFunc:
				*pt.__pop = string(line[i])
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
			items = append(items, matchDynamic(s, pt.funcName, pt.parameters)...)
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
		newLine = expandVariables(line)
		newLine = expandHistory(newLine)
		newPos = len(newLine)
		forward = 0

	case forward == 1 && pos == len(line):
		s := string(line)
		if len(rxVars.FindAllString(s, -1)) > 0 ||
			len(rxHistIndex.FindAllString(s, -1)) > 0 ||
			len(rxHistRegex.FindAllString(s, -1)) > 0 ||
			len(rxHistPrefix.FindAllString(s, -1)) > 0 ||
			len(rxHistTag.FindAllString(s, -1)) > 0 ||
			len(rxHistAllPs.FindAllString(s, -1)) > 0 ||
			len(rxHistParam.FindAllString(s, -1)) > 0 ||
			strings.Contains(s, "^!!") {
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
	// It might seem odd converting this into a string only to convert back to []rune but Go does some pretty fucked
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

func expandVariables(line []rune) []rune {
	s := string(line)
	match := rxVars.FindAllString(s, -1)
	for i := range match {
		s = rxVars.ReplaceAllString(s, proc.GlobalVars.GetString(match[i][1:]))
	}

	return []rune(s)
}
