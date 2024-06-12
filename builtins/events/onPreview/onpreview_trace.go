//go:build trace
// +build trace

package onpreview

func isValidInterruptDebug(interrupt string) {
	if err := isValidInterrupt(interrupt); err != nil {
		panic(err.Error())
	}
}
