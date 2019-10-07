package autocomplete

import (
	"strings"

	"github.com/lmorg/murex/utils/parser"
)

func FormatSuggestions(act *AutoCompleteT) {
	formatSuggestionsArray(act.ParsedTokens, act.Items)
	formatSuggestionsMap(act.ParsedTokens, act.Definitions)
}

func formatSuggestionsArray(pt parser.ParsedTokens, items []string) {
	for i := range items {
		if len(items[i]) == 0 {
			items[i] = " "
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && pt.QuoteBrace == 0 {
			items[i] = strings.Replace(items[i], ` `, `\ `, -1)
			items[i] = strings.Replace(items[i], `'`, `\'`, -1)
			items[i] = strings.Replace(items[i], `"`, `\"`, -1)
			items[i] = strings.Replace(items[i], `(`, `\(`, -1)
			items[i] = strings.Replace(items[i], `)`, `\)`, -1)
			items[i] = strings.Replace(items[i], `{`, `\{`, -1)
			items[i] = strings.Replace(items[i], `}`, `\}`, -1)

			if items[i][len(items[i])-1] != ' ' &&
				items[i][len(items[i])-1] != '=' &&
				items[i][len(items[i])-1] != '/' &&
				len(pt.Variable) == 0 {
				items[i] += " "
			}
		}

	}
}

func formatSuggestionsMap(pt parser.ParsedTokens, items map[string]string) {
	for k := range items {
		if len(items[k]) == 0 {
			items[k] = " "
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && pt.QuoteBrace == 0 {
			items[k] = strings.Replace(items[k], ` `, `\ `, -1)
			items[k] = strings.Replace(items[k], `'`, `\'`, -1)
			items[k] = strings.Replace(items[k], `"`, `\"`, -1)
			items[k] = strings.Replace(items[k], `(`, `\(`, -1)
			items[k] = strings.Replace(items[k], `)`, `\)`, -1)
			items[k] = strings.Replace(items[k], `{`, `\{`, -1)
			items[k] = strings.Replace(items[k], `}`, `\}`, -1)

			if items[k][len(items[k])-1] != ' ' &&
				items[k][len(items[k])-1] != '=' &&
				items[k][len(items[k])-1] != '/' &&
				len(pt.Variable) == 0 {
				items[k] += " "
			}
		}

	}
}
