package structs

import (
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/lang"
)

const MAX_INT = int(^uint(0) >> 1)

func cmdForEachParallel(p *lang.Process, flags map[string]string, additional []string) error {
	block, varName, err := forEachInitializer(p, additional)
	if err != nil {
		return err
	}

	parallel, err := getFlagValueInt(flags, foreachParallel)
	if err != nil {
		return err
	}

	if parallel < 1 {
		parallel = MAX_INT
	}

	var (
		iteration = int64(-1)
		wg        = new(sync.WaitGroup)
		wait      = make(chan struct{}, parallel)
	)

	err = p.Stdin.ReadArrayWithType(p.Context, func(varValue any, dataType string) {
		i := atomic.AddInt64(&iteration, 1)
		wait <- struct{}{}
		wg.Add(1)
		go func() {
			forEachInnerLoop(p, block, varName, varValue, dataType, int(i))
			wg.Done()
			<-wait
		}()
	})

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}
