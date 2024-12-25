package readline

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	rxHistIndex = regexp.MustCompile(`^[0-9]+$`)
	//rxHistRegex   = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
	//rxHistPrefix  = regexp.MustCompile(`(\^\^[a-zA-Z]+)`)
	//rxHistTag     = regexp.MustCompile(`(\^#[-_a-zA-Z0-9]+)`)
	//rxHistParam   = regexp.MustCompile(`\^\[([-]?[0-9]+)]`)
	//rxHistReplace = regexp.MustCompile(`\^s/(.*?[^\\])/(.*?[^\\])/`)
)

func (rl *Instance) vimCommandModeInit() {
	rl.modeViMode = vimCommand
	rl.viUndoSkipAppend = true
	rl.getTabCompletion()
}

func (rl *Instance) vimCommandModeInput(input []rune) string {
	for _, r := range input {
		switch r {
		case '\r', '\n':
			continue
		case '\t':
			rl.viCommandLine = append(rl.viCommandLine, ' ')
		default:
			rl.viCommandLine = append(rl.viCommandLine, r)
		}
	}

	rl.getTabCompletion()
	return "" //output
}

func (rl *Instance) vimCommandModeSuggestions() *TabCompleterReturnT {
	tcr := &TabCompleterReturnT{
		DisplayType: TabDisplayList,
	}

	s := string(rl.viCommandLine)

	switch {
	case s == "":
		fallthrough
	default:
		tcr.Suggestions = []string{
			"Valid commands:",
			"  !!",
			"  n",
			"  m/find/",
			"  s/find/replace/",
		}
		tcr.Descriptions = map[string]string{
			"Valid commands:":   "",
			"  !!":              "Previous command line",
			"  n":               "history item n (where 'n' is an integer)",
			"  m/find/":         "match string (regexp)",
			"  s/find/replace/": "substitute in current line (regexp)",
		}
		tcr.Prefix = " "

	case s == "!!":
		last := rl.History.Len() - 1
		line, err := rl.History.GetLine(last)
		if err != nil {
			return nil
		}
		n := strconv.Itoa(last)
		tcr.Suggestions = []string{n}
		tcr.Descriptions = map[string]string{n: line}

	case rxHistIndex.MatchString(s):
		tcr.Descriptions = make(map[string]string)
		for i := rl.History.Len() - 1; i >= 0; i-- {
			line, err := rl.History.GetLine(i)
			if err != nil {
				return nil // TODO: print err
			}
			n := strconv.Itoa(i)
			if strings.Contains(n, s) {
				tcr.Suggestions = append(tcr.Suggestions, n)
				tcr.Descriptions[n] = line
			}
		}

	case s[0] == 'm':
		if len(s) == 1 {
			return nil
		}
		split := strings.Split(s, string(s[1]))
		rx, err := regexp.Compile(split[1])
		if err != nil {
			return nil // TODO: print err
		}
		tcr.Descriptions = make(map[string]string)
		for i := rl.History.Len() - 1; i >= 0; i-- {
			line, err := rl.History.GetLine(i)
			if err != nil {
				return nil // TODO: print err
			}
			if rx.MatchString(line) {
				n := strconv.Itoa(i)
				tcr.Suggestions = append(tcr.Suggestions, n)
				tcr.Descriptions[n] = line
			}
		}

	case s[0] == 's':
		if len(s) == 1 {
			return nil
		}
		split := strings.Split(s, string(s[1]))
		if len(split) < 3 {
			return nil // TODO: print err
		}
		rx, err := regexp.Compile(split[1])
		if err != nil {
			return nil // TODO: print err
		}
		substitution := rx.ReplaceAllString(rl.line.String(), split[2])
		tcr.Suggestions = []string{substitution}
	}

	return tcr
}

func (rl *Instance) vimCommandModeReturnStr() string {
	var output string

	if len(rl.viCommandLine) == 0 || len(rl.tcSuggestions) == 0 || rl.tcr.Prefix == " " {
		return ""
	}

	cell := (rl.tcMaxX * (rl.tcPosY - 1)) + rl.tcOffset + rl.tcPosX - 1
	v := rl.tcr.Suggestions[cell]

	switch rl.viCommandLine[0] {
	case 's':
		rl.line.Set(rl, []rune(v))
		rl.line.SetRunePos(len(v))

	default:
		_, err := strconv.Atoi(v)
		if err != nil {
			output += rl.insertStr([]rune(v))
		} else {
			output += rl.insertStr([]rune(rl.tcr.Descriptions[v]))
		}
	}

	rl.modeViMode = vimInsert
	rl.viCommandLine = nil

	rl.resetHelpers()
	output += rl.clearHelpersStr()
	//rl.resetTabCompletion()
	output += rl.renderHelpersStr()
	return output
}

func (rl *Instance) vimCommandModeHintText() []rune {
	r := append([]rune("VIM command mode: "), rl.viCommandLine...)
	r = append(r, []rune(seqBlink)...)
	r = append(r, '_')
	return r
}

func (rl *Instance) vimCommandModeBackspaceStr() string {
	if len(rl.viCommandLine) == 0 {
		return "\007" // bell
	}

	rl.viCommandLine = rl.viCommandLine[:len(rl.viCommandLine)-1]
	rl.hintText = rl.vimCommandModeHintText()
	return rl.renderHelpersStr()
}

func (rl *Instance) vimCommandModeCancel() {
	rl.modeViMode = vimInsert
	rl.viCommandLine = nil
}
