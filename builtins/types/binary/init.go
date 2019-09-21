package binary

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.SetMime(types.Binary, "multipart/x-zip")

	lang.SetFileExtensions(types.Binary, "bin")
}
