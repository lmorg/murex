//go:build !linux && !windows
// +build !linux,!windows

package man

var manBlock = []rune(`
	trypipe {
		/usr/bin/zcat -f ${man -w $command} -> mandoc -O width=1000 -c
	}
	catch {
		man $command
	}`)
