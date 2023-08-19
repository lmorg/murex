//go:build !linux && !windows
// +build !linux,!windows

package man

import (
	"fmt"
)

var manBlock = []rune(`
	trypipe {
		/usr/bin/zcat -f ${man -w $command} -> mandoc -O width=%d -c
	}
	catch {
		man <env:MANWIDTH=%d> $command
	}`)

func ManPageExecBlock(width int) []rune {
	return []rune(fmt.Sprintf(string(manBlock), width, width))
}
