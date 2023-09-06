package path

import (
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

func Join(a []string) (s string) {
	switch {
	case len(a) == 0:
		return consts.PathSlash

	case a[0] == consts.PathSlash:
		s = strings.Join(a[1:], consts.PathSlash)
		s = consts.PathSlash + s

	default:
		s = strings.Join(a, consts.PathSlash)
	}

	if s != consts.PathSlash {
		if f, _ := os.Stat(s); f != nil && f.IsDir() {
			s += consts.PathSlash
		}
	}
	return s
}
