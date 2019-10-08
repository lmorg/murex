package autocomplete

import (
	"strings"

	"github.com/lmorg/murex/utils/parser"
)

// FormatSuggestions applies some loose formatting rules to auto-completion
// suggestions
func FormatSuggestions(act *AutoCompleteT) {
	formatSuggestionsArray(act.ParsedTokens, act.Items)
	formatSuggestionsMap(act.ParsedTokens, &act.Definitions)
}

func formatSuggestionsArray(pt parser.ParsedTokens, items []string) {
	for i := range items {
		if len(items[i]) == 0 {
			items[i] = " "
			continue
		}

		if !pt.QuoteSingle && !pt.QuoteDouble && pt.QuoteBrace == 0 {
			items[i] = strings.Replace(items[i], `\`, `\\`, -1)
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

func formatSuggestionsMap(pt parser.ParsedTokens, definitions *map[string]string) {
	var (
		newDef = make(map[string]string)
		newKey string
	)

	for key, value := range *definitions {
		if key == "" {
			newDef[" "] = value
			continue
		}

		newKey = key

		if !pt.QuoteSingle && !pt.QuoteDouble && pt.QuoteBrace == 0 {
			newKey = strings.Replace(newKey, `\`, `\\`, -1)
			newKey = strings.Replace(newKey, ` `, `\ `, -1)
			newKey = strings.Replace(newKey, `'`, `\'`, -1)
			newKey = strings.Replace(newKey, `"`, `\"`, -1)
			newKey = strings.Replace(newKey, `(`, `\(`, -1)
			newKey = strings.Replace(newKey, `)`, `\)`, -1)
			newKey = strings.Replace(newKey, `{`, `\{`, -1)
			newKey = strings.Replace(newKey, `}`, `\}`, -1)

			if newKey[len(newKey)-1] != ' ' &&
				newKey[len(newKey)-1] != '=' &&
				newKey[len(newKey)-1] != '/' &&
				len(pt.Variable) == 0 {

				newKey += " "
			}
		}

		newDef[newKey] = value
	}

	*definitions = newDef
}
