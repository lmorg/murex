package json

import (
	"encoding/json"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
)

func marshal(p *proc.Process, v interface{}) ([]byte, error) {
	return utils.JsonMarshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)

	return
}
