// +build windows

package builtins

import _ "github.com/lmorg/murex/builtins/coreutils"

// This is an optional package that covers some of the basic packages you'd expect to find in your POSIX environment.
// It it not aimed at being a complete like for like rewrite of GNU coreutils (for example) nor will it offer the same
// degree of optimisations. So if you are running on a POSIX environment then it is recommend you leave this package
// disabled by default. And if you are running on Windows and you want better compatibility with GNU coreutils or
// performance is a major concern then I would recommend you install either WSL[1], Cygwin[2] or MinGW[3]
//
// [1] https://msdn.microsoft.com/en-gb/commandline/wsl/install_guide
// [2] https://www.cygwin.com/
// [3] http://www.mingw.org/
