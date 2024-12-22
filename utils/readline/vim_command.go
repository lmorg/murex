package readline

import (
	"regexp"
	"strconv"
)

var (
	rxHistIndex = regexp.MustCompile(`^[0-9]+$`)
	//rxHistRegex   = regexp.MustCompile(`\^m/(.*?[^\\])/`) // Scratchpad: https://play.golang.org/p/Iya2Hx1uxb
	//rxHistPrefix  = regexp.MustCompile(`(\^\^[a-zA-Z]+)`)
	//rxHistTag     = regexp.MustCompile(`(\^#[-_a-zA-Z0-9]+)`)
	//rxHistParam   = regexp.MustCompile(`\^\[([-]?[0-9]+)]`)
	//rxHistReplace = regexp.MustCompile(`\^s/(.*?[^\\])/(.*?[^\\])/`)
)

func (rl *Instance) vimCommandMode(input []rune) string {
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

	return "" //output
}

func (rl *Instance) vimCommandModeReturnStr() string {
	var output string
	s := string(rl.viCommandLine)

	rl.modeViMode = vimInsert
	rl.viCommandLine = nil

	switch {
	case s == "!!":
		line, err := rl.History.GetLine(rl.History.Len() - 1)
		if err != nil {
			rl.hintText = []rune(seqFgRed + err.Error())
		} else {
			output += rl.insertStr([]rune(line))
		}
	case rxHistIndex.MatchString(s):
		i, err := strconv.Atoi(s)
		if err != nil {
			rl.hintText = []rune(seqFgRed + err.Error())
			break
		}
		line, err := rl.History.GetLine(i)
		if err != nil {
			rl.hintText = []rune(seqFgRed + err.Error())
		} else {
			output += rl.insertStr([]rune(line))
		}
	default:
		rl.hintText = []rune(seqFgRed + "Not a valid readline command")
	}

	output += rl.writeHintTextStr()
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
