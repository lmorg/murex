package app

import (
	"fmt"
)

// Name is the name of the $SHELL
const Name = "murex"

// Version number of $SHELL
// Format of version string should be "(major).(minor).(revision) DESCRIPTION"
const (
	version  = "%d.%d.%d"
	Major    = 3
	Minor    = 1
	Revision = 1010
)

var Version string

// Copyright is the copyright owner string
const Copyright = "© 2018-2023 Laurence Morgan"

// License is the projects software license
const License = "License GPL v2"

// ShellModule is the name of the module that REPL code gets imported into
var ShellModule = Name + "/shell"

func init() {
	Version = fmt.Sprintf(version, Major, Minor, Revision)
}
