package streams

import (
	"bufio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func (in *Stdin) GetDataType() (dt string) {
	for dt == "" {
		in.mutex.Lock()
		dt = in.dataType
		in.mutex.Unlock()
	}
	return
}

func (in *Stdin) SetDataType(dt string) {
	in.mutex.Lock()
	in.dataType = dt
	in.mutex.Unlock()
	return
}

func (in *Stdin) DefaultDataType(err bool) {
	in.mutex.Lock()
	if in.dataType == "" {
		if err {
			in.dataType = types.Null
		} else {
			in.dataType = types.Generic
		}
	}
	in.mutex.Unlock()
}

// Stream
func (read *Stdin) ReadDataFunc(callback func([]byte)) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(append(scanner.Bytes(), utils.NewLineByte...))
	}

	if err := scanner.Err(); err != nil {
		panic("ReadLine: " + err.Error())
	}

	return
}
