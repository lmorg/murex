package shell

import (
	"fmt"
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

var (
	Instance *readline.Instance
	forward  int
)

func Start() {
	var (
		err       error
		multiline bool
		lines     []string
	)

	Instance, err = readline.NewEx(&readline.Config{
		HistoryFile:         HomeDirectory + ".murex_history",
		InterruptPrompt:     "^c",
		Stdin:               osstdin.Stdin,
		AutoComplete:        murexCompleter,
		FuncFilterInputRune: filterInput,
		//DisableAutoSaveHistory: true,
	})

	if err != nil {
		panic(err)
	}

	Instance.Config.SetListener(listener)
	defer Instance.Close()
	defer Instance.Terminal.ExitRawMode()

	for {
		if !multiline {
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
		} else {
			Instance.SetPrompt(fmt.Sprintf("%d » ", len(lines)+1))
		}

		line, err := Instance.Readline()
		if err == readline.ErrInterrupt {
			if multiline {
				multiline = false
				lines = make([]string, 0)
				continue
			}
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		lines = append(lines, line+"\n")
		block := []rune(strings.Join(lines, ""))
		_, pErr := lang.ParseBlock(block)
		switch {
		case pErr.Code == lang.ErrUnterminatedBrace,
			pErr.Code == lang.ErrUnterminatedEscape,
			pErr.Code == lang.ErrUnterminatedQuotesSingle:
			multiline = true
		default:
			multiline = false
			lines = make([]string, 0)
			Instance.Terminal.EnterRawMode()
			lang.ShellExitNum, _ = lang.ProcessNewBlock(block, nil, nil, nil, "shell")
			Instance.Terminal.ExitRawMode()
		}
	}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	case readline.CharForward:
		forward++
		return r, true
	}
	return r, true
}
