package path

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/utils/consts"
)

var (
	pathSlashByte   = consts.PathSlash[0]
	pathSlashSlice  = []byte(consts.PathSlash)
	rxCropPathSlash = regexp.MustCompile(fmt.Sprintf(`%s%s+`, consts.PathSlash, consts.PathSlash))
)

func Split(b []byte) ([][]byte, error) {
	if len(b) == 0 {
		return nil, nil
	}

	if b[0] == pathSlashByte {
		b = b[1:]
	}

	if b[len(b)-1] == pathSlashByte {
		b = b[:len(b)-1]
	}

	b = rxCropPathSlash.ReplaceAll(b, pathSlashSlice)
	split := bytes.Split(b, pathSlashSlice)

	return split, nil
}
