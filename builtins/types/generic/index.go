package generic

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/ansi"
	"regexp"
	"strings"
)

var rxWhitespace *regexp.Regexp = regexp.MustCompile(`\s+`)

func index(p *proc.Process, params []string) error {
	cRecords := make(chan []string, 1)

	go func() {
		err := p.Stdin.ReadLine(func(b []byte) {
			cRecords <- rxWhitespace.Split(string(b), -1)
		})
		if err != nil {
			ansi.Stderrln(ansi.FgRed, err.Error())
		}
		close(cRecords)
	}()

	marshaller := func(s []string) (b []byte) {
		b = []byte(strings.Join(s, "\t"))
		return
	}

	return define.IndexTemplateTable(p, params, cRecords, marshaller)
}
