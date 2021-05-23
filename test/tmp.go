package test

import (
	"errors"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
)

func TempDir() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", errors.New("env var GOPATH is not set")
	}

	return gopath + consts.PathSlash + strings.Join([]string{"src", "github.com", "lmorg", "murex", "test", "tmp"}, consts.PathSlash) + consts.PathSlash, nil
}
