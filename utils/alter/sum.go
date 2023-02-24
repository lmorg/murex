package alter

import (
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/lists"
)

func sumMap(v interface{}, new *interface{}) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	switch ret := v.(type) {
	case map[string]int:
		newV := make(map[string]int)
		for k, v := range (*new).(map[string]int) {
			f, err := types.ConvertGoType(v, types.Integer)
			if err != nil {
				return ret, err
			}
			newV[k] = f.(int)
		}

		lists.SumInt(ret, newV)

	case map[string]float64:
		newV := make(map[string]float64)
		for k, v := range (*new).(map[string]float64) {
			f, err := types.ConvertGoType(v, types.Float)
			if err != nil {
				return ret, err
			}
			newV[k] = f.(float64)
		}

		lists.SumFloat64(ret, newV)

	case map[string]interface{}:
		err := lists.SumInterface(ret, (*new).(map[string]interface{}))
		return ret, err

	default:
		return v, fmt.Errorf("path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
