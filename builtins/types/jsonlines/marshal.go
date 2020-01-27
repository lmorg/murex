package jsonlines

import (
	"fmt"

	"github.com/lmorg/murex/utils"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case []string:
		var jsonl []byte
		for i := range t {
			jsonl = append(jsonl, []byte(t[i])...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}
		return jsonl, nil

	case []interface{}:
		var (
			b, jsonl []byte
			err      error
		)

		for i := range t {
			b, err = json.Marshal(t[i], p.Stdout.IsTTY())
			if err != nil {
				err = fmt.Errorf("Unable to marshal %T on line %d: %s", v.([]interface{})[i], i, err)
			}

			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}

		return jsonl, nil

	default:
		return nil, fmt.Errorf("Cannot marshal data into jsonlines. Expecting a slice instead received %T", v)
	}
}
