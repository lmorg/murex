package xml

import (
	"fmt"
	"strings"

	"github.com/clbanning/mxj/v2"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/lists"
)

type mapValueT interface {
	~string | ~bool | ~int | ~float64 | any
}

func marshal(p *lang.Process, v any) ([]byte, error) {
	return MarshalTTY(v, p.Stdout.IsTTY())
}

func MarshalTTY(v any, isTTY bool) ([]byte, error) {
	return marshalTTY(v, isTTY, xmlDefaultRoot, xmlDefaultElement)
}

func marshalTTY(v any, isTTY bool, defaultRoot, defaultElement string) ([]byte, error) {
	switch t := v.(type) {
	case map[string]any:
		err := sanitizeKeysMap(t)
		if err != nil {
			return nil, err
		}

		if t[lang.ELEMENT_META_ROOT] != nil {
			var ok bool
			defaultRoot, ok = t[lang.ELEMENT_META_ROOT].(string)
			if ok {
				delete(t, lang.ELEMENT_META_ROOT)
				break
			}
		}

		if len(t) == 1 {
			for defaultRoot = range t {
			}
			v = t[defaultRoot]
		}

	case []string:
		if len(t) >= 2 {
			key := _elementMeta(t[0], lang.ELEMENT_META_ROOT)
			if key == "" {
				break
			}
			defaultRoot = key
			key = _elementMeta(t[1], lang.ELEMENT_META_ELEMENT)
			if key == "" {
				break
			}
			defaultElement = key
			v = t[2:]
		}

	case []any:
		if len(t) >= 2 {
			key := _elementMeta(t[0], lang.ELEMENT_META_ROOT)
			if key == "" {
				break
			}
			defaultRoot = key
			key = _elementMeta(t[1], lang.ELEMENT_META_ELEMENT)
			if key == "" {
				break
			}
			defaultElement = key
			v = t[2:]
		}

	case [][]string:
		var i int
		v = make([]any, len(t)-1)
		err := types.Table2Map(t, func(m map[string]any) error {
			v.([]any)[i] = m
			i++
			return nil
		})
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("there is currently no support for %T types in XML. Please raise this as an issue at %s", v, consts.IssueTrackerURL)
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

	if len(m) == 1 {
		v, ok := m[xmlDefaultRoot]
		if ok {
			*ptr = v
			return nil
		}
	}

	*ptr = map[string]any(m)
	return nil
}

func sanitizeKeysMap[K comparable, V mapValueT](m map[K]V) error {
	for key := range m {
		new, ok, err := sanitizeKeyName(key)
		if err != nil {
			return err
		}
		if ok {
			continue
		}

		m[new] = m[key]
		delete(m, key)
	}

	return nil
}

func sanitizeKeyName[V mapValueT](key V) (V, bool, error) {
	var err error
	switch v := any(key).(type) {
	case string:
		b := []byte(v)
		for i := 0; i < len(b); {

			switch b[i] {

			case '/':
				if i != 0 {
					b[i] = '.'
					i++
					continue
				}

				fallthrough

			case '[', ']':
				b, err = lists.RemoveOrdered(b, i)
				if err != nil {
					return key, false, err
				}
				continue

			}
			i++
		}

		var s mapValueT = string(b)
		return s.(V), s == v, nil

	default:
		return key, false, fmt.Errorf("invalid key '%v' (%T): XML only support string keys", v, v)
	}
}
