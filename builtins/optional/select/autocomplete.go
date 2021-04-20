package sqlselect

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func dynamicAutocomplete(p *lang.Process, confFailColMismatch, confTableIncHeadings bool) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	inBytes, _ := p.Stdin.Stats()
	if inBytes > 1024*1024*10 { // 10MB
		return fmt.Errorf("File too large to unmarshal")
	}

	var completions []string
	parameters := strings.ToUpper(p.Parameters.StringAll())

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

	if rxQuery.MatchString(parameters) {
		completions = append(completions,
			"=", ">", ">=", "<", "<=", "<>", "!=", "not", "like",
			"AND", "OR",
			"ORDER BY", "GROUP BY",
		)
	} else {
		for i := range completions {
			completions[i] += ","
		}
		completions = append(completions, "*", "WHERE", "ORDER BY", "GROUP BY")
	}

	b, err := lang.MarshalData(p, types.Json, completions)
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
