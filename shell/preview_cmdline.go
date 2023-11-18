package shell

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

var cacheCommandLine []string

func CommandLine(ctx context.Context, block []rune, _ string, _ bool, size *readline.PreviewSizeT, callback readline.PreviewFuncCallbackT) {
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | lang.F_NEW_MODULE | lang.F_BACKGROUND | lang.F_PREVIEW | lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_CREATE_STDERR)
	fork.FileRef = ref.NewModule(app.ShellModule)

	var err error

	fin := make(chan (bool), 1)
	go func() {
		_, err = fork.Execute(block)
		fin <- true
	}()

	select {
	case <-ctx.Done():
		go fork.KillForks(1)
		fork.Stdout.ForceClose()
		return //nil, 0, nil

	case <-fin:
		// continue
	}

	if err != nil {
		callback(clErrorCacheMerge(err, size))
		return
	}

	b, ioErr := fork.Stdout.ReadAll()
	if fork.Stdout.GetDataType() == types.Json {
		var v interface{}
		err = json.Unmarshal(b, &v)
		if err != nil {
			goto output
		}
		j, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			goto output
		}
		b = j
	}

output:

	if ioErr != nil {
		callback(clErrorCacheMerge(err, size))
		return
	}

	sPreview, i, err := previewParse(b, size)

	b, _ = fork.Stderr.ReadAll()
	if len(b) > 0 {
		if len(sPreview) == 1 && strings.TrimSpace(sPreview[0]) == "" {
			sPreview = []string{}
		}
		if len(sPreview) > 0 {
			sPreview = append(sPreview, strings.Repeat("─", size.Width))
		}
		s, _, _ := previewParse(b, size)
		for i := range s {
			s[i] = ansi.ExpandConsts("{RED}") + s[i] + ansi.ExpandConsts("{RESET}") + strings.Repeat(" ", size.Width-len(s[i]))
		}
		sPreview = append(sPreview, s...)
	}

	cacheCommandLine = sPreview
	callback(sPreview, i, err)
}
