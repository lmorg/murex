package paths

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/path"
)

func indexPath(p *lang.Process, params []string) error {
	v, err := unmarshalPath(p)
	if err != nil {
		return err
	}

	return lang.IndexTemplateObject(p, params, &v, path.Marshal)
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
