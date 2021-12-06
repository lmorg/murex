package main

import (
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

func diskSource(filename string) ([]byte, error) {
	var b []byte

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			return nil, err
		}
		b, err = ioutil.ReadAll(gz)

		file.Close()
		gz.Close()

		if err != nil {
			return nil, err
		}

	} else {
		b, err = ioutil.ReadAll(file)
		file.Close()
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func execSource(source []rune, sourceRef *ref.Source) {
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN)
	fork.Stdout = new(term.Out)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	if sourceRef != nil {
		fork.FileRef.Source = sourceRef
	}
	fork.RunMode = lang.ShellProcess.RunMode
	exitNum, err := fork.Execute(source)

	if err != nil {
		if exitNum == 0 {
			exitNum = 1
		}
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		os.Exit(exitNum)
	}

	if exitNum != 0 {
		os.Exit(exitNum)
	}
}

func defaultProfile() {
	os.Stderr.WriteString("Loading default profile" + utils.NewLineString)
	source := defaults.DefaultMurexProfile()
	ref := ref.History.AddSource("(builtin)", "source/builtin", []byte(string(source)))
	execSource(defaults.DefaultMurexProfile(), ref)
}
