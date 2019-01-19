package csv

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types/define"
)

func readIndex(p *lang.Process, params []string) error {
	cRecords := make(chan []string, 1)

	csvParser, err := NewParser(p.Stdin, p.Config)
	if err != nil {
		return err
	}

	go func() {
		var headingsPrinted bool
		err := csvParser.ReadLine(func(recs []string, headings []string) {
			if !headingsPrinted {
				cRecords <- headings
				headingsPrinted = true
			}
			cRecords <- recs
		})
		if err != nil {
			//ansi.Stderrln(p, ansi.FgRed, err.Error())
			p.Stderr.Writeln([]byte(err.Error()))
		}
		close(cRecords)
	}()

	return define.IndexTemplateTable(p, params, cRecords, csvParser.ArrayToCsv)
}
