package readline

import "fmt"

func HkFnMoveToStartOfLine(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.line.RuneLen() == 0 {
		return
	}
	output := rl.clearHelpersStr()
	rl.line.SetCellPos(0)
	output += rl.echoStr()
	output += moveCursorForwardsStr(1)
	print(output)
}

func HkFnMoveToEndOfLine(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.line.RuneLen() == 0 {
		return
	}
	output := rl.clearHelpersStr()
	rl.line.SetRunePos(rl.line.RuneLen())
	output += rl.echoStr()
	output += moveCursorForwardsStr(1)
	print(output)
}

func HkFnClearAfterCursor(rl *Instance) {
	if rl.line.RuneLen() == 0 {
		return
	}
	output := rl.clearHelpersStr()
	rl.line.Set(rl, rl.line.Runes()[:rl.line.RunePos()])
	output += rl.echoStr()
	output += moveCursorForwardsStr(1)
	print(output)
}

func HkFnClearScreen(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.previewMode != previewModeClosed {
		HkFnPreviewToggle(rl)
	}
	output := seqSetCursorPosTopLeft + seqClearScreen
	output += rl.echoStr()
	output += rl.renderHelpersStr()
	print(output)
}

func HkFnClearLine(rl *Instance) {
	rl.clearPrompt()
	rl.resetHelpers()
}

func HkFnFuzzyFind(rl *Instance) {
	rl.viUndoSkipAppend = true
	if !rl.modeTabCompletion {
		rl.modeAutoFind = true
		rl.getTabCompletion()
	}

	rl.modeTabFind = true
	print(rl.updateTabFindStr([]rune{}))
}

func HkFnSearchHistory(rl *Instance) {
	rl.viUndoSkipAppend = true
	rl.modeAutoFind = true
	rl.tcOffset = 0
	rl.modeTabCompletion = true
	rl.tcDisplayType = TabDisplayMap
	rl.tabMutex.Lock()
	rl.tcSuggestions, rl.tcDescriptions = rl.autocompleteHistory()
	rl.tabMutex.Unlock()
	rl.initTabCompletion()

	rl.modeTabFind = true
	print(rl.updateTabFindStr([]rune{}))
}

func HkFnAutocomplete(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.modeTabCompletion {
		rl.moveTabCompletionHighlight(1, 0)
	} else {
		rl.getTabCompletion()
	}

	print(rl.renderHelpersStr())
}

func HkFnJumpForwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	output := rl.moveCursorByRuneAdjustStr(rl.viJumpE(tokeniseLine))
	print(output)
}

func HkFnJumpBackwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	output := rl.moveCursorByRuneAdjustStr(rl.viJumpB(tokeniseLine))
	print(output)
}

func HkFnCancelAction(rl *Instance) {
	rl.viUndoSkipAppend = true
	var output string
	switch {
	case rl.modeAutoFind:
		output = rl.clearPreviewStr()
		output += rl.resetTabFindStr()
		output += rl.clearHelpersStr()
		rl.resetTabCompletion()
		output += rl.renderHelpersStr()

	case rl.modeTabFind:
		output = rl.resetTabFindStr()

	case rl.modeTabCompletion:
		output = rl.clearPreviewStr()
		output += rl.clearHelpersStr()
		rl.resetTabCompletion()
		output += rl.renderHelpersStr()

	default:
		if rl.line.RunePos() == rl.line.RuneLen() && rl.line.RuneLen() > 0 {
			rl.line.SetRunePos(rl.line.RunePos() - 1)
			output = moveCursorBackwardsStr(1)
		}
		rl.modeViMode = vimKeys
		rl.viIteration = ""
		output += rl.viHintMessageStr()
	}

	print(output)
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

	output := rl.insertStr([]rune(tokens[i-1] + " "))
	print(output)
}

func HkFnPreviewToggle(rl *Instance) {
	if !rl.modeAutoFind && !rl.modeTabCompletion && !rl.modeTabFind &&
		rl.previewMode == previewModeClosed {
		HkFnAutocomplete(rl)
		defer func() { rl.previewMode++ }()
	}
	_HkFnPreviewToggle(rl)
}

func _HkFnPreviewToggle(rl *Instance) {
	rl.viUndoSkipAppend = true
	var output string

	switch rl.previewMode {
	case previewModeClosed:
		output = seqSaveBuffer + seqClearScreen
		rl.previewMode++
		size, _ := rl.getPreviewXY()
		if size != nil {
			output += rl.previewMoveToPromptStr(size)
		}

	case previewModeOpen:
		rl.previewMode = previewModeClosed
		output = seqRestoreBuffer

	case previewModeAutocomplete:
		if rl.modeTabFind {
			print(rl.resetTabFindStr())
		}
		HkFnCancelAction(rl)
	}

	output += rl.echoStr()
	output += rl.renderHelpersStr()
	print(output)
}

func HkFnPreviewLine(rl *Instance) {
	if !rl.modeAutoFind && !rl.modeTabCompletion && !rl.modeTabFind &&
		rl.previewMode == previewModeClosed {
		HkFnAutocomplete(rl)
		defer func() { rl.previewMode++ }()
	}

	rl.previewRef = previewRefLine

	if rl.previewMode == previewModeClosed {
		_HkFnPreviewToggle(rl)
	}
}

func HkFnUndo(rl *Instance) {
	rl.viUndoSkipAppend = true
	if len(rl.viUndoHistory) == 0 {
		return
	}
	output := rl.undoLastStr()
	rl.viUndoSkipAppend = true
	rl.line.SetRunePos(rl.line.RuneLen())
	output += moveCursorForwardsStr(1)
	print(output)
}
