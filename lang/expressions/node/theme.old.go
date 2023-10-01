package node

/*import (
	"errors"

	"github.com/lmorg/murex/utils/ansi"
)

type Theme struct {
	theme         themeT
	previousState [][]rune

	lookup          [][]rune
	highlightedCode [][]rune

	parent *Theme

	subExpr   int
	bracePair int
}

func _resetStyle(s string) string {
	if s == "" {
		return "{RESET}"
	}
	return s
}

func (sh *Theme) Clone() ThemeT {
	clone := new(Theme)
	clone.lookup = sh.lookup
	clone.parent = sh
	clone.highlightedCode = make([][]rune, 1)
	return clone
}

func (sh *Theme) BeginSubExpr() {
	sh.subExpr++
	sh.highlightedCode = append(sh.highlightedCode, []rune{})
}

func (sh *Theme) UpdateParent() {
	if sh.subExpr == 0 {
		sh.parent.Append(append(sh.highlightedCode[0], sh.restoreStyle()...)...)
		sh.Clear()
		return
	}

	sh.subExpr--
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.highlightedCode[sh.subExpr+1]...)
	sh.highlightedCode = sh.highlightedCode[:sh.subExpr+1]
}

func (sh *Theme) IsHighlighter() bool { return true }

func (sh *Theme) Clear() {
	sh.highlightedCode = make([][]rune, 1)
	sh.subExpr = 0
}

func (sh *Theme) ClearExpr() {
	sh.highlightedCode[sh.subExpr] = []rune{}
}

func (sh *Theme) GetHighlighted() []rune {
	r := append(sh.highlightedCode[sh.subExpr], sh.restoreStyle()...)
	sh.Clear()
	return r
}

func (sh *Theme) compileTheme() error {
	if len(sh.theme.EndBraces) == 0 {
		sh.theme.EndBraces = make([]string, len(sh.theme.Braces))
	} else if len(sh.theme.Braces) != len(sh.theme.EndBraces) {
		return errors.New("property 'EndBraces' should be empty or same length as property 'Braces'")
	}

	sh.lookup = make([][]rune, 100) //H_END_BRACE+len(sh.theme.EndBraces))
	sh.bracePair = -1

	noColour := !ansi.IsAllowed()

	sh.lookup[H_COMMAND] = []rune(ansi.ForceExpandConsts(sh.theme.Command, noColour))
	sh.lookup[H_CMD_MODIFIER] = []rune(ansi.ForceExpandConsts(sh.theme.CmdModifier, noColour))
	sh.lookup[H_PARAMETER] = []rune(ansi.ForceExpandConsts(sh.theme.Parameter, noColour))
	sh.lookup[H_GLOB] = []rune(ansi.ForceExpandConsts(sh.theme.Glob, noColour))
	sh.lookup[H_NUMBER] = []rune(ansi.ForceExpandConsts(sh.theme.Number, noColour))
	sh.lookup[H_BAREWORD] = []rune(ansi.ForceExpandConsts(sh.theme.Bareword, noColour))
	sh.lookup[H_BOOLEAN] = []rune(ansi.ForceExpandConsts(sh.theme.Boolean, noColour))
	sh.lookup[H_NULL] = []rune(ansi.ForceExpandConsts(sh.theme.Null, noColour))
	sh.lookup[H_VARIABLE] = []rune(ansi.ForceExpandConsts(sh.theme.Variable, noColour))
	sh.lookup[H_MACRO] = []rune(ansi.ForceExpandConsts(sh.theme.Macro, noColour))
	sh.lookup[H_ESCAPE] = []rune(ansi.ForceExpandConsts(sh.theme.Escape, noColour))
	sh.lookup[H_QUOTED_STRING] = []rune(ansi.ForceExpandConsts(sh.theme.QuotedString, noColour))
	sh.lookup[H_ARRAY_ITEM] = []rune(ansi.ForceExpandConsts(sh.theme.ArrayItem, noColour))
	sh.lookup[H_OBJECT_KEY] = []rune(ansi.ForceExpandConsts(sh.theme.ObjectKey, noColour))
	sh.lookup[H_OBJECT_VALUE] = []rune(ansi.ForceExpandConsts(sh.theme.ObjectValue, noColour))
	sh.lookup[H_OPERATOR] = []rune(ansi.ForceExpandConsts(sh.theme.Operator, noColour))
	sh.lookup[H_PIPE] = []rune(ansi.ForceExpandConsts(sh.theme.Pipe, noColour))
	sh.lookup[H_COMMENT] = []rune(ansi.ForceExpandConsts(sh.theme.Comment, noColour))
	sh.lookup[H_ERROR] = []rune(ansi.ForceExpandConsts(sh.theme.Error, noColour))
	for i := range sh.theme.Braces {
		sh.lookup[_H_BRACE+i] = []rune(ansi.ForceExpandConsts(sh.theme.Braces[i], noColour))
	}

	sh.lookup[H_END_COMMAND] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndCommand), noColour))
	sh.lookup[H_END_CMD_MODIFIER] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndCmdModifier), noColour))
	sh.lookup[H_END_PARAMETER] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndParameter), noColour))
	sh.lookup[H_END_GLOB] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndGlob), noColour))
	sh.lookup[H_END_NUMBER] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndNumber), noColour))
	sh.lookup[H_END_BAREWORD] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndBareword), noColour))
	sh.lookup[H_END_BOOLEAN] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndBoolean), noColour))
	sh.lookup[H_END_NULL] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndNull), noColour))
	sh.lookup[H_END_VARIABLE] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndVariable), noColour))
	sh.lookup[H_END_MACRO] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndMacro), noColour))
	sh.lookup[H_END_ESCAPE] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndEscape), noColour))
	sh.lookup[H_END_QUOTED_STRING] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndQuotedString), noColour))
	sh.lookup[H_END_ARRAY_ITEM] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndArrayItem), noColour))
	sh.lookup[H_END_OBJECT_KEY] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndObjectKey), noColour))
	sh.lookup[H_END_OBJECT_VALUE] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndObjectValue), noColour))
	sh.lookup[H_END_OPERATOR] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndOperator), noColour))
	sh.lookup[H_END_PIPE] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndPipe), noColour))
	sh.lookup[H_END_COMMENT] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndComment), noColour))
	sh.lookup[H_END_ERROR] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndError), noColour))
	for i := range sh.theme.EndBraces {
		sh.lookup[_H_END_BRACE+i] = []rune(ansi.ForceExpandConsts(_resetStyle(sh.theme.EndBraces[i]), noColour))
	}

	return nil
}

func (sh *Theme) Begin(keyword int, block ...rune) {
	style := sh.lookup[keyword]
	sh.addStyle(style)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], style...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], block...)
}

func (sh *Theme) Append(block ...rune) {
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], block...)
}

func (sh *Theme) End(keyword int, block ...rune) {
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], block...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.lookup[keyword]...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.restoreStyle()...)
}

func (sh *Theme) Highlight(keyword int, block ...rune) {
	v := sh.braceAdj(keyword)

	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.lookup[v]...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], block...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.lookup[v]...)
	sh.highlightedCode[sh.subExpr] = append(sh.highlightedCode[sh.subExpr], sh.previousStyle()...)
}

func (sh *Theme) addStyle(style []rune) {
	sh.previousState = append(sh.previousState, style)
}

func (sh *Theme) restoreStyle() []rune {
	if len(sh.previousState) == 0 {
		return []rune{}
	}
	style := sh.previousState[len(sh.previousState)-1]
	sh.previousState = sh.previousState[:len(sh.previousState)-1]
	return style
}

func (sh *Theme) previousStyle() []rune {
	if len(sh.previousState) == 0 {
		return []rune{}
	}
	return sh.previousState[len(sh.previousState)-1]
}

func (sh *Theme) braceAdj(keyword int) int {
	switch keyword {
	case H_BRACE_OPEN:
		adj := _H_BRACE + sh.bracePair
		sh.bracePair++
		if sh.bracePair != 0 && len(sh.lookup[_H_BRACE+sh.bracePair]) == 0 {
			sh.bracePair = 0
		}
		return adj

	case H_BRACE_CLOSE:
		sh.bracePair--
		adj := _H_BRACE + sh.bracePair
		if adj < 0 {
			return H_ERROR
		}
		return adj

	default:
		return keyword
	}
}
*/
