package utils_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils"
)

func TestExists(t *testing.T) {
	count.Tests(t, 2)

	if !utils.Exists("README.md") {
		t.Error(`Exists("README.md") == false when it should have been true`)
	}

	if utils.Exists("lkjdslksdjflsdkjflsdkjflsdkfjsldkjfds") {
		t.Error(`Exists("lkjdslksdjflsdkjflsdkjflsdkfjsldkjfds") == true when it should have been false`)
	}
}
