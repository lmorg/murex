package jsonlines

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	switch v.(type) {
	case []interface{}:
		var (
			b, jsonl []byte
			err      error
		)

		for i := range v.([]interface{}) {
			b, err = json.Marshal(v.([]interface{})[i], p.Stdout.IsTTY())
			if err != nil {
				err = fmt.Errorf("Unable to marshal %T on line %d: %s", v.([]interface{})[i], i, err)
			}

			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, '\n')
		}

		return jsonl, nil

	default:
		return nil, fmt.Errorf("Cannot marshal data into jsonlines. Expecting a slice instead received %T", v)
	}
}

func unmarshal(p *lang.Process) (interface{}, error) {
	var (
		jsonl   []interface{}
		v       interface{}
		b       []byte
		err     error
		nextEOF bool
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		b = scanner.Bytes()
		err = json.Unmarshal(b, &v)
		if err != nil {
			if len(jsonl) == 0 && len(b) > 1 && bytes.Contains(b, []byte{'}', '{'}) {
				nextEOF = true
			} else {
				return jsonl, fmt.Errorf("Unable to unmarshal index %d in jsonlines: %s", len(jsonl), err)
			}
		}
		jsonl = append(jsonl, v)
	}

	if err != nil && nextEOF {
		return unmarshalNoCrLF(b)
	}

	err = scanner.Err()
	return jsonl, err
}

func unmarshalNoCrLF(b []byte) (interface{}, error) {
	var (
		start   int
		v       interface{}
		err     error
		jsonl   []interface{}
		quoted  bool
		escaped bool
	)

	// don't range because we want to skip first byte to avoid ugly bounds
	// checks within the loop (eg `b[i]-1`).
	for i := 1; i < len(b)-1; i++ {
		switch b[i] {
		case '\\':
			escaped = !escaped
		case '"':
			if escaped {
				escaped = false
			} else {
				quoted = !quoted
			}
		case '{':
			if escaped {
				escaped = false
				continue
			}
			if !quoted && b[i-1] == '}' {
				err = json.Unmarshal(b[start:i], &v)
				if err != nil {
					return nil, err
				}
				jsonl = append(jsonl, v)
				start = i
			}
		default:
			if escaped {
				escaped = false
			}
		}
	}

	// catch remainder
	err = json.Unmarshal(b[start:], &v)
	if err != nil {
		return nil, err
	}
	jsonl = append(jsonl, v)

	return jsonl, nil
}
