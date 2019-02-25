package autocomplete

import (
	"bytes"
	"sort"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/readline"
)

type dynamicArgs struct {
	exe    string
	params []string
	float  int
}

func matchDynamic(f *Flags, partial string, args dynamicArgs, defs *map[string]string, tdt *readline.TabDisplayType) (items []string) {
	// Default to building up from Dynamic field. Fall back to DynamicDefs
	dynamic := f.Dynamic
	if f.Dynamic == "" {
		dynamic = f.DynamicDesc
	}
	if dynamic == "" {
		return
	}

	if !types.IsBlock([]byte(dynamic)) {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocompleter is not a code block"))
		return
	}
	block := []rune(dynamic[1 : len(dynamic)-1])

	//branch := lang.ShellProcess.BranchFID()
	//branch.Scope = branch.Process
	//branch.Parent = branch.Process
	//branch.IsBackground = true
	//branch.Name = args.exe
	//branch.Parameters = parameters.Parameters{Params: args.params}
	//defer branch.Close()

	//stdout := streams.NewStdin()
	//exitNum, err := lang.RunBlockNewConfigSpace(block, nil, stdout, nil, branch.Process)

	fork := lang.ShellProcess.Fork(lang.F_NEW_CONFIG | lang.F_NEW_TESTS | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name = args.exe
	fork.Parameters = parameters.Parameters{Params: args.params}
	exitNum, err := fork.Execute(block)

	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete code could not compile: " + err.Error()))
	}
	if exitNum != 0 && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete returned a none zero exit number." + utils.NewLineString))
	}

	if f.Dynamic != "" {
		fork.Stdout.ReadArray(func(b []byte) {
			s := string(bytes.TrimSpace(b))
			if len(s) == 0 {
				return
			}
			if strings.HasPrefix(s, partial) {
				items = append(items, s[len(partial):])
			}
		})

	} else {
		if f.ListView {
			*tdt = readline.TabDisplayList
		}

		fork.Stdout.ReadMap(lang.ShellProcess.Config, func(key string, value string, last bool) {
			if strings.HasPrefix(key, partial) {
				items = append(items, key[len(partial):])
				value = strings.Replace(value, "\r", "", -1)
				value = strings.Replace(value, "\n", " ", -1)
				(*defs)[key[len(partial):]+" "] = value
				sort.Strings(items)
			}
		})
	}

	if f.AutoBranch {
		autoBranch(items)
		items = dedup(items)
	}

	return
}

func autoBranch(tree []string) {
	//debug.Json("tree", tree)
	for branch := range tree {

		for i := 0; i < len(tree[branch])-1; i++ {
			if tree[branch][i] == '/' {
				tree[branch] = tree[branch][:i+1]
				break
			}
		}

	}
}

func dedup(items []string) []string {
	m := make(map[string]bool)
	for i := range items {
		m[items[i]] = true
	}

	new := []string{}
	for s := range m {
		new = append(new, s)
	}

	sort.Strings(new)
	return new
}
