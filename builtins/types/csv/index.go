package csv

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/ansi"
)

func readIndex(p *proc.Process, params []string) error {
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
			ansi.Stderrln(ansi.FgRed, err.Error())
		}
		close(cRecords)
	}()

	return define.IndexTemplateTable(p, params, cRecords, csvParser.ArrayToCsv)
}
