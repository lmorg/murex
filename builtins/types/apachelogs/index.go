package apachelogs

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/json"
)

func index(p *lang.Process, params []string) error {
	jInterface, err := unmarshal(p)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return define.IndexTemplateObject(p, params, &jInterface, marshaller)
}
