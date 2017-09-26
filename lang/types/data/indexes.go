package data

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"regexp"
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

func indexString(p *proc.Process, params []string) error {
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

var (
	rxWhitespace   *regexp.Regexp = regexp.MustCompile(`\s+`)
	rxColumnPrefix *regexp.Regexp = regexp.MustCompile(`^:[0-9]+$`)
	rxRowSuffix    *regexp.Regexp = regexp.MustCompile(`^[0-9]+:$`)
)

func indexGeneric(p *proc.Process, params []string) error {
	var (
		mode     int
		matchStr []string
		matchInt []int
	)

	for i := range params {
		switch {
		case rxRowSuffix.MatchString(params[i]):
			if mode != 0 && mode != 1 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 1
			num, _ := strconv.Atoi(params[i][:len(params[i])-1])
			matchInt = append(matchInt, num)

		case rxColumnPrefix.MatchString(params[i]):
			if mode != 0 && mode != 2 {
				return errors.New("You cannot mix and match matching modes.")
			}
			mode = 2
			num, _ := strconv.Atoi(params[i][1:])
			matchInt = append(matchInt, num)

		default:
			if mode != 0 && mode != 3 {
				return errors.New("You cannot mix and match matching modes.")
			}
			matchStr = append(matchStr, params[i])
			mode = 3

		}
	}

	switch mode {
	case 1:
	// TODO: match rows

	case 2:
		// Match column numbers
		p.Stdin.ReadLine(func(b []byte) {
			recs := rxWhitespace.Split(string(b), -1)
			var line string
			for _, i := range matchInt {
				if i < len(recs) {
					line += "\t" + recs[i]
				} else {
					p.ExitNum = 1
				}
			}
			if len(line) > 1 {
				p.Stdout.Writeln([]byte(line[1:]))
			}
		})

	case 3:
		// Match column names
		var (
			lineNum  int
			headings map[string]int = make(map[string]int)
		)

		p.Stdin.ReadLine(func(b []byte) {
			recs := rxWhitespace.Split(string(b), -1)
			if lineNum == 0 {
				for i := range recs {
					headings[recs[i]] = i + 1
				}

			} else {
				var line string
				for i := range matchStr {
					col := headings[matchStr[i]]
					if col != 0 && col < len(recs) {
						line += "\t" + recs[col-1]
					} else {
						p.ExitNum = 1
					}
				}
				if len(line) > 1 {
					p.Stdout.Writeln([]byte(line[1:]))
				}
			}
			lineNum++
		})

	default:
		return errors.New("You haven't selected any rows / columns.")
	}

	return nil
}
