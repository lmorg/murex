//go:generate ./update-version.mx

package app

import (
	"fmt"

	"github.com/lmorg/murex/utils/semver"
)

// Name is the name of the $SHELL
const Name = "murex"

// Version number of $SHELL
// Format of version string should be "(major).(minor).(revision) BRANCH-NAME"
const (
	version  = "%d.%d.%04d"
	Major    = 6
	Minor    = 3
	Revision = 587
	Branch   = "863/job-control"
)

// Copyright is the copyright owner string
const Copyright = "© 2018-2024 Laurence Morgan"

// License is the projects software license
const License = "License GPL v2"

// ShellModule is the name of the module that REPL code gets imported into
var ShellModule = Name + "/shell"

func init() {
	v = fmt.Sprintf(version, Major, Minor, Revision)
	sv, _ = semver.Parse(v)
}

var v string

func Version() string {
	return fmt.Sprintf("%s %s", v, Branch)
}

var sv *semver.Version

func Semver() *semver.Version {
	return sv
}
