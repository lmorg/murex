package data

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/csv"
	"strconv"
)

func indexJson(p *proc.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	var jArray []interface{}
	switch v := jInterface.(type) {
	case []interface{}:
		for _, key := range params {
			i, err := strconv.Atoi(key)
			if err != nil {
				return err
			}
			if i < 0 {
				return errors.New("Cannot have negative keys in array.")
			}
			if i >= len(v) {
				return errors.New("Key '" + key + "' greater than number of items in array.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[i])

			} else {
				switch v[i].(type) {
				case string:
					p.Stdout.Write([]byte(v[i].(string)))
				default:
					b, err := utils.JsonMarshal(v[i], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[string]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := utils.JsonMarshal(v[key], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[interface{}]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := utils.JsonMarshal(v[key], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	default:
		return errors.New("JSON object cannot be indexed.")
	}
}

func indexCsv(p *proc.Process, params []string) error {
	match := make(map[string]int)
	for i := range params {
		match[params[i]] = i + 1
	}

	csvParser, err := csv.NewParser(nil, &proc.GlobalConf)
	if err != nil {
		return err
	}
	records := make([]string, len(params)+1)
	var matched bool

	err = p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
		if match[key] != 0 {
			matched = true
			records[match[key]] = value
		}

		if last && matched {
			p.Stdout.Writeln(csvParser.ArrayToCsv(records[1:]))
			matched = false
			records = make([]string, len(params)+1)
		}
	})

	return err
}

func indexGeneric(p *proc.Process, params []string) error {
	p.Stdout.SetDataType(types.String)

	match := make(map[string]bool)
	for i := range params {
		match[params[i]] = true
	}

	err := p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
		if match[key] {
			p.Stdout.Writeln([]byte(value))
		}
	})

	return err
}
