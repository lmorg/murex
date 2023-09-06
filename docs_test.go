package main

import (
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func TestCoreDocs(t *testing.T) {
	count.Tests(t, len(lang.GoFunctions)*3)

	for name := range lang.GoFunctions {

		syn, ok := docs.Synonym[name]
		if !ok || syn == "" {
			t.Errorf("Synonym for `%s` does not exist or is empty", name)
			continue
		}

		sum, ok := docs.Summary[syn]
		if !ok || sum == "" {
			t.Errorf("Summary for `%s` does not exist or is empty", name)
			continue
		}

		doc := docs.Definition(syn)
		if len(doc) == 0 {
			t.Errorf("document for `%s` does not exist or is empty", name)
		}
	}
}
