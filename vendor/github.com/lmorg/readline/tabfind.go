package readline

import "regexp"

func (rl *Instance) backspaceTabFind() {
	if len(rl.tfLine) > 0 {
		rl.tfLine = rl.tfLine[:len(rl.tfLine)-1]
	}
	rl.updateTabFind([]rune{})
}

func (rl *Instance) updateTabFind(r []rune) {
	rl.tfLine = append(rl.tfLine, r...)
	rl.hintText = append([]rune("regexp find: "), rl.tfLine...)

	defer func() {
		rl.clearHelpers()
		rl.initTabCompletion()
		rl.renderHelpers()
	}()

	if len(rl.tfLine) == 0 {
		rl.tfSuggestions = append(rl.tcSuggestions, []string{}...)
		return
	}

	rx, err := regexp.Compile("(?i)" + string(rl.tfLine))
	if err != nil {
		rl.tfSuggestions = []string{err.Error()}
		return
	}

	rl.tfSuggestions = make([]string, 0)
	for i := range rl.tcSuggestions {
		if rx.MatchString(rl.tcSuggestions[i]) {
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])

		} else if rl.tcDisplayType == TabDisplayList && rx.MatchString(rl.tcDescriptions[rl.tcSuggestions[i]]) {
			// this is a list so lets also check the descriptions
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])
		}
	}
}

func (rl *Instance) resetTabFind() {
	rl.modeTabFind = false
	rl.tfLine = []rune{}
	if rl.modeAutoFind {
		rl.hintText = []rune{}
	} else {
		rl.hintText = []rune("Cancelled regexp suggestion find.")
	}

	rl.clearHelpers()
	rl.initTabCompletion()
	rl.renderHelpers()
}
