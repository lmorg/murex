package main

import (
	_ "embed"

	"github.com/lmorg/murex/app"
)

//go:embed LICENSE
var license string

func init() {
	app.SetLicenseFull(license)
}
