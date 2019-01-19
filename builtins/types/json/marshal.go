package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	return json.Marshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *lang.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)

	return
}
