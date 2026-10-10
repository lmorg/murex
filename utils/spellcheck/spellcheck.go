package spellcheck

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/spellcheck/userdictionary"
)

var rxRemoveEsc = regexp.MustCompile(`\\[a-zA-Z]`)

// String spellchecks a line of type string and returns an underlined (ANSI escaped) string
func String(line string) (string, error) {
	enabled, err := lang.ShellProcess.Config.Get("shell", "spellcheck-enabled", types.Boolean)
	if err != nil || !enabled.(bool) {
		return line, err
	}

	block, err := lang.ShellProcess.Config.Get("shell", "spellcheck-func", types.CodeBlock)
	if err != nil || len(block.(string)) == 0 {
		return line, err
	}

	check := rxRemoveEsc.ReplaceAll([]byte(line), []byte{' '})

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_BACKGROUND | lang.F_CREATE_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	fork.Name.Set("(spellcheck)")
	fork.Stdin.SetDataType(types.Generic)
	_, err = fork.Stdin.Writeln(check)
	if err != nil {
		return line, err
	}

	_, err = fork.Execute([]rune(block.(string)))
	if err != nil {
		return line, err
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil {
		return line, err
	}
	if len(b) != 0 {
		return line, fmt.Errorf("`config get shell spellcheck-func` STDERR: %s", string(utils.CrLfTrim(b)))
	}

	err = fork.Stdout.ReadArray(context.Background(), func(bWord []byte) {
		if len(bWord) == 0 {
			return
		}

		sWord := string(bytes.TrimSpace(bWord))

		if (*autocomplete.GlobalExes.Get())[sWord] || lang.MxFunctions.Exists(sWord) || lang.GoFunctions[sWord] != nil || lang.GlobalAliases.Exists(sWord) {
			return
		}

		v, _ := lang.ShellProcess.Variables.GetValue(sWord)
		if v != nil {
			return
		}

		if userdictionary.IsInDictionary(sWord) {
			return
		}

		highlighter(&line, []rune(sWord), highlight)
	})

	return line, err
}

func Exclusions(p *lang.Process) []string {
	exclusions := autocomplete.GlobalExes.List()
	exclusions = append(exclusions, lang.MxFunctions.List()...)
	exclusions = append(exclusions, goFuncList()...)
	exclusions = append(exclusions, lang.GlobalAliases.List()...)
	exclusions = append(exclusions, lang.ListVariables(p)...)
	exclusions = append(exclusions, userdictionary.Get()...)

	pwd, err := os.Getwd()
	if err != nil {
		return exclusions
	}

	files, err := os.ReadDir(pwd)
	if err != nil {
		return exclusions
	}

	for i := range files {
		exclusions = append(exclusions, strings.Split(files[i].Name(), ".")...)
	}

	return exclusions
}

func goFuncList() []string {
	slice := make([]string, len(lang.GoFunctions))
	var i int

	for name := range lang.GoFunctions {
		slice[i] = name
	}

	return slice
}
