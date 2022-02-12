package alter

import (
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
)

func sumMap(v interface{}, new *string) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	var jMap map[string]interface{}
	err = json.Unmarshal([]byte(*new), &jMap)
	if err != nil {
		return ret, fmt.Errorf("cannot unmarshal new map: %s", err)
	}

	switch ret := v.(type) {
	case map[string]int:
		newV := make(map[string]int)
		for k, v := range jMap {
			f, err := types.ConvertGoType(v, types.Integer)
			if err != nil {
				return ret, err
			}
			newV[k] = f.(int)
		}

		lists.SumInt(ret, newV)

	case map[string]float64:
		newV := make(map[string]float64)
		for k, v := range jMap {
			f, err := types.ConvertGoType(v, types.Float)
			if err != nil {
				return ret, err
			}
			newV[k] = f.(float64)
		}

		lists.SumFloat64(ret, newV)

	case map[string]interface{}:
		err := lists.SumInterface(ret, jMap)
		return ret, err

	default:
		return v, fmt.Errorf("path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
