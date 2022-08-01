//go:build plan9
// +build plan9

package io

import (
	"os"
)

func matchFlags(add, remove fFlagsT, info os.FileInfo) bool {
	mode := info.Mode()

	return ((add.File() && mode.IsRegular()) ||
		(add.Dir() && mode.IsDir()) ||
		(add.Symlink() && mode&os.ModeSymlink != 0) ||
		(add.DevBlock() && mode&os.ModeDevice != 0) ||
		(add.DevChar() && mode&os.ModeCharDevice != 0) ||
		(add.Socket() && mode&os.ModeSocket != 0) ||
		(add.NamedPipe() && mode&os.ModeNamedPipe != 0) ||

		(add.SetUid() && mode&os.ModeSetuid != 0) ||
		(add.SetGid() && mode&os.ModeSetgid != 0) ||
		(add.Sticky() && mode&os.ModeSticky != 0) ||

		(add.Irregular() && mode&os.ModeIrregular != 0)) &&

		!((remove.File() && mode.IsRegular()) ||
			(remove.Dir() && mode.IsDir()) ||
			(remove.Symlink() && mode&os.ModeSymlink != 0) ||
			(remove.DevBlock() && mode&os.ModeDevice != 0) ||
			(remove.DevChar() && mode&os.ModeCharDevice != 0) ||
			(remove.Socket() && mode&os.ModeSocket != 0) ||
			(remove.NamedPipe() && mode&os.ModeNamedPipe != 0) ||

			(remove.SetUid() && mode&os.ModeSetuid != 0) ||
			(remove.SetGid() && mode&os.ModeSetgid != 0) ||
			(remove.Sticky() && mode&os.ModeSticky != 0) ||

			(remove.Irregular() && mode&os.ModeIrregular != 0))
}
