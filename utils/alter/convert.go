package alter

import (
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func StrToInterface(s string) any {
	var new any

	err := json.Unmarshal([]byte(s), &new)
	if err == nil {
		goto newConverted
	}
	new, err = types.ConvertGoType(s, types.Integer)
	if err == nil {
		goto newConverted
	}
	new, err = types.ConvertGoType(s, types.Float)
	if err == nil {
		goto newConverted
	}
	/*new, err = types.ConvertGoType(s, types.Boolean)
	if err == nil {
		goto newConverted
	}*/
	new = s

newConverted:

	return new
}
