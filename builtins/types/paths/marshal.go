package paths

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/path"
)

func marshalPath(_ *lang.Process, v interface{}) ([]byte, error) {
	return path.Marshal(v)
}

func unmarshalPath(p *lang.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	return path.Unmarshal(b)
}

func marshalPaths(_ *lang.Process, v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		return []byte(t), nil
	case []string:
		s := consts.PathSlash + strings.Join(t, string(pathsSeparator))
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
		return []byte(strings.Join(a, consts.PathSlash+string(pathsSeparator))), nil
	default:
		return nil, fmt.Errorf("%s can only marshal arrays. Instead got %T", types.Paths, t)
	}
}

func unmarshalPaths(p *lang.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	split := bytes.Split(b, pathsSeparator)
	a := make([]string, len(split))
	for i := range split {
		a[i] = string(split[i])
	}
	return a, nil
}
