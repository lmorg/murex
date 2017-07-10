package table

import (
	"errors"
	"sync"
)

// This needs some debugging...

type table struct {
	rows      [][]string
	mutex     sync.Mutex
	separator byte
	headings  []string
	newRecord bool
	closed    bool
}

func NewTable() table {
	return table{
		newRecord: true,
		separator: ' ',
		rows:      make([][]string, 1),
	}
}

/*func (t table) ReadFrom(r io.Reader) (n int64, err error) {
	var iRead, iWrite int
	for {
		p := make([]byte, 1024)
		iRead, err = r.Read(p)
		if err != nil {
			return
		}

		iWrite, err = t.Write(p[:iRead])
		if err != nil {
			return
		}

		n += int64(iWrite)
	}
}*/

func (t table) Write(b []byte) (i int, err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.closed {
		err = errors.New("Trying to write to a closed table")
		return
	}

	t.newRecord = true
	for i = range b {
		switch b[i] {
		case t.separator:
			t.newRecord = true
		case '\r':
			continue
		case '\n':
			t.rows = append(t.rows, []string{})
			//t.row = &t.rows[len(t.rows)-1]
			t.newRecord = true
		default:
			if t.newRecord {
				t.rows[len(t.rows)-1] = append(t.rows[len(t.rows)-1], string(b[i]))
				t.newRecord = false
			}
			l := len(t.rows) - 1
			t.rows[l][len(t.rows[l])-1] += string(b[i])
		}
	}

	return i + 1, nil
}

func (t table) Close() {
	t.mutex.Lock()
	t.closed = true
	t.mutex.Unlock()
}

func (t table) GetRow() (row []string, eof bool) {
	t.mutex.Lock()

	if len(t.rows) == 0 && t.closed {
		t.mutex.Unlock()
		eof = true
		return
	}

checkAgain:
	i := len(t.rows)
	t.mutex.Unlock()
	if i == 0 {
		t.mutex.Lock()
		goto checkAgain
	}

	t.mutex.Lock()
	row = t.rows[0]
	t.rows = t.rows[1:]
	t.mutex.Unlock()
	return
}
