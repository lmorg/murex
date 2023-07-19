package readline

import "fmt"

func HkFnMoveToStartOfLine(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	rl.clearHelpers()
	rl.line.SetCellPos(0)
	rl.echo()
	moveCursorForwards(1)
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
		rl.clearPreview()
		rl.resetTabFind()
		rl.clearHelpers()
		rl.resetTabCompletion()
		rl.renderHelpers()

	case rl.modeTabFind:
		rl.resetTabFind()

	case rl.modeTabCompletion:
		rl.clearPreview()
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

func HkFnRecallWord1(rl *Instance)  { hkFnRecallWord(rl, 1) }
func HkFnRecallWord2(rl *Instance)  { hkFnRecallWord(rl, 2) }
func HkFnRecallWord3(rl *Instance)  { hkFnRecallWord(rl, 3) }
func HkFnRecallWord4(rl *Instance)  { hkFnRecallWord(rl, 4) }
func HkFnRecallWord5(rl *Instance)  { hkFnRecallWord(rl, 5) }
func HkFnRecallWord6(rl *Instance)  { hkFnRecallWord(rl, 6) }
func HkFnRecallWord7(rl *Instance)  { hkFnRecallWord(rl, 7) }
func HkFnRecallWord8(rl *Instance)  { hkFnRecallWord(rl, 8) }
func HkFnRecallWord9(rl *Instance)  { hkFnRecallWord(rl, 9) }
func HkFnRecallWord10(rl *Instance) { hkFnRecallWord(rl, 10) }
func HkFnRecallWord11(rl *Instance) { hkFnRecallWord(rl, 11) }
func HkFnRecallWord12(rl *Instance) { hkFnRecallWord(rl, 12) }

const errCannotRecallWord = "Cannot recall word"

func hkFnRecallWord(rl *Instance, i int) {
	line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: empty history", errCannotRecallWord, i))
		return
	}

	tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	if i > len(tokens) {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: previous line contained fewer words", errCannotRecallWord, i))
		return
	}

	rl.insert([]rune(tokens[i-1] + " "))
}

func HkFnPreviewToggle(rl *Instance) {
			rl.showPreviews = !rl.showPreviews
		if rl.showPreviews {
			print(seqSaveBuffer)
		} else {
			print(seqRestoreBuffer)
		}

		rl.echo()
		rl.renderHelpers()
}