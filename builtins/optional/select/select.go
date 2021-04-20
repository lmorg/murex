package sqlselect

import (
	"fmt"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["select"] = cmdSelect

	defaults.AppendProfile(`
		config: eval  shell safe-commands { -> append select }

		autocomplete set select { [{ 
			"Dynamic": ({ -> select --autocomplete }),
			"AllowMultiple": true,
			"ExecCmdline": true
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

func parseSQL(sql string, headings []string) string {
	return "" // TODO: parse SQL string
}

func dynamicAutocomplete(p *lang.Process, confFailColMismatch, confTableIncHeadings bool) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	inBytes, _ := p.Stdin.Stats()
	if inBytes > 1024*1024*10 { // 10MB
		return fmt.Errorf("File too large to unmarshal")
	}

	var completions []string

	if confTableIncHeadings {
		v, err := lang.UnmarshalData(p, dt)
		if err != nil {
			return fmt.Errorf("Unable to unmarshal STDIN: %s", err.Error())
		}
		switch v.(type) {
		case [][]string:
			completions = v.([][]string)[0]

		case [][]interface{}:
			completions = make([]string, len(v.([][]interface{})[0]))
			for i := range completions {
				completions[i] = fmt.Sprint(v.([][]interface{})[0][i])
			}

		default:
			return fmt.Errorf("Not a table") // TODO: better error message please
		}
	}

	completions = append(completions,
		"=", ">", ">=", "<", "<=", "<>", "!=", "not", "like",
		"AND", "OR",
		"ORDER BY", "GROUP BY", "SORT",
	)

	b, err := lang.MarshalData(p, types.Json, completions)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
