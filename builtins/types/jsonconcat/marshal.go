package jsonconcat

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v any) ([]byte, error) {
	var (
		b, jsonl []byte
		err      error
	)

	switch t := v.(type) {
	case []string:
		for i := range t {
			jsonl = append(jsonl, []byte(t[i])...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}
		return jsonl, nil

	case []any:
		for i := range t {
			b, err = json.Marshal(t[i], p.Stdout.IsTTY())
			if err != nil {
				err = fmt.Errorf("unable to marshal %T on line %d: %s", v.([]any)[i], i, err)
				return nil, err
			}

			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}

		return jsonl, nil

	case [][]string:
		for i := range t {
			b, err = json.Marshal(t[i], false)
			if err != nil {
				err = fmt.Errorf("unable to marshal %T on line %d: %s", v.([][]string)[i], i, err)
				return nil, err
			}

			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}

		return jsonl, nil

	case [][]any:
		for i := range t {
			b, err = json.Marshal(t[i], false)
			if err != nil {
				err = fmt.Errorf("unable to marshal %T on line %d: %s", v.([][]any)[i], i, err)
				return nil, err
			}

			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, utils.NewLineByte...)
		}

		return jsonl, nil

	default:
		return nil, fmt.Errorf("cannot marshal data into concatenated JSON. Expecting a slice instead received %T", v)
	}
}
