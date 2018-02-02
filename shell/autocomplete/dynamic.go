package autocomplete

import (
	"bytes"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"strings"
)

func matchDynamic(f *Flags, partial, exe string, params []string) (items []string) {
	//if len(f.Dynamic) == 0 {
	if f.Dynamic == "" {
		return
	}

	p := &proc.Process{
		Name:       exe,
		Parameters: parameters.Parameters{Params: params},
		Parent:     proc.ShellProcess,
	}
	p.Scope = p

	if !types.IsBlock([]byte(f.Dynamic)) {
		ansi.Stderrln(ansi.FgRed, "Dynamic autocompleter is not a code block.")
		return
	}
	block := []rune(f.Dynamic[1 : len(f.Dynamic)-1])

	stdout := streams.NewStdin()
	stderr := streams.NewStdin()
	exitNum, err := lang.ProcessNewBlock(block, nil, stdout, stderr, p)
	stdout.Close()
	stderr.Close()

	b, _ := stderr.ReadAll()
	s := strings.TrimSpace(string(b))

	if err != nil {
		ansi.Stderrln(ansi.FgRed, "Dynamic autocomplete code could not compile: "+err.Error())
	}
	if exitNum != 0 && debug.Enable {
		ansi.Stderrln(ansi.FgRed, "Dynamic autocomplete returned a none zero exit number.")
	}

	if len(s) > 0 && debug.Enable {
		ansi.Stderrln(ansi.FgRed, utils.NewLineString+s)
	}

	stdout.ReadArray(func(b []byte) {
		s := string(bytes.TrimSpace(b))
		if len(s) == 0 {
			return
		}
		if strings.HasPrefix(s, partial) {
			items = append(items, s[len(partial):])
		}
	})

	return
}
