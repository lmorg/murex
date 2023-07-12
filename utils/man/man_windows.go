//go:build windows
// +build windows

package man

import (
	"github.com/lmorg/murex/lang/stdio"
)

// ScanManPages - windows doesn't have man pages so lets just create an empty
// function that returns a zero-length array so murex can compile on Windows
// but without support for flag auto-detection.
func GetManPages(_ string) []string { return []string{} }

// ParseByPaths - windows doesn't have man pages so lets just create an empty
// function that returns a zero-length array so murex can compile on Windows
// but without support for flag auto-detection.
func ParseByPaths(_ []string) ([]string, map[string]string) { return []string{}, map[string]string{} }

// ParseByStdio - windows doesn't have man pages so lets just create an empty
// function that returns a zero-length array so murex can compile on Windows
// but without support for flag auto-detection.
func ParseByStdio(_ stdio.Io) ([]string, map[string]string) { return []string{}, map[string]string{} }

// ParseDescription - windows doesn't have man pages so lets just create an empty
// function that returns a zero-length string so murex can compile on Windows
// but without support for flag auto-detection.
func ParseSummary(_ []string) string { return "" }
