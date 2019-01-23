package jsonlines

import (
	"bufio"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	switch v.(type) {
	default:
		return nil, fmt.Errorf("Cannot marshal data into jsonlines. Expecting an array of structures and instead received %T", v)

	case []interface{}:
		var (
			b, jsonl []byte
			err      error
		)

		for i, line := range v.([]interface{}) {
			b, err = json.Marshal(line, p.Stdout.IsTTY())
			if err != nil {
				return jsonl, fmt.Errorf("Unable to marshal index %d in array. %s", i, err)
			}
			jsonl = append(jsonl, b...)
			jsonl = append(jsonl, '\n')
		}

		return jsonl, nil
	}
}

func unmarshal(p *lang.Process) (interface{}, error) {
	var (
		jsonl []interface{}
		v     interface{}
		err   error
	)

	scanner := bufio.NewScanner(p.Stdin)
	for scanner.Scan() {
		err = json.Unmarshal(scanner.Bytes(), &v)
		if err != nil {
			return jsonl, fmt.Errorf("Unable to unmarshal index %d in jsonlines: %s", len(jsonl), err)
		}
		jsonl = append(jsonl, v)
	}

	err = scanner.Err()
	return jsonl, err
}
