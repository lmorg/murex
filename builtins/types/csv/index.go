package csv

import (
	"github.com/lmorg/murex/lang/proc"
)

func readIndex(p *proc.Process, params []string) error {
	match := make(map[string]int)
	for i := range params {
		match[params[i]] = i + 1
	}

	csvParser, err := NewParser(nil, &proc.GlobalConf)
	if err != nil {
		return err
	}
	records := make([]string, len(params)+1)
	var matched bool

	err = p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
		if match[key] != 0 {
			matched = true
			records[match[key]] = value
		}

		if last && matched {
			p.Stdout.Writeln(csvParser.ArrayToCsv(records[1:]))
			matched = false
			records = make([]string, len(params)+1)
		}
	})

	return err
}
