//go:build !windows
// +build !windows

package man

type flagsT struct {
	Flags        []string
	Descriptions map[string]string
}
