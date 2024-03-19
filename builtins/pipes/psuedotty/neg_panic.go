package psuedotty

import "github.com/lmorg/murex/debug"

var panicOnNegDeps = func(i int32) {
	if i < 0 && debug.Enabled {
		panic("More closed dependents than open")
	}
}
