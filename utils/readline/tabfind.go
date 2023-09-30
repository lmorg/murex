package readline

import "strings"

func (rl *Instance) backspaceTabFindStr() string {
	if len(rl.tfLine) > 0 {
		rl.tfLine = rl.tfLine[:len(rl.tfLine)-1]
	}
	return rl.updateTabFindStr([]rune{})
}

func _updateTabFindHelpersStr(rl *Instance) (output string) {
	output = rl._clearHelpers()
	rl.initTabCompletion()
	output += rl._renderHelpers()
	return
}

func (rl *Instance) updateTabFindStr(r []rune) string {
	rl.tfLine = append(rl.tfLine, r...)

	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	if len(rl.tfLine) == 0 {
		rl.hintText = rFindSearchPart
		rl.tfSuggestions = append(rl.tcSuggestions, []string{}...)
		return _updateTabFindHelpersStr(rl)
	}

	var (
		find findT
		err  error
	)

	find, rl.rFindSearch, rl.rFindCancel, err = newFuzzyFind(string(rl.tfLine))
	if err != nil {
		rl.tfSuggestions = []string{err.Error()}
		return _updateTabFindHelpersStr(rl)
	}

	rl.hintText = append(rl.rFindSearch, rl.tfLine...)
	rl.hintText = append(rl.hintText, []rune(seqReset+seqBlink+"_"+seqReset)...)

	rl.tfSuggestions = make([]string, 0)
	for i := range rl.tcSuggestions {
		if find.MatchString(strings.TrimSpace(rl.tcSuggestions[i])) {
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])

		} else if rl.tcDisplayType == TabDisplayList && find.MatchString(rl.tcDescriptions[rl.tcSuggestions[i]]) {
			// this is a list so lets also check the descriptions
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])
		}
	}

	return _updateTabFindHelpersStr(rl)
}

func (rl *Instance) resetTabFindStr() string {
	rl.modeTabFind = false
	rl.tfLine = []rune{}
	if rl.modeAutoFind {
		rl.hintText = []rune{}
	} else {
		rl.hintText = rl.rFindCancel
	}
	rl.modeAutoFind = false

	output := rl._clearHelpers()
	rl.initTabCompletion()
	output += rl._renderHelpers()
	return output
}
