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

const (
	IS_RELATIVE = "IsRelative"
	IS_DIR      = "IsDir"
	VALUE       = "Value"
	EXISTS      = "Exists"
)

func Marshal(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		return []byte(t), nil

	case []string:
		s := consts.PathSlash + strings.Join(t, consts.PathSlash)
		return []byte(s), nil

	case map[string]interface{}:
		name, err := types.ConvertGoType(t[VALUE], types.String)
		if err != nil {
			return nil, fmt.Errorf("unable to get '%s' from %v", VALUE, t)
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

	relative, err := types.ConvertGoType(v[0].(map[string]interface{})[IS_RELATIVE], types.Boolean)
	if err != nil {
		return nil, fmt.Errorf("unable to get '%s' from %v", IS_RELATIVE, v[0])
	}

	dir, err := types.ConvertGoType(v[len(v)-1].(map[string]interface{})[IS_DIR], types.Boolean)
	if err != nil {
		return nil, fmt.Errorf("unable to get '%s' from %v", IS_DIR, v[len(v)-1])
	}

	for i := range v {
		switch v[i].(type) {
		case map[string]interface{}:
			name, err := types.ConvertGoType(v[i].(map[string]interface{})[VALUE], types.String)
			if err != nil {
				return nil, fmt.Errorf("unable to get '%s' from %v", VALUE, v[i])
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
	if relative.(bool) {
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

	relative := b[0] != pathSlashByte

	f, err := os.Stat(string(b))
	dir := err == nil && f.IsDir()

	split, err := Split(b)
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
			IS_RELATIVE: relative && i == 0,
			IS_DIR:      (dir && i == len(split)-1) || i < len(split)-1,
			VALUE:       string(split[i]),
		}
	}

	return v, nil
}
