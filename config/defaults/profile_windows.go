//go:build windows
// +build windows

package defaults

import (
	_ "embed"
)

//go:embed profile_windows.mx
var profileWindows []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_windows",
		Block: profileWindows,
	})
}
