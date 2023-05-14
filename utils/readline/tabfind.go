package readline

import "strings"

func (rl *Instance) backspaceTabFind() {
	if len(rl.tfLine) > 0 {
		rl.tfLine = rl.tfLine[:len(rl.tfLine)-1]
	}
	rl.updateTabFind([]rune{})
}

func (rl *Instance) updateTabFind(r []rune) {
	rl.tfLine = append(rl.tfLine, r...)

	defer func() {
		rl.clearHelpers()
		rl.initTabCompletion()
		rl.renderHelpers()
	}()

	rl.tabMutex.Lock()
	defer rl.tabMutex.Unlock()

	if len(rl.tfLine) == 0 {
		rl.hintText = rFindSearchPart
		rl.tfSuggestions = append(rl.tcSuggestions, []string{}...)
		return
	}

	var (
		find findT
		err  error
	)

	find, rl.rFindSearch, rl.rFindCancel, err = newFuzzyFind(string(rl.tfLine))
	if err != nil {
		rl.tfSuggestions = []string{err.Error()}
		return
	}

	rl.hintText = append(rl.rFindSearch, rl.tfLine...)

	rl.tfSuggestions = make([]string, 0)
	for i := range rl.tcSuggestions {
		if find.MatchString(strings.TrimSpace(rl.tcSuggestions[i])) {
			rl.tfSuggestions = append(rl.tfSuggestions, rl.tcSuggestions[i])

		} else if rl.tcDisplayType == TabDisplayList && find.MatchString(rl.tcDescriptions[rl.tcSuggestions[i]]) {
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
		rl.hintText = rl.rFindCancel
	}
	rl.modeAutoFind = false

	rl.clearHelpers()
	rl.initTabCompletion()
	rl.renderHelpers()
}
