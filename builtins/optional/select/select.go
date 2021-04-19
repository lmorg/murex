package sqlselect

import (
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.GoFunctions["select"] = cmdSelect
}

func cmdSelect(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	confFailColMismatch := false
	confTableIncHeadings := true

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
