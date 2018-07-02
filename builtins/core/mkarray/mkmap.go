package mkarray

import (
	"errors"
	"sync"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["map"] = mkmap
}

func mkmap(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	blockKey, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	blockValue, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	debug.Log("block key:", string(blockKey))
	debug.Log("block value:", string(blockValue))

	var wg sync.WaitGroup
	var errKeys, errValues error
	outKeys := streams.NewStdin()
	outValues := streams.NewStdin()
	var aKeys, aValues []string

	//go func() {
	//	wg.Add(1)
	_, errKeys = lang.RunBlockExistingConfigSpace(blockKey, nil, outKeys, p.Stderr, p)
	//	wg.Done()
	//}()

	//go func() {
	//	wg.Add(1)
	_, errValues = lang.RunBlockExistingConfigSpace(blockValue, nil, outValues, p.Stderr, p)
	//	wg.Done()
	//}()

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}
	//wg.Wait()

	//go func() {
	//	wg.Add(1)
	errKeys = outKeys.ReadArray(func(b []byte) {
		aKeys = append(aKeys, string(b))
	})
	//	wg.Done()
	//}()

	//go func() {
	//	wg.Add(1)
	errValues = outValues.ReadArray(func(b []byte) {
		aValues = append(aValues, string(b))
	})
	//	wg.Done()
	//}()

	if errKeys != nil {
		return errKeys
	}
	if errValues != nil {
		return errValues
	}
	wg.Wait()

	debug.Json("a keys", aKeys)
	debug.Json("a values", aValues)

	if len(aKeys) > len(aValues) {
		return errors.New("There are more keys than values.")
	}

	if len(aKeys) < len(aValues) {
		return errors.New("There are more values than keys.")
	}

	m := make(map[string]string)
	for i := range aKeys {
		m[aKeys[i]] = aValues[i]
	}
	debug.Json("m", m)

	b, err := json.Marshal(m, p.Stdout.IsTTY())
	if err != nil {
		return err
	}
	_, err = p.Stdout.Write(b)
	return err
}
