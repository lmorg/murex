//go:build trace
// +build trace

package onprompt

func isValidInterruptDebug(interrupt string) {
	if err := isValidInterrupt(interrupt); err != nil {
		panic(err.Error())
	}
}
