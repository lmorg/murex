package gopath

import (
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

// Source returns the absolute path to a murex packages source. It uses GOPATH
// environmental variable when it is available.
//
// (packagePath []string) refers to the relative path of a specific Go package
// within the murex source tree
func Source(packagePath []string) string {
	var err error

	path := []string{
		"src",
		"github.com",
		"lmorg",
		"murex",
	}

	// This should work in most cases but not everyone has GOPATH set these days
	GOPATH := os.Getenv("GOPATH")
	if GOPATH != "" {
		GOPATH += consts.PathSlash + strings.Join(path, consts.PathSlash)

		// OK, GOPATH isn't set but we might be able to have an accurate
		// guess based on the current working directory (we hope!)
	} else {
		GOPATH, err = os.Getwd()

		// Lets guess the GOPATH based on a common install location
		// (this is an all else fails approach - not ideal but the former
		// two tests should work)
		if err != nil {
			GOPATH = home.MyDir + consts.PathSlash + "go" + consts.PathSlash + strings.Join(path, consts.PathSlash)
		}
	}

	return GOPATH + consts.PathSlash + strings.Join(packagePath, consts.PathSlash) + consts.PathSlash

}
