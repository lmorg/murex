package shell

import (
	"github.com/chzyer/readline"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Start() {
	rl, err := readline.NewEx(&readline.Config{
		//Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "murex.history",
		AutoComplete:    createCompleter(),
		InterruptPrompt: "^c",
		//EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})

	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		prompt, _ := proc.GlobalConf.Get("shell", "Prompt", types.CodeBlock)
		out := streams.NewStdin()
		exitNum, err := lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, types.Null)
		out.Close()
		b := out.ReadAll()
		if exitNum != 0 || err != nil {
			os.Stderr.WriteString("Invalid prompt. Block returned false." + utils.NewLineString)
			b = []byte("murex » ")
		}

		if b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}

		if b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}

		rl.SetPrompt(string(b))

		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case line == "":
		default:
			lang.ProcessNewBlock(
				[]rune(line),
				nil,
				nil,
				nil,
				types.Null,
			)

		}
	}
}

func dynFunctions() func(string) []string {
	return func(line string) (items []string) {
		for name := range proc.GoFunctions {
			if proc.GoFunctions[name].TypeIn == types.Null || proc.GoFunctions[name].TypeIn == types.Generic {
				items = append(items, name+":")
			}
		}
		return
	}
}

func dynMethods() func(string) []string {
	return func(line string) (items []string) {
		for name := range proc.GoFunctions {
			//if proc.GoFunctions[name].TypeIn == types.Null || proc.GoFunctions[name].TypeIn == types.Generic {
			items = append(items, "-> "+name)
			//}
		}
		return
	}
}

func dynParameters() func(string) []string {
	return func(line string) (items []string) {
		items = listFiles("./")
		//items = append([]string{")"}, items...)
		return
	}
}

func listFiles(path string) (filenames []string) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		filenames = append(filenames, f.Name())
	}
	return
}

func createCompleter() (completer *readline.PrefixCompleter) {
	newFunc := readline.PcItemDynamic(dynFunctions())

	newMethod := []readline.PrefixCompleterInterface{
		readline.PcItem("->", readline.PcItemDynamic(dynMethods())),
		readline.PcItem("|"),
		readline.PcItem("?"),
		readline.PcItem(";"),
	}

	params := readline.PcItemDynamic(
		dynParameters(),
		readline.PcItem(" "),
		readline.PcItem("->", newMethod...),
		readline.PcItem("|"),
		readline.PcItem("?"),
		readline.PcItem(";"),
		//readline.PcItemDynamic(dynMethods()),
	)

	newFunc.SetChildren([]readline.PrefixCompleterInterface{params})
	newMethod[2].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[0].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[1].SetChildren([]readline.PrefixCompleterInterface{params})
	params.Children[2].SetChildren([]readline.PrefixCompleterInterface{params})

	/*newFunc := readline.PcItemDynamic(dynComplete())
	newFunc.SetChildren([]readline.PrefixCompleterInterface{newFunc})*/

	completer = readline.NewPrefixCompleter(newFunc)

	//completer.Children[0].GetChildren()[0].
	return
}

/*var completer2 = readline.NewPrefixCompleter(
	readline.PcItem("text(",
		readline.PcItemDynamic(listFiles("./"),
			readline.PcItem(","),
			readline.PcItem(")",
				readline.PcItem(".", compChild),
				readline.PcItem("|", compChild),
				readline.PcItem("?", compChild),
				compChild,
			),
		),
	),
)*/

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
