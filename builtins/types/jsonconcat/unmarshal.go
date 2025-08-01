package jsonconcat

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func unmarshal(p *lang.Process) (any, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	var jsonc []any
	var i int

	cb := func(j []byte) {
		var v any
		err = json.Unmarshal(j, &v)
		debug.Json(string(j), v)
		jsonc = append(jsonc, v)
		i++
	}

	pErr := parse(b, cb)
	if err != nil {
		return nil, err
	}
	if pErr != nil {
		return nil, pErr
	}

	return jsonc, nil
}
