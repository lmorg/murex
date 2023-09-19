package docgen

import "fmt"

const strAlreadyExists = "'%s' already exists in '%s' (%s)"

func unique() bool {
	var failed bool

	m := make(map[string]document)
	for _, d := range Documents {
		dup, ok := m[d.DocumentID]
		if ok {
			warning(d.SourcePath, fmt.Sprintf(strAlreadyExists, d.DocumentID, dup.SourcePath, dup.CategoryID))
			failed = true
			continue
		}

		for _, sym := range d.Synonyms {
			dup, ok := m[sym]
			if ok {
				warning(d.SourcePath, fmt.Sprintf(strAlreadyExists, sym, dup.SourcePath, dup.CategoryID))
				failed = true
				continue
			}
			m[sym] = d
		}

		m[d.DocumentID] = d
	}

	return !failed
}
