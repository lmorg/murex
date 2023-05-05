package jsonlines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func iface2Str(a []interface{}) []string {
	s := make([]string, len(a))
	for i := range a {
		s[i] = fmt.Sprint(a[i])
	}
	return s
}

func unmarshal(p *lang.Process) (interface{}, error) {
	var (
		jStruct  []interface{}
		v        interface{}
		jTable   [][]string
		row      []interface{}
		b        []byte
		err      error
		nextEOF  bool
		isStruct bool
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		b = scanner.Bytes()
		if len(bytes.TrimSpace(b)) == 0 {
			continue
		}

		isStruct = isStruct || noSquare(b)

		switch {
		case !isStruct:
			// is a table

			err = json.Unmarshal(b, &row)
			if err == nil {
				jTable = append(jTable, iface2Str(row))
				continue
			}

			isStruct = true
			for i := range jTable {
				jStruct = append(jStruct, jTable[i])
			}

			fallthrough

		case len(jTable) != 0:
			// not a row

			for i := range jTable {
				jStruct = append(jStruct, jTable[i])
			}
			jTable = nil

			fallthrough

		default:
			// is a struct
			err = json.Unmarshal(b, &v)
			switch {
			case err == nil:
				jStruct = append(jStruct, v)
				continue

			case len(jStruct) == 0 && len(b) > 1 && bytes.Contains(b, []byte{'}', '{'}):
				nextEOF = true
				continue

			case noQuote(b) && noSquare(b) && noCurly(b):
				b = append([]byte{'"'}, b...)
				b = append(b, '"')
				err = json.Unmarshal(b, &v)
				if err == nil {
					jStruct = append(jStruct, v)
					continue
				}
				fallthrough

			default:
				return jStruct, fmt.Errorf("unable to unmarshal index %d in jsonlines: %s", len(jStruct), err)
			}
		}

	}

	if err != nil && nextEOF {
		return unmarshalNoCrLF(b)
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling a %s: %s", types.JsonLines, err.Error())
	}

	if isStruct {
		return jStruct, err
	}
	return jTable, err
}

func unmarshalNoCrLF(b []byte) (interface{}, error) {
	var (
		start   int
		v       interface{}
		err     error
		jStruct []interface{}
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
				jStruct = append(jStruct, v)
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
	jStruct = append(jStruct, v)

	return jStruct, nil
}
