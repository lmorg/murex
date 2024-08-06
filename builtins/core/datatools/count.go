package datatools

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

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
		
		autocomplete: set count %[{
			FlagsDesc: {
				"--duplications": "Output a JSON map of items and the number of their occurrences in a list or array"
				"--unique": "Print the number of unique elements in a list or array"
				"--sum": "Read an array, list or map from stdin and output the sum of all the values (ignore non-numeric values)"
				"--sum-strict": "Read an array, list or map from stdin and output the sum of all the values (error on non-numeric values)"
				"--total": "Read an array, list or map from stdin and output the length for that array (default behaviour)"
				"--bytes": "Count the total number of bytes read from stdin",
				"--runes": "Count the total number of unicode characters (runes) read from stdin. Zero width symbols, wide characters and other non-typical graphemes are all each treated as a single rune"
			}
		}]
	`)
}

const (
	argDuplications = "--duplications"
	argUnique       = "--unique"
	argSum          = "--sum"
	argSumStrict    = "--sum-strict"
	argTotal        = "--total"
	argBytes        = "--bytes"
	argRunes        = "--runes"
)

var argsCount = parameters.Arguments{
	AllowAdditional: false,
	Flags: map[string]string{
		argDuplications: types.Boolean,
		"-d":            argDuplications,
		argUnique:       types.Boolean,
		"-u":            argUnique,
		argSum:          types.Boolean,
		"-s":            argSum,
		argSumStrict:    types.Boolean,
		argTotal:        types.Boolean,
		"-t":            argTotal,
		argBytes:        types.Boolean,
		"-b":            argBytes,
		argRunes:        types.Boolean,
		"-r":            argRunes,
	},
}

func cmdCount(p *lang.Process) error {
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

		case argSum, argSumStrict:
			f, err := countSum(p, flags[argSumStrict] == types.TrueString)
			if err != nil {
				p.Stdout.SetDataType(types.Null)
				return err
			}

			p.Stdout.SetDataType(types.Number)
			_, err = p.Stdout.Write([]byte(types.FloatToString(f)))
			return err

		case argBytes:
			p.Stdout.SetDataType(types.Integer)
			b, err := p.Stdin.ReadAll()
			if err != nil {
				return err
			}
			_, err = p.Stdout.Write([]byte(fmt.Sprintf("%d", len(b))))
			return err

		case argRunes:
			p.Stdout.SetDataType(types.Integer)
			b, err := p.Stdin.ReadAll()
			if err != nil {
				return err
			}
			_, err = p.Stdout.Write([]byte(fmt.Sprintf("%d", utf8.RuneCount(b))))
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

func countSum(p *lang.Process, strict bool) (float64, error) {
	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return 0, err
	}

	switch t := v.(type) {
	case nil:
		return 0, nil

	case int:
		return float64(t), nil
	case float64:
		return t, nil

	case []int:
		return sumArrayNum(t)
	case []float64:
		return sumArrayNum(t)

	case []string:
		return sumArrayStr(t, strict)
	case []interface{}:
		return sumArrayStr(t, strict)

	case map[string]int:
		return sumMapNum(t)
	case map[string]float64:
		return sumMapNum(t)
	case map[int]int:
		return sumMapNum(t)
	case map[int]float64:
		return sumMapNum(t)
	case map[float64]int:
		return sumMapNum(t)
	case map[float64]float64:
		return sumMapNum(t)
	case map[any]int:
		return sumMapNum(t)
	case map[any]float64:
		return sumMapNum(t)

	case map[string]string:
		return sumMapStr(t, strict)
	case map[string]any:
		return sumMapStr(t, strict)
	case map[int]string:
		return sumMapStr(t, strict)
	case map[int]any:
		return sumMapStr(t, strict)
	case map[float64]string:
		return sumMapStr(t, strict)
	case map[float64]any:
		return sumMapStr(t, strict)
	case map[any]string:
		return sumMapStr(t, strict)
	case map[any]any:
		return sumMapStr(t, strict)

	case [][]string:
		slice := make([]string, len(t))
		for i := range t {
			slice[i] = strings.Join(t[i], " ")
		}
		return sumArrayStr(slice, strict)

	default:
		return 0, fmt.Errorf("data type (%T) not supported, please report this at https://github.com/lmorg/murex/issues", v)
	}
}

func sumArrayNum[N int | float64](slice []N) (float64, error) {
	var n N

	for _, v := range slice {
		n += v
	}

	return float64(n), nil
}

func sumArrayStr[V any](slice []V, strict bool) (float64, error) {
	var f float64

	for i := range slice {
		v, err := types.ConvertGoType(slice[i], types.Number)
		if err != nil {
			if strict {
				return 0, fmt.Errorf(`cannot convert index %d to %s ("%v"): %s`,
					i, types.Number, slice[i], err.Error())
			} else {
				continue
			}
		}
		f += v.(float64)
	}

	return f, nil
}

func sumMapNum[K comparable, V int | float64](m map[K]V) (float64, error) {
	var n V

	for _, v := range m {
		n += v
	}

	return float64(n), nil
}

func sumMapStr[K comparable, V string | any](m map[K]V, strict bool) (float64, error) {
	var f float64

	for k, v := range m {
		v, err := types.ConvertGoType(v, types.Number)
		if err != nil {
			if strict {
				return 0, fmt.Errorf(`cannot convert index %v to %s ("%v"): %s`,
					k, types.Number, v, err.Error())
			} else {
				continue
			}
		}
		f += v.(float64)
	}

	return f, nil
}
