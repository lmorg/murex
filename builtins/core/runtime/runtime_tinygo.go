//go:build tinygo
// +build tinygo

package cmdruntime

import "runtime"

const unsupportedTinyGo = "unsupported in TinyGo"

func memNumGC(mem *runtime.MemStats) any {
	return unsupportedTinyGo
}

func memNumForcedGC(mem *runtime.MemStats) any {
	return unsupportedTinyGo
}
