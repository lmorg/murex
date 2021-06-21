package lang

import (
	"path/filepath"
	"strings"
)

func init() {
	// add auto globbing to autocomplete
	GoFunctions["@g"] = nil
}

func autoGlob(p *Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if name[len(name)-1] == ':' {
		p.Name.Set(name[:len(name)-1])
	} else {
		p.Name.Set(name)
	}

	var (
		old     = p.Parameters.StringArray()[1:]
		new     []string
		globbed []string
	)

	for i := range old {
		if strings.ContainsAny(old[i], "?*") {
			globbed, err = filepath.Glob(old[i])
			if err != nil {
				return err
			}
			new = append(new, globbed...)
		} else {
			new = append(new, old[i])
		}
	}

	p.Parameters.DefineParsed(new)
	return nil
}
