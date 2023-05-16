package io

//go:generate stringer -linecomment -type=fFlagsT

type fFlagsT rune

const (
	fFile      fFlagsT = 1 << iota // file
	fDirectory                     // directory
	fSymlink                       // symlink
	fDevBlock                      // block device
	fDevChar                       // character device
	fSocket                        // socket
	fNamedPipe                     // named pipe
	fIrregular                     // irregular file

	fUserRead  // user read
	fGroupRead // group read
	fOtherRead // others read

	fUserWrite  // user write
	fGroupWrite // group write
	fOtherWrite // others write

	fUserExec  // user execute
	fGroupExec // group execute
	fOtherExec // others execute

	fSetUid // set uid
	fSetGid // set gid
	fSticky // sticky bit

	fHelp // help
)
