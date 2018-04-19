// +build windows

package consts

// PathSlash is an OS specific directory separator.
// Normally in Windows this would be a \ but lets standardise everything in murex to be /
const PathSlash = "/"

// TempDir is the location of temp directory
const TempDir = `c:\temp\murex\`
