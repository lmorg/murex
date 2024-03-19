//go:build pprof
// +build pprof

package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
)

const (
	fCpuProfile   = "./cpu.pprof"
	fMemProfile   = "./mem.pprof"
	fTraceProfile = "./trace.pprof"
)

func init() {
	lang.ProfCpuCleanUp = cpuProfile()
	lang.ProfMemCleanUp = memProfile()
	lang.ProfTraceCleanUp = traceProfile()
}

func cpuProfile() func() {
	if fCpuProfile != "" {

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
		}
	}

	return func() {}
}

func memProfile() func() {
	if fMemProfile != "" {

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
		}
	}

	return func() {}
}

func traceProfile() func() {
	if fTraceProfile != "" {

		f, err := os.Create(fTraceProfile)
		if err != nil {
			panic(err)
		}
		if err := trace.Start(f); err != nil {
			panic(err)
		}

		return func() {
			trace.Stop()
			if err = f.Close(); err != nil {
				panic(err)
			}
		}
	}

	return func() {}

}
