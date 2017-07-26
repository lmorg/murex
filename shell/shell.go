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
	History  history
	forward  int
)

func Start() {
	var (
		err       error
		multiline bool
		lines     []string
	)

	Instance, err = readline.NewEx(&readline.Config{
		InterruptPrompt:        "^c",
		Stdin:                  osstdin.Stdin,
		AutoComplete:           murexCompleter,
		FuncFilterInputRune:    filterInput,
		DisableAutoSaveHistory: true,
	})
	if err != nil {
		panic(err)
	}

	History, err = openHistFile(HomeDirectory + ".murex_history")
	if err != nil {
		os.Stderr.WriteString("Error opening history file: " + err.Error())
	}

	Instance.Config.SetListener(listener)
	defer Instance.Close()
	defer Instance.Terminal.ExitRawMode()

	for {
		if !multiline {
			var (
				err, err2 error
				exitNum   int
				b         []byte
			)

			proc.GlobalVars.Set("linenum", 1, types.Number)
			prompt, err := proc.GlobalConf.Get("shell", "prompt", types.CodeBlock)
			if err == nil {
				out := streams.NewStdin()
				exitNum, err = lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, "shell")
				out.Close()

				b, err2 = out.ReadAll()
				if len(b) > 1 && b[len(b)-1] == '\n' {
					b = b[:len(b)-1]
				}

				if len(b) > 1 && b[len(b)-1] == '\r' {
					b = b[:len(b)-1]
				}
			}

			if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
				os.Stderr.WriteString("Invalid prompt. Block returned false." + utils.NewLineString)
				b = []byte("murex » ")
			}

			Instance.SetPrompt(string(b))
		} else {
			var (
				err, err2 error
				exitNum   int
				b         []byte
			)

			proc.GlobalVars.Set("linenum", len(lines)+1, types.Number)
			prompt, err := proc.GlobalConf.Get("shell", "prompt-multiline", types.CodeBlock)
			if err == nil {
				out := streams.NewStdin()
				exitNum, err = lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, "shell")
				out.Close()

				b, err2 = out.ReadAll()
				if len(b) > 1 && b[len(b)-1] == '\n' {
					b = b[:len(b)-1]
				}

				if len(b) > 1 && b[len(b)-1] == '\r' {
					b = b[:len(b)-1]
				}
			}

			if exitNum != 0 || err != nil || len(b) == 0 || err2 != nil {
				os.Stderr.WriteString("Invalid prompt. Block returned false." + utils.NewLineString)
				b = []byte(fmt.Sprintf("%5d » ", len(lines)+1))
			}

			Instance.SetPrompt(string(b))
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

		lines = append(lines, line)
		block := []rune(strings.Join(lines, "\n"))
		_, pErr := lang.ParseBlock(block)
		switch {
		case pErr.Code == lang.ErrUnterminatedBrace,
			pErr.Code == lang.ErrUnterminatedEscape,
			pErr.Code == lang.ErrUnterminatedQuotesSingle:
			multiline = true
		case len(block) == 0:
			continue
		default:
			History.Last = strings.Join(lines, " ")
			multiline = false
			lines = make([]string, 0)
			Instance.SaveHistory(History.Last)
			History.Write(block)
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
