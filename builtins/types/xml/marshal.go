package xml

import (
	"fmt"

	"github.com/clbanning/mxj/v2"
	"github.com/lmorg/murex/lang"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	return MarshalTTY(v, p.Stdout.IsTTY())
}

/*const (
	_ROOT    = ".."
	_ELEMENT = "."
)*/

func MarshalTTY(v any, isTTY bool) ([]byte, error) {
	defaultRoot := "xml"

	switch m := v.(type) {
	case map[string]any:
		/*if m[_ROOT] != nil {
			var ok bool
			defaultRoot, ok = m[_ROOT].(string)
			if ok {
				delete(m, _ROOT)
				break
			}
		}*/

		if len(m) == 1 {
			for defaultRoot = range m {
			}
			v = m[defaultRoot]
		}
	}

	if isTTY {
		return mxj.AnyXmlIndent(v, "", "    ", defaultRoot)
	}

	return mxj.AnyXml(v, defaultRoot)
}

func UnmarshalFromProcess(p *lang.Process) (v any, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = unmarshaller(b, &v)
	return v, err
}

func unmarshaller(b []byte, v any) error {
	ptr, ok := v.(*any)
	if !ok {
		return fmt.Errorf("cannot unmarshal XML into %T: expecting a pointer", v)
	}

	m, err := mxj.NewMapXml(b, true)
	if err != nil {
		return err
	}

	/*if len(m) == 1 {
		var root string
		for root = range m {
		}

		switch t := m[root].(type) {
		case map[string]any:
			t[_ROOT] = root
			*ptr = t
			return nil

		case []any:
			t = append([]any{root}, t...)
			*ptr = t
			return nil
		}
	}*/

	*ptr = map[string]any(m)
	return nil
}
