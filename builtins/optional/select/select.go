package sqlselect

import (
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

const ( // Config key names
	sFailColMismatch      = "fail-irregular-columns"
	sTableIncHeadings     = "table-includes-headings"
	sMergeTrailingColumns = "merge-trailing-columns"
	sPrintHeadings        = "print-headings"
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

	config.InitConf.Define("select", sFailColMismatch, config.Properties{
		Description: "When importing a table into sqlite3, fail if there is an irregular number of columns",
		Default:     false,
		DataType:    types.Boolean,
		Global:      false,
	})

	config.InitConf.Define("select", sTableIncHeadings, config.Properties{
		Description: "When importing a table into sqlite3, treat the first row as headings (if `false`, headings are Excel style column references starting at `A`)",
		Default:     true,
		DataType:    types.Boolean,
		Global:      false,
	})

	config.InitConf.Define("select", sMergeTrailingColumns, config.Properties{
		Description: "When importing a table into sqlite3, if `fail-irregular-columns` is set to `false` and there are more columns than headings, then any additional columns are concatenated into the last column (space delimitated). If `merge-trailing-columns` is set to `false` then any trailing columns are ignored",
		Default:     true,
		DataType:    types.Boolean,
		Global:      false,
	})

	config.InitConf.Define("select", sPrintHeadings, config.Properties{
		Description: "Print headings when writing results",
		Default:     true,
		DataType:    types.Boolean,
		Global:      false,
	})
}

func cmdSelect(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	confFailColMismatch, err := p.Config.Get("select", sFailColMismatch, types.Boolean)
	if err != nil {
		return err
	}

	confTableIncHeadings, err := p.Config.Get("select", sTableIncHeadings, types.Boolean)
	if err != nil {
		return err
	}

	confMergeTrailingColumns, err := p.Config.Get("select", sMergeTrailingColumns, types.Boolean)
	if err != nil {
		return err
	}

	confPrintHeadings, err := p.Config.Get("select", sPrintHeadings, types.Boolean)
	if err != nil {
		return err
	}

	if flag, _ := p.Parameters.String(0); flag == "--autocomplete" {
		return dynamicAutocomplete(p, confFailColMismatch.(bool), confTableIncHeadings.(bool))
	}

	return loadAll(p, confFailColMismatch.(bool), confMergeTrailingColumns.(bool), confTableIncHeadings.(bool), confPrintHeadings.(bool))
}

func stringToInterfaceTrim(s []string, max int) []interface{} {
	slice := make([]interface{}, max)

	if max <= len(s) {
		var i int
		for ; i < max; i++ {
			slice[i] = s[i]
		}

		return slice
	}

	var i int
	for ; i < len(s); i++ {
		slice[i] = s[i]
	}

	for ; i < max; i++ {
		slice[i] = ""
	}

	return slice
}

func stringToInterfaceMerge(s []string, max int) []interface{} {
	slice := make([]interface{}, max)

	switch {
	case max == 0:
		// return empty slice

	case max < len(s):
		var i int
		for ; i < max-1; i++ {
			slice[i] = s[i]
		}
		slice[i] = strings.Join(s[i:], " ")

	case max == len(s):
		var i int
		for ; i < max; i++ {
			slice[i] = s[i]
		}

	case max > len(s):
		var i int
		for ; i < len(s); i++ {
			slice[i] = s[i]
		}
		for ; i < max; i++ {
			slice[i] = ""
		}
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

func iToColumnLetter(i int) string {
	var col string

	for i >= 0 {
		mod := i % 26
		col = string([]byte{byte(mod) + 65}) + col
		i = ((i - mod) / 26) - 1
	}

	return col
}
