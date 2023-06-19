package readline

func HkFnMoveToStartOfLine(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.SetCellPos(1)
	rl.echo()
}

func HkFnMoveToEndOfLine(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.SetRunePos(rl.line.RuneLen())
	rl.echo()
	moveCursorForwards(1)
}

func HkFnClearAfterCursor(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.Set(rl.line.Runes()[:rl.line.RunePos()])
	rl.echo()
	moveCursorForwards(1)
}

func HkFnClearScreen(rl *Instance) {
	print(seqSetCursorPosTopLeft + seqClearScreen)
	rl.echo()
	rl.renderHelpers()
}

func HkFnClearLine(rl *Instance) {
	rl.clearPrompt()
	rl.resetHelpers()
}

func HkFnFuzzyFind(rl *Instance) {
	if !rl.modeTabCompletion {
		rl.modeAutoFind = true
		rl.getTabCompletion()
	}

	rl.modeTabFind = true
	rl.updateTabFind([]rune{})
	rl.viUndoSkipAppend = true
}

func HkFnSearchHistory(rl *Instance) {
	rl.modeAutoFind = true
	rl.tcOffset = 0
	rl.modeTabCompletion = true
	rl.tcDisplayType = TabDisplayMap
	rl.tabMutex.Lock()
	rl.tcSuggestions, rl.tcDescriptions = rl.autocompleteHistory()
	rl.tabMutex.Unlock()
	rl.initTabCompletion()

	rl.modeTabFind = true
	rl.updateTabFind([]rune{})
	rl.viUndoSkipAppend = true
}

func HkFnAutocomplete(rl *Instance) {
	if rl.modeTabCompletion {
		rl.moveTabCompletionHighlight(1, 0)
	} else {
		rl.getTabCompletion()
	}

	rl.renderHelpers()
	rl.viUndoSkipAppend = true
}

func HkFnJumpForwards(rl *Instance) {
	rl.moveCursorByRuneAdjust(rl.viJumpE(tokeniseLine))
}

func HkFnJumpBackwards(rl *Instance) {
	rl.moveCursorByRuneAdjust(rl.viJumpB(tokeniseLine))
}

func HkFnCancelAction(rl *Instance) {
	switch {
	case rl.modeAutoFind:
		rl.resetTabFind()
		rl.clearHelpers()
		rl.resetTabCompletion()
		rl.renderHelpers()

	case rl.modeTabFind:
		rl.resetTabFind()

	case rl.modeTabCompletion:
		rl.clearHelpers()
		rl.resetTabCompletion()
		rl.renderHelpers()

	default:
		if rl.line.RunePos() == rl.line.RuneLen() && rl.line.RuneLen() > 0 {
			rl.line.SetRunePos(rl.line.RunePos() - 1)
			moveCursorBackwards(1)
		}
		rl.modeViMode = vimKeys
		rl.viIteration = ""
		rl.viHintMessage()
	}
	rl.viUndoSkipAppend = true
}
