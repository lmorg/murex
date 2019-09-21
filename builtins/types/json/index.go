package json

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/json"
)

func index(p *lang.Process, params []string) error {
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
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &jInterface, marshaller)
}
