package main

import (
	"testing"

	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
)

func TestGoPath(t *testing.T) {
	count.Tests(t, 1)

	pwd := goPath()

	if pwd == "" {
		t.Error("$GOPATH env var appears to be unset. This is needed for wasmserver to locate which directory to serve")
	}
}

func TestPathBuilder(t *testing.T) {
	count.Tests(t, 1)

	test.Exists(t, pathBuilder()+"/index.html")
}
