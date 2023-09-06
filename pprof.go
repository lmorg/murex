//go:build pprof
// +build pprof

package main

import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
)

const (
	fCpuProfile = "./cpu.pprof"
	fMemProfile = "./mem.pprof"
)

func init() {
	lang.ProfCpuCleanUp = cpuProfile()
	lang.ProfMemCleanUp = memProfile()
}

func cpuProfile() func() {
	if fCpuProfile != "" {
		//fmt.Fprintf(tty.Stderr, "Writing CPU profile to '%s'\n", fCpuProfile)

		f, err := os.Create(fCpuProfile)
		if err != nil {
			panic(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		return func() {
			pprof.StopCPUProfile()
			if err = f.Close(); err != nil && debug.Enabled {
				panic(err)
			}

			//fmt.Fprintf(tty.Stderr, "CPU profile written to '%s'\n", fCpuProfile)
		}
	}

	return func() {}
}

func memProfile() func() {
	if fMemProfile != "" {
		//fmt.Fprintf(tty.Stderr, "Writing memory profile to '%s'\n", fMemProfile)

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
			//fmt.Fprintf(tty.Stderr, "Memory profile written to '%s'\n", fMemProfile)
		}
	}

	return func() {}
}
