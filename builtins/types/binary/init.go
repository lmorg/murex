package binary

import (
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

func init() {
	define.SetMime(types.Binary, "multipart/x-zip")

	define.SetFileExtensions(types.Binary, "bin")
}
