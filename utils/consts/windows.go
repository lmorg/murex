// +build windows

package consts

// PathSlash is an OS specific directory separator.
// Normally in Windows this would be a \ but lets standardise everything in murex to be /
const PathSlash = "/"

// tempDir is the location of temp directory if it cannot be automatically determind
const tempDir = `c:/temp/murex/`
