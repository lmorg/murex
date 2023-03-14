package path

import (
	"path"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

func Split(s string) []string {
	if len(s) == 0 {
		return []string{"."}
	}

	s = path.Clean(s)

	split := strings.Split(s, consts.PathSlash)

	if len(split) == 0 {
		// this should never happen
		return []string{"."}
	}

	if len(split) > 0 && split[0] == "" {
		split[0] = consts.PathSlash
	}

	return split
}
