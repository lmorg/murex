package objectkeys

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
)

func Recursive(ctx context.Context, path string, v any, dataType string, separator string, writeString func(string) error, iteration int) error {
	if dataType == "xml" {
		m, ok := v.(map[string]any)
		if ok && len(m) == 1 {
			for root := range m {
				return recursive(ctx, path, m[root], separator, writeString, iteration)
			}
		}
	}

	return recursive(ctx, path, v, separator, writeString, iteration)
}

func recursive(ctx context.Context, path string, v any, separator string, writeString func(string) error, iteration int) error {
	if iteration == 0 {
		return nil
	}

	switch t := v.(type) {
	case []string:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			err := writeString(path + separator + strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
		return nil

	case []int:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			err := writeString(path + separator + strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
		return nil

	case []float64:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			err := writeString(path + separator + strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
		return nil

	case []uint32:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			err := writeString(path + separator + strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
		return nil

	case []bool:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			err := writeString(path + separator + strconv.Itoa(i))
			if err != nil {
				return err
			}
		}
		return nil

	case []any:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + strconv.Itoa(i)
			err := writeString(newPath)
			if err != nil {
				return err
			}
			err = recursive(ctx, newPath, t[i], separator, writeString, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case [][]string:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + strconv.Itoa(i)
			err := writeString(newPath)
			if err != nil {
				return err
			}
			err = recursive(ctx, newPath, t[i], separator, writeString, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[string]any:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + key
			err := writeString(newPath)
			if err != nil {
				return err
			}
			err = recursive(ctx, newPath, t[key], separator, writeString, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[int]any:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + strconv.Itoa(key)
			err := writeString(newPath)
			if err != nil {
				return err
			}
			err = recursive(ctx, newPath, t[key], separator, writeString, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[any]any:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + fmt.Sprint(key)
			err := writeString(newPath)
			if err != nil {
				return err
			}
			err = recursive(ctx, newPath, t[key], separator, writeString, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case string, bool, int, float64, nil, float32, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return nil

	case lang.MxInterface:
		return recursive(ctx, path, t.GetValue(), separator, writeString, iteration)

	default:
		return fmt.Errorf("found %T but no case exists for handling it", t)
	}
}
