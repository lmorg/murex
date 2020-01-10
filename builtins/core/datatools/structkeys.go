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

	recursiveMap(p.Context, "", v, aw)
	return aw.Close()
}

func recursiveMap(ctx context.Context, path string, v interface{}, aw stdio.ArrayWriter) {
	switch v.(type) {
	case []interface{}:
		for key := range v.([]interface{}) {
			//debug.Log("v.([]interface{})")
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + strconv.Itoa(key)
			aw.WriteString(newPath)
			recursiveMap(ctx, newPath, v.([]interface{})[key], aw)
		}

	case map[string]interface{}:
		for key := range v.(map[string]interface{}) {
			//debug.Log("v.(map[string]interface{})")
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + key
			aw.WriteString(newPath)
			recursiveMap(ctx, newPath, v.(map[string]interface{})[key], aw)
		}

	case map[int]interface{}:
		for key := range v.(map[int]interface{}) {
			//debug.Log("v.(map[int]interface{})")
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + strconv.Itoa(key)
			aw.WriteString(newPath)
			recursiveMap(ctx, newPath, v.(map[int]interface{})[key], aw)
		}

	case map[interface{}]interface{}:
		for key := range v.(map[interface{}]interface{}) {
			//debug.Log("v.(map[interface{}]interface{})")
			select {
			case <-ctx.Done():
				return
			default:
			}

			newPath := path + "/" + fmt.Sprint(key)
			aw.WriteString(newPath)
			recursiveMap(ctx, newPath, v.(map[interface{}]interface{})[key], aw)
		}

	default:
		//debug.Log(fmt.Sprintf("default: %T", v))
		return
	}
}
