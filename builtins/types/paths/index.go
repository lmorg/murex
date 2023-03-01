package paths

import (
	"github.com/lmorg/murex/lang"
)

func indexPath(p *lang.Process, params []string) error {
	v, err := unmarshalPath(p)
	if err != nil {
		return err
	}

	marshaller := func(v interface{}) ([]byte, error) {
		return marshalPath(nil, v)
	}

	return lang.IndexTemplateObject(p, params, &v, marshaller)
}

func indexPaths(p *lang.Process, params []string) error {
	v, err := unmarshalPaths(p)
	if err != nil {
		return err
	}

	marshaller := func(v interface{}) ([]byte, error) {
		return marshalPaths(nil, v)
	}

	return lang.IndexTemplateObject(p, params, &v, marshaller)
}
