package shell

import (
	"github.com/lmorg/murex/utils/consts"
	"strings"
)

func partialPath(s string) (path, partial string) {
	expanded := expandVariables([]rune(s))
	split := strings.Split(string(expanded), consts.PathSlash)
	path = strings.Join(split[:len(split)-1], consts.PathSlash)
	partial = split[len(split)-1]

	if len(s) > 0 && s[0] == consts.PathSlash[0] {
		path = consts.PathSlash + path
	}

	if path == "" {
		path = "."
	}
	return
}
