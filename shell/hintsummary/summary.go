package hintsummary

import (
	"strings"

	"github.com/lmorg/murex/builtins/docs"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/which"
)

var (
	// Summary is an overriding summary for readline hints
	Summary = New()
)

func Get(cmd string, checkManPage bool) (r []rune) {
	//var summary string
	summary := Cache.Get(cmd)
	if summary != "" {
		return []rune(summary)
	}

	defer func() { Cache.Set(cmd, r) }()

	custom := Summary.Get(cmd)
	if custom != "" {
		summary = custom
	}

	if lang.GlobalAliases.Exists(cmd) {
		a := lang.GlobalAliases.Get(cmd)
		alias := make([]string, len(a))
		copy(alias, a)
		escape.CommandLine(alias)
		s := strings.Join(alias, " ")
		r = []rune("(alias) " + s + " => ")
		cmd = alias[0]
	}

	if lang.MxFunctions.Exists(cmd) {
		if summary == "" {
			summary, _ = lang.MxFunctions.Summary(cmd)
		}

		if summary == "" {
			summary = "no summary written"
		}
		return append(r, []rune("(murex function) "+summary)...)
	}

	if lang.GoFunctions[cmd] != nil {
		if summary == "" {
			synonym := docs.Synonym[cmd]
			summary = docs.Summary[synonym]
		}

		if summary == "" {
			summary = "no doc written"
		}
		r = append(r, []rune("(builtin) "+summary)...)
		return r
	}

	if checkManPage {
		if summary == "" {
			summary = man.ParseSummary(man.GetManPages(cmd))
		}

		if summary == "" {
			summary = "no man page found"
		}

		w := readlink(which.Which(cmd))

		r = append(r, []rune("("+w+") "+summary)...)
		if len(r) > 0 {
			return r
		}
	}

	return nil
}
