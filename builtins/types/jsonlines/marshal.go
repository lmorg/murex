package jsonlines

import (
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
