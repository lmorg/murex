package generic

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types/define"
	"regexp"
	"strings"
)

var (
	rxWhitespace *regexp.Regexp = regexp.MustCompile(`\s+`)
)

func index(p *proc.Process, params []string) error {
	unmarshaller := func(b []byte) (s []string, err error) {
		s = rxWhitespace.Split(string(b), -1)
		return
	}

	marshaller := func(s []string) (b []byte) {
		b = []byte(strings.Join(s, "\t"))
		return
	}

	return define.IndexTemplateTable(p, params, unmarshaller, marshaller)
}
