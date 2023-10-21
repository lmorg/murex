//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package preview

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/builtins/pipes/psuedotty"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

var cacheCommandLine []string

func CommandLine(ctx context.Context, block []rune, _ string, _ bool, size *readline.PreviewSizeT) ([]string, int, error) {
	usePty := false

	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_BACKGROUND | lang.F_PREVIEW | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR)
	fork.FileRef = ref.NewModule(app.ShellModule)

	var (
		err, ioErr error
		buf        *bytes.Buffer
	)

	if usePty {
		fork.Stdout, err = psuedotty.NewPTY(size.Width, size.Height)
		if err != nil {
			panic("TODO")
		}

		fork.Stdout.Open()

		buf = bytes.NewBuffer(nil)
		go func() {
			_, ioErr = io.Copy(buf, fork.Stdout)
		}()
	}

	fork.Stderr = fork.Stdout

	fin := make(chan (bool), 1)
	go func() {
		_, err = fork.Execute(block)
		fin <- true
	}()

	select {
	case <-ctx.Done():
		go fork.KillForks(1)
		fork.Stdout.ForceClose()
		return nil, 0, nil

	case <-fin:
		// continue
	}

	if err != nil {
		return clErrorCacheMerge(err, size)
	}

	var b []byte
	if usePty {
		b = buf.Bytes()
	} else {
		b, ioErr = fork.Stdout.ReadAll()
	}
	if ioErr != nil {
		return clErrorCacheMerge(err, size)
	}

	s, i, err := parse(b, size)
	cacheCommandLine = s
	return s, i, err
}

func clErrorCacheMerge(err error, size *readline.PreviewSizeT) ([]string, int, error) {
	s, _, err := parse([]byte(err.Error()), size)
	for i := range s {
		s[i] = ansi.ExpandConsts("{RED}") + s[i] + ansi.ExpandConsts("{RESET}") + strings.Repeat(" ", size.Width-len(s[i]))
	}

	if len(cacheCommandLine) == 0 {
		return s, 0, err
	}

	s = append(s, strings.Repeat("─", size.Width))
	return append(s, cacheCommandLine...), 0, err
}
