package datatools

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["struct-keys"] = cmdStructKeys
}

func cmdStructKeys(p *lang.Process) error {
	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, p.Stdin.GetDataType())
	if err != nil {
		return err
	}

	p.Stdout.SetDataType(types.Json)

	aw, err := p.Stdout.WriteArray(types.Json)
	if err != nil {
		return err
	}

	recursive(p.Context, "", v, aw)
	return aw.Close()
}

func recursive(ctx context.Context, path string, v interface{}, aw stdio.ArrayWriter) {
	switch t := v.(type) {
	case []string:
		for i := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			aw.WriteString(path + "/" + strconv.Itoa(i))
		}

	case []int:
		for i := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			aw.WriteString(path + "/" + strconv.Itoa(i))
		}

	case []float64:
		for i := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			aw.WriteString(path + "/" + strconv.Itoa(i))
		}

	case []uint32:
		for i := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			aw.WriteString(path + "/" + strconv.Itoa(i))
		}

	case []bool:
		for i := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			aw.WriteString(path + "/" + strconv.Itoa(i))
		}

	case []interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + strconv.Itoa(key)
			aw.WriteString(newPath)
			recursive(ctx, newPath, t, aw)
		}

	case map[string]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + key
			aw.WriteString(newPath)
			recursive(ctx, newPath, t, aw)
		}

	case map[int]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + strconv.Itoa(key)
			aw.WriteString(newPath)
			recursive(ctx, newPath, t, aw)
		}

	case map[interface{}]interface{}:
		for key := range t {
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + fmt.Sprint(key)
			aw.WriteString(newPath)
			recursive(ctx, newPath, t, aw)
		}

	default:
		//debug.Log(fmt.Sprintf("default: %T", v))
		return
	}
}
