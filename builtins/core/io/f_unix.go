//go:build !windows && !plan9
// +build !windows,!plan9

package io

import (
	"os"

	"github.com/phayes/permbits"
)

func matchFlags(add, remove fFlagsT, info os.FileInfo) bool {
	mode := info.Mode()
	perm := permbits.FileMode(mode)

	return ((add.File() && mode.IsRegular()) ||
		(add.Dir() && mode.IsDir()) ||
		(add.Symlink() && mode&os.ModeSymlink != 0) ||
		(add.DevBlock() && mode&os.ModeDevice != 0) ||
		(add.DevChar() && mode&os.ModeCharDevice != 0) ||
		(add.Socket() && mode&os.ModeSocket != 0) ||
		(add.NamedPipe() && mode&os.ModeNamedPipe != 0) ||

		(add.UserRead() && perm.UserRead()) ||
		(add.GroupRead() && perm.GroupRead()) ||
		(add.OtherRead() && perm.OtherRead()) ||

		(add.UserWrite() && perm.UserWrite()) ||
		(add.GroupWrite() && perm.GroupWrite()) ||
		(add.OtherWrite() && perm.OtherWrite()) ||

		(add.UserExecute() && perm.UserExecute()) ||
		(add.GroupExecute() && perm.GroupExecute()) ||
		(add.OtherExecute() && perm.OtherExecute()) ||

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

			(remove.UserRead() && perm.UserRead()) ||
			(remove.GroupRead() && perm.GroupRead()) ||
			(remove.OtherRead() && perm.OtherRead()) ||

			(remove.UserWrite() && perm.UserWrite()) ||
			(remove.GroupWrite() && perm.GroupWrite()) ||
			(remove.OtherWrite() && perm.OtherWrite()) ||

			(remove.UserExecute() && perm.UserExecute()) ||
			(remove.GroupExecute() && perm.GroupExecute()) ||
			(remove.OtherExecute() && perm.OtherExecute()) ||

			(remove.SetUid() && mode&os.ModeSetuid != 0) ||
			(remove.SetGid() && mode&os.ModeSetgid != 0) ||
			(remove.Sticky() && mode&os.ModeSticky != 0) ||

			(remove.Irregular() && mode&os.ModeIrregular != 0))
}
