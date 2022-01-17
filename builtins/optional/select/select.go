package sqlselect

import (
	"fmt"
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
	lang.DefineMethod("select", cmdSelect, types.Unmarshal, types.Marshal)

	defaults.AppendProfile(`
		config: eval shell safe-commands { -> append select }

		autocomplete set select { [{ 
			"Dynamic": ({ -> select --autocomplete @{$ARGS->@[1..] } }),
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

	parameters, fromFile, err := dissectParameters(p)
	if err != nil {
		return err
	}

	return loadAll(p, fromFile, confFailColMismatch.(bool), confMergeTrailingColumns.(bool), confTableIncHeadings.(bool), confPrintHeadings.(bool), parameters)
}

func dissectParameters(p *lang.Process) (parameters, fromFile string, err error) {
	if p.IsMethod {
		s := p.Parameters.StringAll()
		if rxCheckFrom.MatchString(s) {
			return "", "", fmt.Errorf("SQL contains FROM clause. This should not be included when using `select` as a method")
		}
		return s, "", nil

	} else {
		params := p.Parameters.StringArray()
		i := 0
		for ; i < len(params); i++ {
			if strings.ToLower(params[i]) == "from" {
				goto fromFound
			}
		}
		return "", "", fmt.Errorf("invalid usage. `select` should either be called as a method or include a `FROM file` statement")

	fromFound:
		fromFile = params[i+1]
		if i == 0 {
			params = append([]string{"*"}, params...)
			i++
		}
		if i == len(params)-1 {
			return "", "", fmt.Errorf("invalid usage: `FROM` used but no source file specified")
		}

		return strings.Join(append(params[:i], params[i+2:]...), " "), fromFile, nil
	}
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
