package structs

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/mattn/go-runewidth"
)

func cmdForEachParallel(p *lang.Process, flags map[string]string, additional []string) error {
	dataType := p.Stdin.GetDataType()
	if dataType == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dataType)
	}

	var (
		block   []rune
		varName string
	)

	switch len(additional) {
	case 1:
		varName = "!"
		block = []rune(additional[0])

	case 2:
		varName = additional[0]
		block = []rune(additional[1])

	default:
		return errors.New("invalid number of parameters")
	}
	if !types.IsBlockRune(block) {
		return fmt.Errorf("invalid code block: `%s`", runewidth.Truncate(string(block), 70, "…"))
	}

	parallel, err := getFlagValueInt(flags, foreachParallel)
	if err != nil {
		return err
	}

	var (
		max       = int64(parallel)
		queue     atomic.Int64
		iteration = -1
		wg        = new(sync.WaitGroup)
	)

	err = p.Stdin.ReadArrayWithType(p.Context, func(varValue any, dataType string) {
		iteration++
		for {
			if max < 1 || queue.Load() < max {
				break
			}
			time.Sleep(250 * time.Millisecond)
		}
		queue.Add(1)
		wg.Add(1)
		go func() {
			forEachInnerLoop(p, block, varName, varValue, dataType, iteration)
			queue.Add(-1)
			wg.Done()
		}()
	})

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
