package lang

import (
	"os"

	"github.com/lmorg/murex/utils/cache/cachelib"
)

var (
	ProfCpuCleanUp func() = func() {}
	ProfMemCleanUp func() = func() {}
)

func Exit(exitNum int) {
	ProfCpuCleanUp()
	ProfMemCleanUp()

	cachelib.CloseDb()
	os.Exit(exitNum)
}
