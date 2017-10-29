package json

import (
	"encoding/json"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
)

func index(p *proc.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return utils.JsonMarshal(iface, p.Stdout.IsTTY())
	}

	return define.IndexTemplateObject(p, params, &jInterface, marshaller)
}
