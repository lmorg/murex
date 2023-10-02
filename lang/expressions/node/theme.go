package node

import (
	"errors"

	"github.com/lmorg/murex/utils/ansi"
)

type ThemeT struct {
	Command      string
	CmdModifier  string
	Parameter    string
	Glob         string
	Number       string
	Bareword     string
	Boolean      string
	Null         string
	Variable     string
	Macro        string
	Escape       string
	QuotedString string
	ArrayItem    string
	ObjectKey    string
	ObjectValue  string
	Operator     string
	Pipe         string
	Comment      string
	Error        string
	Braces       []string

	EndCommand      string
	EndCmdModifier  string
	EndParameter    string
	EndGlob         string
	EndNumber       string
	EndBareword     string
	EndBoolean      string
	EndNull         string
	EndVariable     string
	EndMacro        string
	EndEscape       string
	EndQuotedString string
	EndArrayItem    string
	EndObjectKey    string
	EndObjectValue  string
	EndOperator     string
	EndPipe         string
	EndComment      string
	EndError        string
	EndBraces       []string

	lookup [][]rune
	//previousState [][]rune
	bracePair Symbol
}

func _resetStyle(s string) string {
	if s == "" {
		return "{RESET}"
	}
	return s
}

func (theme *ThemeT) CompileTheme() error {
	if len(theme.EndBraces) == 0 {
		theme.EndBraces = make([]string, len(theme.Braces))
	} else if len(theme.Braces) != len(theme.EndBraces) {
		return errors.New("property 'EndBraces' should be empty or same length as property 'Braces'")
	}

	theme.lookup = make([][]rune, 100) //H_END_BRACE+len(theme.EndBraces))
	theme.bracePair = -1

	noColour := !ansi.IsAllowed()

	theme.lookup[H_COMMAND] = []rune(ansi.ForceExpandConsts(theme.Command, noColour))
	theme.lookup[H_CMD_MODIFIER] = []rune(ansi.ForceExpandConsts(theme.CmdModifier, noColour))
	theme.lookup[H_PARAMETER] = []rune(ansi.ForceExpandConsts(theme.Parameter, noColour))
	theme.lookup[H_GLOB] = []rune(ansi.ForceExpandConsts(theme.Glob, noColour))
	theme.lookup[H_NUMBER] = []rune(ansi.ForceExpandConsts(theme.Number, noColour))
	theme.lookup[H_BAREWORD] = []rune(ansi.ForceExpandConsts(theme.Bareword, noColour))
	theme.lookup[H_BOOLEAN] = []rune(ansi.ForceExpandConsts(theme.Boolean, noColour))
	theme.lookup[H_NULL] = []rune(ansi.ForceExpandConsts(theme.Null, noColour))
	theme.lookup[H_VARIABLE] = []rune(ansi.ForceExpandConsts(theme.Variable, noColour))
	theme.lookup[H_MACRO] = []rune(ansi.ForceExpandConsts(theme.Macro, noColour))
	theme.lookup[H_ESCAPE] = []rune(ansi.ForceExpandConsts(theme.Escape, noColour))
	theme.lookup[H_QUOTED_STRING] = []rune(ansi.ForceExpandConsts(theme.QuotedString, noColour))
	theme.lookup[H_ARRAY_ITEM] = []rune(ansi.ForceExpandConsts(theme.ArrayItem, noColour))
	theme.lookup[H_OBJECT_KEY] = []rune(ansi.ForceExpandConsts(theme.ObjectKey, noColour))
	theme.lookup[H_OBJECT_VALUE] = []rune(ansi.ForceExpandConsts(theme.ObjectValue, noColour))
	theme.lookup[H_OPERATOR] = []rune(ansi.ForceExpandConsts(theme.Operator, noColour))
	theme.lookup[H_PIPE] = []rune(ansi.ForceExpandConsts(theme.Pipe, noColour))
	theme.lookup[H_COMMENT] = []rune(ansi.ForceExpandConsts(theme.Comment, noColour))
	theme.lookup[H_ERROR] = []rune(ansi.ForceExpandConsts(theme.Error, noColour))
	for i := range theme.Braces {
		theme.lookup[_H_BRACE+Symbol(i)] = []rune(ansi.ForceExpandConsts(theme.Braces[i], noColour))
	}

	theme.lookup[H_END_COMMAND] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndCommand), noColour))
	theme.lookup[H_END_CMD_MODIFIER] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndCmdModifier), noColour))
	theme.lookup[H_END_PARAMETER] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndParameter), noColour))
	theme.lookup[H_END_GLOB] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndGlob), noColour))
	theme.lookup[H_END_NUMBER] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndNumber), noColour))
	theme.lookup[H_END_BAREWORD] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndBareword), noColour))
	theme.lookup[H_END_BOOLEAN] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndBoolean), noColour))
	theme.lookup[H_END_NULL] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndNull), noColour))
	theme.lookup[H_END_VARIABLE] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndVariable), noColour))
	theme.lookup[H_END_MACRO] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndMacro), noColour))
	theme.lookup[H_END_ESCAPE] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndEscape), noColour))
	theme.lookup[H_END_QUOTED_STRING] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndQuotedString), noColour))
	theme.lookup[H_END_ARRAY_ITEM] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndArrayItem), noColour))
	theme.lookup[H_END_OBJECT_KEY] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndObjectKey), noColour))
	theme.lookup[H_END_OBJECT_VALUE] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndObjectValue), noColour))
	theme.lookup[H_END_OPERATOR] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndOperator), noColour))
	theme.lookup[H_END_PIPE] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndPipe), noColour))
	theme.lookup[H_END_COMMENT] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndComment), noColour))
	theme.lookup[H_END_ERROR] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndError), noColour))
	for i := range theme.EndBraces {
		theme.lookup[_H_END_BRACE+Symbol(i)] = []rune(ansi.ForceExpandConsts(_resetStyle(theme.EndBraces[i]), noColour))
	}

	return nil
}

func (theme *ThemeT) highlight(keyword Symbol, block ...rune) []rune {
	v := theme.braceAdj(keyword)

	r := theme.lookup[v]
	r = append(r, block...)
	r = append(r, theme.lookup[v+_adjust+1]...)
	//r = append(r, theme.previousStyle()...)

	return r
}

func (theme *ThemeT) braceAdj(keyword Symbol) Symbol {
	switch keyword {
	case H_BRACE_OPEN:
		adj := _H_BRACE + theme.bracePair
		theme.bracePair++
		if theme.bracePair != 0 && len(theme.lookup[_H_BRACE+theme.bracePair]) == 0 {
			theme.bracePair = 0
		}
		return adj

	case H_BRACE_CLOSE:
		theme.bracePair--
		adj := _H_BRACE + theme.bracePair
		if adj < 0 {
			return H_ERROR
		}
		return adj

	default:
		return keyword
	}
}

/*func (theme *ThemeT) addStyle(style []rune) {
	theme.previousState = append(theme.previousState, style)
}

func (theme *ThemeT) restoreStyle() []rune {
	if len(theme.previousState) == 0 {
		return []rune{}
	}
	style := theme.previousState[len(theme.previousState)-1]
	theme.previousState = theme.previousState[:len(theme.previousState)-1]
	return style
}
*/
