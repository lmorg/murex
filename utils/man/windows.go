// +build windows

package man

// This file only exists so murex compiles successfully on Windows
func ScanManPages(exe string) []string { return []string{} }
