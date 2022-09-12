//go:build !windows
// +build !windows

package which

import "strings"

// SplitPath takes a $PATH (or similar) string and splits it into an array of paths
func SplitPath(envPath string) []string {
	split := strings.Split(envPath, ":")
	return split
}
