package shell

import (
	"strings"

	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

func previewError(err error, size *readline.PreviewSizeT) ([]string, int, error) {
	s, _, err := previewParse([]byte(err.Error()), size)
	for i := range s {
		s[i] = ansi.ExpandConsts("{RED}") + s[i] + ansi.ExpandConsts("{RESET}") + strings.Repeat(" ", size.Width-len(s[i]))
	}
	return s, 0, err
}

func clErrorCacheMerge(err error, size *readline.PreviewSizeT) ([]string, int, error) {
	s, _, err := previewError(err, size)

	if len(cacheCommandLine) == 0 {
		return s, 0, err
	}

	s = previewHr(s, size)
	return append(s, cacheCommandLine...), 0, err
}

func previewHr(s []string, size *readline.PreviewSizeT) []string {
	return append(s, strings.Repeat("─", size.Width))
}
