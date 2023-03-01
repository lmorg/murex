package paths

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func marshalPath(_ *lang.Process, v interface{}) ([]byte, error) {
	return marshal(typePath, v, []byte(consts.PathSlash))
}

func marshalPaths(_ *lang.Process, v interface{}) ([]byte, error) {
	return marshal(typePath, v, []byte{':'})
}

func marshal(dataType string, v interface{}, separator []byte) ([]byte, error) {
	switch t := v.(type) {
	case string:
		return []byte(t), nil
	case []string:
		s := strings.Join(t, string(separator))
		return []byte(s), nil
	case []interface{}:
		a := make([]string, len(t))
		for i := range t {
			s, err := types.ConvertGoType(t[i], types.String)
			if err != nil {
				return nil, err
			}
			a[i] = s.(string)
		}
		return []byte(strings.Join(a, string(separator))), nil
	default:
		return nil, fmt.Errorf("%s can only marshal arrays. Instead got %T", dataType, t)
	}
}

func unmarshalPath(p *lang.Process) (interface{}, error) {
	return unmarshal(p, []byte(consts.PathSlash))
}

func unmarshalPaths(p *lang.Process) (interface{}, error) {
	return unmarshal(p, []byte{':'})
}

func unmarshal(p *lang.Process, separator []byte) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	split := bytes.Split(b, separator)
	a := make([]string, len(split))
	for i := range split {
		a[i] = string(split[i])
	}
	return a, nil
}
