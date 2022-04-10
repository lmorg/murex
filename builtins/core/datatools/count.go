package datatools

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func init() {
	lang.DefineMethod("count", cmdCount, types.Unmarshal, types.Json)

	defaults.AppendProfile(`
		alias len=count --total
		
		autocomplete: set count { [{
			"FlagsDesc": {
				"--duplications": "Output a JSON map of items and the number of their occurrences in a list or array",
				"--unique": "Print the number of unique elements in a list or array",
				"--total": "Read an array, list or map from STDIN and output the length for that array (default behaviour)"
			}
		} ]}
	`)
}

const (
	argDuplications = "--duplications"
	argUnique       = "--unique"
	argTotal        = "--total"
)

var argsCount = parameters.Arguments{
	AllowAdditional: false,
	Flags: map[string]string{
		argDuplications: types.Boolean,
		"-d":            argDuplications,
		argUnique:       types.Boolean,
		"-u":            argUnique,
		argTotal:        types.Boolean,
		"-t":            argTotal,
	},
}

func cmdCount(p *lang.Process) error {
	//p.Stdout.SetDataType(types.Json)

	err := p.ErrIfNotAMethod()
	if err != nil {
		return err
	}

	flags, _, err := p.Parameters.ParseFlags(&argsCount)
	if err != nil {
		return err
	}

	if len(flags) == 0 {
		flags[argTotal] = types.TrueString
	}

	for f := range flags {
		switch f {
		case argDuplications:
			v, err := countDuplications(p)
			if err != nil {
				p.Stdout.SetDataType(types.Null)
				return err
			}

			b, err := json.Marshal(v, p.Stdout.IsTTY())
			if err != nil {
				p.Stdout.SetDataType(types.Null)
				return err
			}

			p.Stdout.SetDataType(types.Json)
			_, err = p.Stdout.Write(b)
			return err

		case argUnique:
			v, err := countUnique(p)
			if err != nil {
				p.Stdout.SetDataType(types.Null)
				return err
			}

			p.Stdout.SetDataType(types.Integer)
			_, err = p.Stdout.Write([]byte(strconv.Itoa(v)))
			return err

		case argTotal:
			v, err := countTotal(p)
			if err != nil {
				p.Stdout.SetDataType(types.Null)
				return err
			}

			p.Stdout.SetDataType(types.Integer)
			_, err = p.Stdout.Write([]byte(strconv.Itoa(v)))
			return err
		}
	}

	return nil
}

func countDuplications(p *lang.Process) (map[string]int, error) {
	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return make(map[string]int), err
	}

	return lists.Count(v)
}

func countUnique(p *lang.Process) (int, error) {
	m, err := countDuplications(p)
	if err != nil {
		return 0, err
	}

	return len(m), nil
}

func countTotal(p *lang.Process) (int, error) {
	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return 0, err
	}

	switch v := v.(type) {
	case nil:
		return 0, nil

	case int, float64, string, bool:
		return 1, nil

	case []int:
		return len(v), nil
	case []float64:
		return len(v), nil
	case []string:
		return len(v), nil
	case []bool:
		return len(v), nil
	case []interface{}:
		return len(v), nil

	case map[string]string:
		return len(v), nil
	case map[interface{}]string:
		return len(v), nil

	case map[string]int:
		return len(v), nil
	case map[interface{}]int:
		return len(v), nil

	case map[string]float64:
		return len(v), nil
	case map[interface{}]float64:
		return len(v), nil

	case map[string]bool:
		return len(v), nil
	case map[interface{}]bool:
		return len(v), nil

	case map[string]interface{}:
		return len(v), nil
	case map[interface{}]interface{}:
		return len(v), nil

	case [][]string:
		return len(v), nil

	default:
		return 0, fmt.Errorf("data type (%T) not supported, please report this at https://github.com/lmorg/murex/issues", v)
	}
}
