// +build windows

package man

// ScanManPages - windows doesn't have man pages so lets just create an empty
// function that returns a zero-lengthed array so murex can compile on Windows
// but without support for flag auto-detection.
func GetManPages(exe string) []string { return []string{} }

// ParseFlags - windows doesn't have man pages so lets just create an empty
// function that returns a zero-lengthed array so murex can compile on Windows
// but without support for flag auto-detection.
func ParseFlags(paths []string) (flags []string) { return []string{} }

// ParseDescription - windows doesn't have man pages so lets just create an empty
// function that returns a zero-lengthed string so murex can compile on Windows
// but without support for flag auto-detection.
func ParseDescription(paths []string) string { return "" }
