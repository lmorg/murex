//go:build linux
// +build linux

package man

import "fmt"

var manBlock = `man <env:MANWIDTH=%d> $command`

func ManPageExecBlock(width int) []rune {
	return []rune(fmt.Sprintf(manBlock, width))
}
