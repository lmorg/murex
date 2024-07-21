package file

func init() {
	panicOnNegDeps = func(i int32) {
		if i < 0 {
			panic("more closed dependents than open")
		}
	}
}
