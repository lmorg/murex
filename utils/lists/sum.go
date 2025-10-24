package lists

import (
	"fmt"

	"github.com/lmorg/murex/lang/types"
)

func SumInt(dst, src map[string]int) {
	for key, i := range src {
		dst[key] += i
	}
}

func SumFloat64(dst, src map[string]float64) {
	for key, i := range src {
		dst[key] += i
	}
}

func SumInterface(dst, src map[string]any) error {
	for key, v := range src {
		f1, err := types.ConvertGoType(dst[key], types.Float)
		if err != nil {
			return fmt.Errorf("cannot convert '%v' in destination to `number` (float64)", dst[key])
		}

		f2, err := types.ConvertGoType(v, types.Float)
		if err != nil {
			return fmt.Errorf("cannot convert '%v' in source to `number` (float64)", v)
		}

		dst[key] = f1.(float64) + f2.(float64)
	}

	return nil
}
