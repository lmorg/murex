//go:generate ./update-version.mx

package app

import (
	"fmt"

	"github.com/lmorg/murex/utils/semver"
)

// Name is the name of the $SHELL
const Name = "murex"

// Version number of $SHELL
// Format of version string should be "$(Major).$(Minor).$(Revision) ($Branch)"
const (
	Major     = 7
	Minor     = 0
	Revision  = 2013
	Branch    = "develop"
	BuildDate = "2025-04-05  12:00:49"
)

// Copyright is the copyright owner string
const Copyright = "2018-2025 Laurence Morgan"

// License is the projects software license
const License = "GPL v2"

var licenseFull string

func GetLicenseFull() string  { return licenseFull }
func SetLicenseFull(s string) { licenseFull = s }

const (
	// ShellModule is the name of the module that REPL code gets imported into
	ShellModule  = Name + "/shell"
	ShellProfile = "builtin/"
	UserProfile  = "profile/"
)

func Version() string {
	return fmt.Sprintf("%d.%d.%04d (%s)", Major, Minor, Revision, Branch)
}

func Semver() *semver.Version {
	return &semver.Version{Major, Minor, Revision}
}
