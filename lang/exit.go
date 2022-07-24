package lang

import "os"

var (
	ProfCpuCleanUp func() = func() {}
	ProfMemCleanUp func() = func() {}
)

func Exit(exitNum int) {
	ProfCpuCleanUp()
	ProfMemCleanUp()

	os.Exit(exitNum)
}
