package arraytools

import (
	"errors"
	"sync"

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

	var wg sync.WaitGroup
	array := newMultiArray(p.Parameters.Len())

	for i := 0; i < p.Parameters.Len(); i++ {
		wg.Add(1)

		index := i
		count := 0

		go func() {
			//_, err := lang.RunBlockExistingConfigSpace(block[index], nil, out, p.Stderr, p)
			fork := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
			_, err := fork.Execute(block[index])

			if err != nil {
				fork.Stderr.Write([]byte(err.Error()))
			}

			/*err =*/ fork.Stdout.ReadArray(func(b []byte) {
				count++
				array.Append(index, count, string(b))
			})
                        // this will likely need to be done via a channel and selects
			//if err != nil {
			//	return err
			//}

			wg.Done()
		}()
	}

	wg.Wait()

	b, err := json.Marshal(array.array, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
