package sqlselect

import (
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["select"] = cmdSelect

	defaults.AppendProfile(`
		config: eval  shell safe-commands { -> append select }

		autocomplete set select { [{ 
			"Dynamic": ({ -> select --autocomplete ${$ARGS->@[1..] } }),
			"AllowMultiple": true,
			"AnyValue":      true,
			"ExecCmdline":   true
		}] }
	`)
}

func cmdSelect(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	confFailColMismatch := false
	confTableIncHeadings := true

	if flag, _ := p.Parameters.String(0); flag == "--autocomplete" {
		return dynamicAutocomplete(p, confFailColMismatch, confTableIncHeadings)
	}

	return loadAll(p, confFailColMismatch, confTableIncHeadings)
}

func stringToInterface(s []string, max int) []interface{} {
	slice := make([]interface{}, max)
	for i := range slice {
		slice[i] = s[i]
	}

	return slice
}

func stringToInterfacePtr(s *[]string, max int) []interface{} {
	slice := make([]interface{}, max)
	for i := range slice {
		slice[i] = &(*s)[i]
	}

	return slice
}
