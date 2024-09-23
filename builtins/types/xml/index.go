package xml

import (
	"github.com/lmorg/murex/lang"
)

func index(p *lang.Process, params []string) error {
	v, err := UnmarshalFromProcess(p)
	if err != nil {
		return err
	}

	marshaller := func(v interface{}) ([]byte, error) {
		return MarshalTTY(v, p.Stdout.IsTTY())
	}

	return lang.IndexTemplateObject(p, params, &v, marshaller)
}
