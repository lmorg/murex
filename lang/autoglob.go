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

	params := p.Parameters.Params[1:]
	p.Parameters.Params = []string{}
	var globbed []string

	for i := range params {
		if strings.ContainsAny(params[i], "?*") {
			globbed, err = filepath.Glob(params[i])
			if err != nil {
				return err
			}
			p.Parameters.Params = append(p.Parameters.Params, globbed...)
		} else {
			p.Parameters.Params = append(p.Parameters.Params, params[i])
		}

	}

	return err
}
