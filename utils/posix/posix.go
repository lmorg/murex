package posix

import "runtime"

// IsPosix returns `true` when running on a POSIX host
func IsPosix() bool {
	return isPosix(runtime.GOOS)
}

func isPosix(os string) bool {
	return os != "windows" && os != "plan9"
}
