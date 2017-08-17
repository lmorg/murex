package shell

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
	"strings"
)

var (
	Instance *readline.Instance // Readline instance
	History  history            // History file
	forward  int
)

// Start interactive shell
func Start() {
	var (
		err       error
		multiline bool
		lines     []string
	)

	Instance, err = readline.NewEx(&readline.Config{
		InterruptPrompt:        "^c",
		AutoComplete:           murexCompleter,
		FuncFilterInputRune:    filterInput,
		DisableAutoSaveHistory: true,
	})

	if err != nil {
		panic(err)
	}

	History, err = newHist(HomeDirectory + ".murex_history")
	if err != nil {
		os.Stderr.WriteString("Error opening history file: " + err.Error())
	}

	Instance.Config.SetListener(listener)
	defer Instance.Close()
	SigHandler()

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
				exitNum, err = lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, proc.ShellProcess)
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
				exitNum, err = lang.ProcessNewBlock([]rune(prompt.(string)), nil, out, nil, proc.ShellProcess)
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
			expanded := expandHistory(block)
			if string(expanded) != string(block) {
				os.Stderr.WriteString(string(expanded) + utils.NewLineString)
				block = expanded
			}

			hist := strings.TrimSpace(strings.Join(lines, " "))
			if History.Last != hist {
				History.Last = hist
				Instance.SaveHistory(hist)
				History.Write(block)
			}

			multiline = false
			lines = make([]string, 0)

			lang.ShellExitNum, _ = lang.ProcessNewBlock(block, nil, nil, nil, proc.ShellProcess)
			streams.CrLf.Write()
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
