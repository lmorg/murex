package xml

import (
	"fmt"
	"strings"

	"github.com/clbanning/mxj/v2"
	"github.com/lmorg/murex/lang"
)

func marshal(p *lang.Process, v interface{}) ([]byte, error) {
	return MarshalTTY(v, p.Stdout.IsTTY())
}

func MarshalTTY(v any, isTTY bool) ([]byte, error) {
	return marshalTTY(v, isTTY, xmlDefaultRoot, xmlDefaultElement)
}

func marshalTTY(v any, isTTY bool, defaultRoot, defaultElement string) ([]byte, error) {
	switch m := v.(type) {
	case map[string]any:
		if m[lang.ELEMENT_META_ROOT] != nil {
			var ok bool
			defaultRoot, ok = m[lang.ELEMENT_META_ROOT].(string)
			if ok {
				delete(m, lang.ELEMENT_META_ROOT)
				break
			}
		}

		if len(m) == 1 {
			for defaultRoot = range m {
			}
			v = m[defaultRoot]
		}

	case []string:
		if len(m) >= 2 {
			key := _elementMeta(m[0], lang.ELEMENT_META_ROOT)
			if key == "" {
				break
			}
			defaultRoot = key
			key = _elementMeta(m[1], lang.ELEMENT_META_ELEMENT)
			if key == "" {
				break
			}
			defaultElement = key
			v = m[2:]
		}

	case []any:
		if len(m) >= 2 {
			key := _elementMeta(m[0], lang.ELEMENT_META_ROOT)
			if key == "" {
				break
			}
			defaultRoot = key
			key = _elementMeta(m[1], lang.ELEMENT_META_ELEMENT)
			if key == "" {
				break
			}
			defaultElement = key
			v = m[2:]
		}
	}

	if isTTY {
		return mxj.AnyXmlIndent(v, "", "    ", defaultRoot, defaultElement)
	}

	return mxj.AnyXml(v, defaultRoot, defaultElement)
}

func _elementMeta(v any, prefix string) string {
	key, ok := v.(string)
	if !ok {
		return ""
	}

	if !strings.HasPrefix(key, prefix) {
		return ""
	}

	return key[len(prefix):]
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

	*ptr = map[string]any(m)
	return nil
}

/*func demeta(v any) error {
	ptr, ok := v.(*any)
	if !ok {
		return fmt.Errorf("expecting a pointer, instead got %T", v)
	}

	switch t := (*ptr).(type) {
	case map[string]any:
		delete(t, lang.ELEMENT_META_ROOT)
		delete(t, lang.ELEMENT_META_ELEMENT)
		*ptr = t
	}

	return nil
}
*/