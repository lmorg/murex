package preview

import (
	"context"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

var cacheCommandLine []string

func CommandLine(ctx context.Context, block []rune, _ string, _ bool, size *readline.PreviewSizeT) ([]string, int, error) {
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_NO_STDIN | lang.F_FAKETTY_STDOUT | lang.F_BACKGROUND | lang.F_PREVIEW)
	fork.FileRef = ref.NewModule(app.ShellModule)
	fork.Stderr = fork.Stdout

	var err error
	fin := make(chan (bool), 1)
	go func() {
		_, err = fork.Execute(block)
		fin <- true
	}()

	select {
	case <-ctx.Done():
		fork.KillForks(1)
		return nil, 0, nil

	case <-fin:
		// continue
	}

	if err != nil {
		return clErrorCacheMerge(err, size)
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
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
