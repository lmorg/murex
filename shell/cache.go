package shell

import (
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/shell/hintsummary"
)

func warmCache() {
	act := autocomplete.AutoCompleteT{
		//Definitions:       make(map[string]string),
		//ErrCallback:       errCallback,
		//DelayedTabContext: dtc,
		//ParsedTokens:      parser.ParsedTokens,
	}

	items := autocomplete.MatchFunction("", &act)
	for i := range items {
		hintsummary.Get(items[i], true)
	}

	//panic("Cache warmed!")
}
