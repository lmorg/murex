//go:build !tinygo
// +build !tinygo

package cmdruntime

import "runtime"

func memNumGC(mem *runtime.MemStats) any {
	return mem.NumGC
}

func memNumForcedGC(mem *runtime.MemStats) any {
	return mem.NumForcedGC
}
