package readline

func (rl *Instance) getTabCompletion() {
	if rl.TabCompleter == nil {
		return
	}

	rl.tcPrefix, rl.tcSuggestions, rl.tcDefinitions = rl.TabCompleter(rl.line, rl.pos)
	if len(rl.tcSuggestions) == 0 {
		return
	}

	/*if len(rl.tcSuggestions) == 1 && !rl.modeTabCompletion {
		if len(rl.tcSuggestions[0]) == 0 || rl.tcSuggestions[0] == " " || rl.tcSuggestions[0] == "\t" {
			return
		}
		rl.insert([]byte(rl.tcSuggestions[0]))
		return
	}*/

	rl.initTabCompletion()
}

func (rl *Instance) initTabCompletion() {
	if len(rl.tcDefinitions) == 0 {
		rl.initTabGrid()
	} else {
		rl.initTabMap()
	}
}

func (rl *Instance) moveTabCompletionHighlight(x, y int) {
	if len(rl.tcDefinitions) == 0 {
		rl.moveTabGridHighlight(x, y)
	} else {
		rl.moveTabMapHighlight(x, y)
	}
}

func (rl *Instance) writeTabCompletion() {
	if !rl.modeTabCompletion {
		return
	}

	if len(rl.tcDefinitions) == 0 {
		rl.writeTabGrid()
	} else {
		rl.writeTabMap()
	}
}

func (rl *Instance) resetTabCompletion() {
	rl.modeTabCompletion = false
	rl.tcOffset = 0
	rl.tcUsedY = 0
	rl.modeTabFind = false
	rl.tfLine = []rune{}
}
