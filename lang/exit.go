package lang

import (
	"os"

	"github.com/lmorg/murex/utils/cache"
)

var (
	ProfCpuCleanUp   func() = func() {}
	ProfMemCleanUp   func() = func() {}
	ProfTraceCleanUp func() = func() {}
)

func Exit(exitNum int) {
	ProfCpuCleanUp()
	ProfMemCleanUp()
	ProfTraceCleanUp()

	cache.CloseDb()
	os.Exit(exitNum)
}
