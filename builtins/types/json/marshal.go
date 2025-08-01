package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v any) ([]byte, error) {
	b, err := json.Marshal(v, p.Stdout.IsTTY())
	if err == nil {
		return b, err
	}

	if err.Error() != json.NoData {
		return b, err
	}

	strict, _ := p.Config.Get("proc", "strict-arrays", types.Boolean)
	if strict.(bool) {
		return b, err
	}

	return []byte{'[', ']'}, nil
}

func unmarshal(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)

	// Initially I really liked the idea of JSON files automatically falling
	// back to jsonlines. However on reflection I think it is a bad idea
	// because it then means we need to cover any and all instances where JSON
	// is read but not calling unmarshal - which will be plentiful - else we
	// end up with inconsistent and confusing behavior. But in the current
	// modal all we need to do is the following (see below) so we're not
	// really saving a significant amount of effort.
	//
	//     open ~/jsonlines.json -> cast jsonl -> format json
	/*if err.Error() == "invalid character '{' after top-level value" {
		// ^ this needs a test case so we catch any failures should Go ever
		// change the reported error message
		if jsonl, errJl := unmarshalJsonLines(b); errJl != nil {
			debug.Json(err.Error(), jsonl)
			return jsonl, nil
		}
	}*/

	return
}
