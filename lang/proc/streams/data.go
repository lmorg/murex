package streams

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"github.com/lmorg/murex/lang/types"
	"io"
	"strconv"
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

// Stream arrays regardless of data type.
// Though currently only 'strings' support streaming, but since this is now a single API it gives an easy place to
// upgrade multiple builtins.
func (read *Stdin) ReadArray(callback func([]byte)) {
	read.mutex.Lock()
	dt := read.dataType
	read.mutex.Unlock()

	switch dt {
	case types.Json:
		b := read.ReadAll()
		j := make([]string, 0)
		err := json.Unmarshal(b, &j)
		if err == nil {
			for i := range j {
				callback([]byte(j[i]))
			}
			return
		}
		fallthrough

	default:
		read.ReadLine(callback)
	}

	return
}

func (read *Stdin) ReadHash(callback func(key, value []byte)) error {
	read.mutex.Lock()
	dt := read.dataType
	read.mutex.Unlock()

	switch dt {
	/*case types.Json:
	b := read.ReadAll()
	j := make(map[string]string, 0)
	err := json.Unmarshal(b, &j)
	if err == nil {
		for i := range j {
			callback([]byte(j[i]))
		}
		return
	}
	fallthrough*/

	case types.Csv:
		r := csv.NewReader(read)
		for {
			_, err := r.Read()
			switch {
			case err == io.EOF:
				return nil
			case err != nil:
				return err
			}

		}

	default:
		scanner := bufio.NewScanner(read)
		var i int
		for scanner.Scan() {
			i++
			callback([]byte(strconv.Itoa(i)), scanner.Bytes())
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
