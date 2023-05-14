package io

import "fmt"

func (f fFlagsT) File() bool      { return f&fFile != 0 }
func (f fFlagsT) Dir() bool       { return f&fDirectory != 0 }
func (f fFlagsT) Symlink() bool   { return f&fSymlink != 0 }
func (f fFlagsT) DevBlock() bool  { return f&fDevBlock != 0 }
func (f fFlagsT) DevChar() bool   { return f&fDevChar != 0 }
func (f fFlagsT) Socket() bool    { return f&fSocket != 0 }
func (f fFlagsT) NamedPipe() bool { return f&fNamedPipe != 0 }
func (f fFlagsT) Irregular() bool { return f&fIrregular != 0 }

func (f fFlagsT) UserRead() bool  { return f&fUserRead != 0 }
func (f fFlagsT) GroupRead() bool { return f&fGroupRead != 0 }
func (f fFlagsT) OtherRead() bool { return f&fOtherRead != 0 }

func (f fFlagsT) UserWrite() bool  { return f&fUserWrite != 0 }
func (f fFlagsT) GroupWrite() bool { return f&fGroupWrite != 0 }
func (f fFlagsT) OtherWrite() bool { return f&fOtherWrite != 0 }

func (f fFlagsT) UserExecute() bool  { return f&fUserExec != 0 }
func (f fFlagsT) GroupExecute() bool { return f&fGroupExec != 0 }
func (f fFlagsT) OtherExecute() bool { return f&fOtherExec != 0 }

func (f fFlagsT) SetUid() bool { return f&fSetUid != 0 }
func (f fFlagsT) SetGid() bool { return f&fSetGid != 0 }
func (f fFlagsT) Sticky() bool { return f&fSticky != 0 }

var fFlagLookup = map[rune]fFlagsT{
	'f': fFile | fSymlink | fDevBlock | fDevChar | fSocket | fNamedPipe | fIrregular,
	'F': fFile,
	'd': fDirectory | fSymlink,
	'D': fDirectory,
	's': fSymlink,
	'l': fSymlink,
	'b': fDevBlock,
	'c': fDevChar,
	'S': fSocket,
	'p': fNamedPipe,
	'?': fIrregular,

	'r': fUserRead | fGroupRead | fOtherRead,
	'R': fUserRead,
	'e': fGroupRead,
	'E': fOtherRead,

	'w': fUserWrite | fGroupWrite | fOtherWrite,
	'W': fUserWrite,
	'Q': fGroupWrite,
	'q': fOtherWrite,

	'x': fUserExec | fGroupExec | fOtherExec,
	'X': fUserExec,
	'Z': fGroupExec,
	'z': fOtherExec,

	'u': fSetUid,
	'g': fSetGid,
	't': fSticky,

	'h': fHelp,
}

func fFlagsParser(param []rune, add fFlagsT, remove fFlagsT) (fFlagsT, fFlagsT, error) {
	if len(param) < 2 {
		return add, remove, fmt.Errorf("invalid parameter '%s'", string(param))
	}

	var mode bool

	switch param[0] {
	case '-':
		// already assigned as
		// mode = false

	case '+':
		mode = true

	default:
		return add, remove, fmt.Errorf("flags should begin either with '-' or '+'. Instead got '%s'", string([]rune{param[0]}))
	}

	for i := 1; i < len(param); i++ {
		f := fFlagLookup[param[i]]
		if f == 0 {
			return add, remove, fmt.Errorf("invalid flag: '%s'", string([]rune{param[i]}))
		}

		if mode {
			add |= f
		} else {
			remove |= f
		}
	}

	return add, remove, nil
}
