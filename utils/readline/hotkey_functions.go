package readline

import "fmt"

func HkFnCursorMoveToStartOfLine(rl *Instance) {
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

func HkFnCursorMoveToEndOfLine(rl *Instance) {
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
		HkFnModePreviewToggle(rl)
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

func HkFnModeFuzzyFind(rl *Instance) {
	rl.viUndoSkipAppend = true
	if !rl.modeTabCompletion {
		rl.modeAutoFind = true
		rl.getTabCompletion()
	}

	rl.modeTabFind = true
	print(rl.updateTabFindStr([]rune{}))
}

func HkFnModeSearchHistory(rl *Instance) {
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

func HkFnModeAutocomplete(rl *Instance) {
	rl.viUndoSkipAppend = true
	if rl.modeTabCompletion {
		rl.moveTabCompletionHighlight(1, 0)
	} else {
		rl.getTabCompletion()
	}

	if rl.previewMode == previewModeOpen || rl.previewRef == previewRefLine {
		rl.previewMode = previewModeAutocomplete
	}

	print(rl.renderHelpersStr())
}

func HkFnCursorJumpForwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	output := rl.moveCursorByRuneAdjustStr(rl.viJumpE(tokeniseLine))
	print(output)
}

func HkFnCursorJumpBackwards(rl *Instance) {
	rl.viUndoSkipAppend = true
	output := rl.moveCursorByRuneAdjustStr(rl.viJumpB(tokeniseLine))
	print(output)
}

func HkFnCancelAction(rl *Instance) {
	rl.viUndoSkipAppend = true
	var output string
	switch {
	case rl.modeAutoFind:
		//output += rl.clearPreviewStr()
		output += rl.resetTabFindStr()
		output += rl.clearHelpersStr()
		rl.resetTabCompletion()
		output += rl.renderHelpersStr()

	case rl.modeTabFind:
		output += rl.resetTabFindStr()

	case rl.modeTabCompletion:
		//output = rl.clearPreviewStr()
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

func HkFnRecallWord1(rl *Instance)    { hkFnRecallWord(rl, 1) }
func HkFnRecallWord2(rl *Instance)    { hkFnRecallWord(rl, 2) }
func HkFnRecallWord3(rl *Instance)    { hkFnRecallWord(rl, 3) }
func HkFnRecallWord4(rl *Instance)    { hkFnRecallWord(rl, 4) }
func HkFnRecallWord5(rl *Instance)    { hkFnRecallWord(rl, 5) }
func HkFnRecallWord6(rl *Instance)    { hkFnRecallWord(rl, 6) }
func HkFnRecallWord7(rl *Instance)    { hkFnRecallWord(rl, 7) }
func HkFnRecallWord8(rl *Instance)    { hkFnRecallWord(rl, 8) }
func HkFnRecallWord9(rl *Instance)    { hkFnRecallWord(rl, 9) }
func HkFnRecallWord10(rl *Instance)   { hkFnRecallWord(rl, 10) }
func HkFnRecallWord11(rl *Instance)   { hkFnRecallWord(rl, 11) }
func HkFnRecallWord12(rl *Instance)   { hkFnRecallWord(rl, 12) }
func HkFnRecallWordLast(rl *Instance) { hkFnRecallWordLast(rl) }

const errCannotRecallWord = "Cannot recall word"

func hkFnRecallWord(rl *Instance, i int) {
	tokens := getLastLineTokens(rl)
	if tokens == nil {
		return
	} else if i > len(tokens) {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s %d: previous line contained fewer words", errCannotRecallWord, i))
		return
	}

	printRecalledWord(tokens[i-1], rl)
}

func hkFnRecallWordLast(rl *Instance) {
	tokens := getLastLineTokens(rl)
	if tokens == nil {
		return
	} else if len(tokens) == 0 {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s: previous line contained no words", errCannotRecallWord))
	}

	printRecalledWord(tokens[len(tokens)-1], rl)
}

func getLastLineTokens(rl *Instance) []string {
	line, err := rl.History.GetLine(rl.History.Len() - 1)
	if err != nil {
		rl.ForceHintTextUpdate(fmt.Sprintf("%s: empty history", errCannotRecallWord))
		return nil
	}

	tokens, _, _ := tokeniseSplitSpaces([]rune(line), 0)
	return tokens
}

func printRecalledWord(word string, rl *Instance) {
	word = rTrimWhiteSpace(word)
	output := rl.insertStr([]rune(word + " "))
	print(output)
}

func HkFnModePreviewToggle(rl *Instance) {
	if !rl.modeAutoFind && !rl.modeTabCompletion && !rl.modeTabFind &&
		rl.previewMode == previewModeClosed {

		if rl.modeTabCompletion {
			rl.moveTabCompletionHighlight(1, 0)
		} else {
			rl.getTabCompletion()
		}
		defer func() { rl.previewMode++ }()
	}

	_fnPreviewToggle(rl)
}

func _fnPreviewToggle(rl *Instance) {
	rl.viUndoSkipAppend = true
	var output string

	switch rl.previewMode {
	case previewModeClosed:
		output = curPosSave + seqSaveBuffer + seqClearScreen
		rl.previewMode++
		size, _ := rl.getPreviewXY()
		if size != nil {
			output += rl.previewMoveToPromptStr(size)
		}

	case previewModeOpen:
		print(rl.clearPreviewStr())

	case previewModeAutocomplete:
		print(rl.clearPreviewStr())
		rl.resetHelpers()
	}

	output += rl.echoStr()
	output += rl.renderHelpersStr()
	print(output)
}

func HkFnModePreviewLine(rl *Instance) {
	if rl.PreviewInit != nil {
		// forced rerun of command line preview
		rl.PreviewInit()
		rl.previewCache = nil
	}

	if !rl.modeAutoFind && !rl.modeTabCompletion && !rl.modeTabFind &&
		rl.previewMode == previewModeClosed {
		defer func() { rl.previewMode++ }()
	}

	rl.previewRef = previewRefLine

	if rl.previewMode == previewModeClosed {
		_fnPreviewToggle(rl)
	} else {
		print(rl.renderHelpersStr())
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
