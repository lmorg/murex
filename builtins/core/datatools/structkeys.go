package datatools

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("struct-keys", cmdStructKeys, types.Unmarshal, types.Json)
}

func cmdStructKeys(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags: map[string]string{
			"--depth":     "int",
			"-d":          "--depth",
			"--separator": "str",
			"-s":          "--separator",
		},
	})

	if err != nil {
		return err
	}

	separator := flags["--separator"]
	if separator == "" {
		separator = "/"
	}

	depth := flags["--depth"]
	if depth == "" && len(additional) == 1 {
		depth = additional[0]
	}

	nDeep, _ := strconv.Atoi(depth)
	if nDeep < 1 {
		nDeep = -1
	}

	/*nDeep, _ := p.Parameters.Int(0)
	if nDeep < 1 {
		nDeep = -1 // lets hardcode the max number of iterations for now...
	}*/

	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)

	aw, err := p.Stdout.WriteArray(types.Json)
	if err != nil {
		return err
	}

	err = recursive(p.Context, "", v, separator, aw, nDeep)
	if err != nil {
		return err
	}
	return aw.Close()
}

func recursive(ctx context.Context, path string, v interface{}, separator string, aw stdio.ArrayWriter, iteration int) error {
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

			aw.WriteString(path + separator + strconv.Itoa(i))
		}
		return nil

	case []int:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			aw.WriteString(path + separator + strconv.Itoa(i))
		}
		return nil

	case []float64:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			aw.WriteString(path + separator + strconv.Itoa(i))
		}
		return nil

	case []uint32:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			aw.WriteString(path + separator + strconv.Itoa(i))
		}
		return nil

	case []bool:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			aw.WriteString(path + separator + strconv.Itoa(i))
		}
		return nil

	case []interface{}:
		for i := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + strconv.Itoa(i)
			aw.WriteString(newPath)
			err := recursive(ctx, newPath, t[i], separator, aw, iteration-1)
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
			aw.WriteString(newPath)
			err := recursive(ctx, newPath, t[i], separator, aw, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[string]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + key
			aw.WriteString(newPath)
			err := recursive(ctx, newPath, t[key], separator, aw, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[int]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + strconv.Itoa(key)
			aw.WriteString(newPath)
			err := recursive(ctx, newPath, t[key], separator, aw, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case map[interface{}]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return nil
			default:
			}

			newPath := path + separator + fmt.Sprint(key)
			aw.WriteString(newPath)
			err := recursive(ctx, newPath, t[key], separator, aw, iteration-1)
			if err != nil {
				return err
			}
		}
		return nil

	case string, bool, int, float64, nil, float32, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return nil

	default:
		return fmt.Errorf("found %T but no case exists for handling it", t)
	}
}
