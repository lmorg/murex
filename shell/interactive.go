package shell

import (
	"github.com/lmorg/murex/lang/proc"
	"sort"
	"strings"
)

var murexCompleter *MurexCompleter = new(MurexCompleter)

type MurexCompleter struct{}

func (fz MurexCompleter) Do(line []rune, pos int) (suggest [][]rune, retPos int) {
	var (
		loc     int = -1
		escaped bool
		qSingle bool
		qDouble bool
		//bracket int
		expectFunc bool = true
		readFunc   bool
	)

	for i := range line {
		switch line[i] {
		case '#':
			loc = i
			switch {
			case escaped, qSingle, qDouble:
			default:
				return
			}

		case '\'':
			loc = i
			switch {
			case escaped, qDouble:
			case qSingle:
				qSingle = false
			default:
				qSingle = true
			}

		case '"':
			loc = i
			switch {
			case escaped, qSingle:
			case qDouble:
				qDouble = false
			default:
				qDouble = true
			}

		case ';', '|':
			loc = i
			switch {
			case escaped, qSingle, qDouble:
			default:
				expectFunc = true
			}

		case ' ':
			loc = i
			switch {
			case escaped, qSingle, qDouble:
			case expectFunc && readFunc:
				expectFunc = false
				readFunc = false
			}

		default:
			switch {
			case expectFunc:
				readFunc = true
			}
		}
	}

	loc++
	var items []string

	switch {
	case qSingle:
		items = []string{"'"}
	case qDouble:
		items = []string{"\""}
	case expectFunc:
		//fmt.Println(len(line), loc)
		var s string
		if loc < len(line) {
			s = strings.TrimSpace(string(line[loc:]))
		}
		retPos = len(s)
		for name := range proc.GoFunctions {
			if strings.HasPrefix(name, s) {
				items = append(items, name[len(s):])
			}
		}
		if len(items) == 0 {
			items = []string{": "}
		} else {
			sort.Strings(items)
		}
	}

	suggest = make([][]rune, len(items))
	for i := range items {
		suggest[i] = []rune(items[i])
	}

	return
}
