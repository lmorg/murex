package xml

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
)

func index(p *lang.Process, params []string) error {
	v, err := UnmarshalFromProcess(p)
	if err != nil {
		return err
	}

	m, ok := v.(map[string]any)
	if !ok || len(m) != 1 {
		return _index(p, params, v, xmlDefaultRoot)
	}

	for root := range m {

		if len(params) == 1 {
			return _index(p, params, m[root], root)
		}

		return _index(p, params, m[root], root)
	}

	panic("unhandled branch") // could should never reach this point
}

func _index(p *lang.Process, params []string, v any, root string) error {
	element := xmlDefaultElement

	marshaller := func(v any) ([]byte, error) {
		switch t := v.(type) {
		case []any:
			if len(params) == 1 {
				element = params[0]
				break
			}

			if len(t) != len(params) {
				// I wouldn't expect this to happen, but just in case
				break
			}

			m := make(map[string]any)
			for i := range t {
				m[params[i]] = t[i]
			}
			v = m

		case map[string]any:
			root = params[0]

		default:
			debug.Logf("unknown type %T", v)
		}

		return marshalTTY(v, p.Stdout.IsTTY(), root, element)
	}

	return lang.IndexTemplateObject(p, params, &v, marshaller)
}
