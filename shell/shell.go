package shell

import (
	"github.com/chzyer/readline"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/osstdin"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"strings"
)

var Instance *readline.Instance

func Start() {
	var err error

	Instance, err = readline.NewEx(&readline.Config{
		HistoryFile:         HomeDirectory + ".murex_history",
		AutoComplete:        murexCompleter,
		InterruptPrompt:     "^c",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
		Stdin:               osstdin.Stdin,
	})
	defer Instance.Terminal.ExitRawMode()

	if err != nil {
		panic(err)
	}
	defer Instance.Close()

	for {
		prompt, _ := proc.GlobalConf.Get("shell", "prompt", types.CodeBlock)
		out := streams.NewStdin()
		exitNum, err := lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, "shell")
		out.Close()

		b, err2 := out.ReadAll()
		if len(b) > 1 && b[len(b)-1] == '\n' {
			b = b[:len(b)-1]
		}

		if len(b) > 1 && b[len(b)-1] == '\r' {
			b = b[:len(b)-1]
		}

		if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
			os.Stderr.WriteString("Invalid prompt. Block returned false." + utils.NewLineString)
			b = []byte("murex » ")
		}

		Instance.SetPrompt(string(b))

		line, err := Instance.Readline()
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
			Instance.Terminal.EnterRawMode()
			lang.ShellExitNum, _ = lang.ProcessNewBlock(
				[]rune(line),
				nil,
				nil,
				nil,
				"shell",
			)
			Instance.Terminal.ExitRawMode()
		}
	}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
