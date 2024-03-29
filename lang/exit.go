package lang

import (
	"os"

	"github.com/lmorg/murex/utils/cache"
)

var (
	ProfCpuCleanUp func() = func() {}
	ProfMemCleanUp func() = func() {}
)

func Exit(exitNum int) {
	ProfCpuCleanUp()
	ProfMemCleanUp()

	cache.CloseDb()
	os.Exit(exitNum)
}
