package readline

import "github.com/mattn/go-runewidth"

type suggestionsT struct {
	rl          *Instance
	suggestions []string
	prefixWidth int
}

func newSuggestionsT(rl *Instance, suggestions []string) *suggestionsT {
	return &suggestionsT{
		rl:          rl,
		suggestions: suggestions,
		prefixWidth: runewidth.StringWidth(rl.tcPrefix),
	}
}

func (s *suggestionsT) Len() int {
	return len(s.suggestions)
}

func (s *suggestionsT) ItemLen(index int) int {
	// fast
	switch {
	case len(s.suggestions[index]) == 0:
		return runewidth.StringWidth(s.rl.tcPrefix)

	case s.suggestions[index][0] == 2:
		if len(s.suggestions[index]) == 1 {
			return 0
		}
		return len(s.suggestions[index][1:])
	default:
		return len(s.suggestions[index]) + s.prefixWidth
	}

	// accurate
	/*switch {
	case len(s.suggestions[index]) == 0:
		return runewidth.StringWidth(s.rl.tcPrefix)

	case s.suggestions[index][0] == 2:
		if len(s.suggestions[index]) == 1 {
			return 0
		}
		return runewidth.StringWidth(s.suggestions[index][1:])
	default:
		return runewidth.StringWidth(s.suggestions[index]) + s.prefixWidth
	}*/
}

func (s *suggestionsT) ItemValue(index int) string {
	switch {
	case len(s.suggestions[index]) == 0:
		return s.rl.tcPrefix

	case s.suggestions[index][0] == 2:
		if len(s.suggestions[index]) == 1 {
			return ""
		}
		return s.suggestions[index][1:]
	default:
		return s.rl.tcPrefix + s.suggestions[index]
	}
}

func (s *suggestionsT) ItemLookupValue(index int) string {
	return s.suggestions[index]
	/*switch {
	case len(s.suggestions[index]) == 0:
		return ""

	case s.suggestions[index][0] == 2:
		if len(s.suggestions[index]) == 1 {
			return ""
		}
		return s.suggestions[index][1:]
	default:
		return s.suggestions[index]
	}*/
}

func (s *suggestionsT) ItemCompletionReturn(index int) (string, string) {
	switch {
	case len(s.suggestions[index]) == 0:
		return s.rl.tcPrefix, ""

	case s.suggestions[index][0] == 2:
		if len(s.suggestions[index]) == 1 {
			return s.rl.tcPrefix, ""
		}
		return "", s.suggestions[index][1:]
	default:
		return s.rl.tcPrefix, s.suggestions[index]
	}
}
