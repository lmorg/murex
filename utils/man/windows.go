// +build windows

package man

// Windows doesn't have man pages so lets just create an empty function that returns a zero-lengthed array so murex can
// compile on Windows but without support for flag auto-detection.
func ScanManPages(exe string) []string { return []string{} }
