//go:build pprof
// +build pprof

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/lmorg/murex/lang"
)

const (
	fCpuProfile = "./cpu.pprof"
	fMemProfile = "./mem.pprof"
)

func init() {
	lang.ProfCpuCleanUp = cpuProfile
	lang.ProfMemCleanUp = memProfile
}

func cpuProfile() {
	fmt.Fprintf(os.Stderr, "Writing CPU profile to '%s'\n", fCpuProfile)

	f, err := os.Create(fCpuProfile)
	if err != nil {
		panic(err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}

	return func() {
		pprof.StopCPUProfile()
		if err = f.Close(); err != nil {
			panic(err)
		}

		fmt.Fprintf(os.Stderr, "CPU profile written to '%s'\n", fCpuProfile)
	}
}

func memProfile() {
	fmt.Fprintf(os.Stderr, "Writing memory profile to '%s'\n", fMemProfile)

	f, err := os.Create(fMemProfile)
	if err != nil {
		panic(err)
	}

	return func() {
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(err)
		}
		if err = f.Close(); err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "Memory profile written to '%s'\n", fMemProfile)
	}
}
