package autocomplete

import (
	"bytes"
	"strings"

	"github.com/lmorg/murex/builtins/pipes/streams"

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

func matchDynamic(f *Flags, partial string, args dynamicArgs, act *AutoCompleteT) {
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

	var fStdin int
	pipelineStdout := streams.NewStdin()
	if f.ExecPipeline && !act.ParsedTokens.Unsafe {
		fork := lang.ShellFork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_BACKGROUND | lang.F_NO_STDIN | lang.F_NO_STDERR)
		fork.Stdout = pipelineStdout
		fork.Name = args.exe
		fork.FileRef = ExesFlagsFileRef[args.exe]
		fork.Execute(act.ParsedTokens.Source[:act.ParsedTokens.LastFlowToken])

	} else {
		fStdin = lang.F_NO_STDIN
	}

	fork := lang.ShellFork(lang.F_FUNCTION | lang.F_NEW_MODULE | lang.F_BACKGROUND | fStdin | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.Name = args.exe
	fork.Parameters = parameters.Parameters{Params: args.params}
	fork.FileRef = ExesFlagsFileRef[args.exe]
	if f.ExecPipeline && !act.ParsedTokens.Unsafe {
		fork.Stdin = pipelineStdout
	}
	exitNum, err := fork.Execute(block)

	if err != nil {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete code could not compile: " + err.Error()))
	}
	if exitNum != 0 && debug.Enabled {
		lang.ShellProcess.Stderr.Writeln([]byte("Dynamic autocomplete returned a none zero exit number." + utils.NewLineString))
	}

	if f.Dynamic != "" {
		var items []string
		fork.Stdout.ReadArray(func(b []byte) {
			s := string(bytes.TrimSpace(b))
			if len(s) == 0 {
				return
			}
			if strings.HasPrefix(s, partial) {
				items = append(items, s[len(partial):])
				//act.append(s[len(partial):])
			}
		})

		if f.AutoBranch {
			autoBranch(&items)
		}

		act.append(items...)

	} else {
		if f.ListView {
			//*tdt = readline.TabDisplayList
			act.TabDisplayType = readline.TabDisplayList
		}

		fork.Stdout.ReadMap(lang.ShellProcess.Config, func(key string, value string, last bool) {
			if strings.HasPrefix(key, partial) {
				//items = append(items, key[len(partial):])
				value = strings.Replace(value, "\r", "", -1)
				value = strings.Replace(value, "\n", " ", -1)
				//(*defs)[key[len(partial):]+" "] = value
				//sort.Strings(items)
				act.appendDef(key[len(partial):], value)
			}
		})
	}

	return
}

func execPipeline() {

}
