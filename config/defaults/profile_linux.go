//go:build linux
// +build linux

package defaults

import (
	_ "embed"
)

//go:embed profile_linux.mx
var profileLinux []byte

func init() {
	DefaultProfiles = append(DefaultProfiles, &DefaultProfileT{
		Name:  "profile_linux.mx",
		Block: profileLinux,
	})
}
