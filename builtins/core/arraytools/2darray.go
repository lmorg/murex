package arraytools

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["2darray"] = twoDArray
}

type mdarray struct {
	mutex sync.Mutex
	array [][]string
	len   int
}

func newMultiArray(len int) mdarray {
	array := [][]string{make([]string, len)}
	return mdarray{array: array, len: len}
}

func (a *mdarray) Append(index int, count int, value string) {
	a.mutex.Lock()

	if len(a.array) <= count {
		a.array = append(a.array, make([]string, a.len))
	}

	a.array[count][index] = value
	a.mutex.Unlock()
}

func twoDArray(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters. Expecting code blocks to populate array")
	}

	block := make(map[int][]rune)

	for i := 0; i < p.Parameters.Len(); i++ {
		block[i], err = p.Parameters.Block(i)
		if err != nil {
			return err
		}
	}

	var (
		wg       sync.WaitGroup
		array    = newMultiArray(p.Parameters.Len())
		errCount int32
		i        int
	)

	for i = 0; i < p.Parameters.Len(); i++ {
		wg.Add(1)

		index := i
		count := 0

		go func() {
			fork := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
			_, err := fork.Execute(block[index])

			if err != nil {
				fork.Stderr.Write([]byte(fmt.Sprintf("Error executing fork (block %d): %s", index, err.Error())))
			}

			err = fork.Stdout.ReadArray(func(b []byte) {
				count++
				array.Append(index, count, string(b))
			})

			if err != nil {
				p.Stderr.Writeln([]byte(fmt.Sprintf("Error in ReadArray() (block %d): %s: ", index, err.Error())))
				atomic.AddInt32(&errCount, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	if errCount > 0 {
		return fmt.Errorf("%d/%d blocks contained errors. Please read STDERR for details", errCount, i+1)
	}

	b, err := json.Marshal(array.array, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
