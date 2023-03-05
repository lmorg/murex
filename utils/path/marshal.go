package path

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

/*
	struct:
		[
			{
				"root":  bool,    // only true on top level element if absolute path
				"dir":   bool,    // only false on last element if exists and a directory
				"value": string,  // value of element
			},
			...
		]
*/

func Marshal(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		return []byte(t), nil

	case []string:
		s := consts.PathSlash + strings.Join(t, consts.PathSlash)
		return []byte(s), nil

	case map[string]interface{}:
		name, err := types.ConvertGoType(t["value"], types.String)
		if err != nil {
			return nil, fmt.Errorf("unable to get 'value' from %v", t)
		}
		return []byte(name.(string)), nil

	case []interface{}:
		if len(t) == 0 {
			return nil, nil
		}
		return marshalPathInterface(t)

	default:
		return nil, fmt.Errorf("%s can only marshal arrays. Instead got %T", types.Path, t)
	}
}

func marshalPathInterface(v []interface{}) ([]byte, error) {
	a := make([]string, len(v))

	root, err := types.ConvertGoType(v[0].(map[string]interface{})["root"], types.Boolean)
	if err != nil {
		return nil, fmt.Errorf("unable to get 'root' from %v", v[0])
	}

	dir, err := types.ConvertGoType(v[len(v)-1].(map[string]interface{})["dir"], types.Boolean)
	if err != nil {
		return nil, fmt.Errorf("unable to get 'dir' from %v", v[len(v)-1])
	}

	for i := range v {
		switch v[i].(type) {
		case map[string]interface{}:
			name, err := types.ConvertGoType(v[i].(map[string]interface{})["value"], types.String)
			if err != nil {
				return nil, fmt.Errorf("unable to get 'value' from %v", v[i])
			}
			a[i] = name.(string)

		default:
			name, err := types.ConvertGoType(v[i], types.String)
			if err != nil {
				return nil, err
			}
			a[i] = name.(string)
		}
	}

	b := []byte(strings.Join(a, consts.PathSlash))
	if root.(bool) {
		b = append(pathSlashSlice, b...)
	}

	if dir.(bool) {
		b = append(b, pathSlashByte)
	}

	return b, nil
}

func Unmarshal(b []byte) (interface{}, error) {
	if len(b) == 0 {
		b = []byte{'.'}
	}

	root := b[0] == pathSlashByte

	/*if !root && !bytes.HasPrefix(b, relativePrefix) {
		b = append(relativePrefix, b...)
	}*/

	f, err := os.Stat(string(b))
	dir := err == nil && f.IsDir()

	split, err := SplitPath(b)
	if err != nil {
		return nil, err
	}

	a := make([]string, len(split))
	for i := range split {
		a[i] = string(split[i])
	}

	v := make([]interface{}, len(split))

	for i := range split {
		v[i] = map[string]interface{}{
			"root":  root && i == 0,
			"dir":   (dir && i == len(split)-1) || i < len(split)-1,
			"value": string(split[i]),
		}
	}

	return v, nil
}
