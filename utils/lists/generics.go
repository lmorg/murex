package lists

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
)

// GenericToString converts []interface to []string
func GenericToString(list interface{}) ([]string, error) {
	switch t := list.(type) {
	case []string:
		return t, nil

	case []any:
		new := make([]string, len(t))
		for i := range t {
			v, err := types.ConvertGoType(t[i], types.String)
			if err != nil {
				return nil, fmt.Errorf("cannot convert element %d: %s", i, err.Error())
			}
			new[i] = v.(string)
		}
		return new, nil

	default:
		return nil, fmt.Errorf("expecting []string or []any, instead got %T", t)
	}
}
