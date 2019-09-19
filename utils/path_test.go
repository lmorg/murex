package utils

import (
	"runtime"
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/consts"
)

// TestNormalisePath tests NormalisePath function
func TestNormalisePath(t *testing.T) {
	count.Tests(t, 12)

	path := NormalisePath(consts.PathSlash)
	if path != consts.PathSlash {
		t.Error("Root slash, /, (absolute path) not returning itself in NormalisePath")
	}

	path = NormalisePath(consts.PathSlash + "test")
	if path != consts.PathSlash+"test" {
		t.Error(consts.PathSlash + "test (absolute path) not returning itself in NormalisePath")
	}

	path = NormalisePath("test")
	if path == "test" {
		t.Error("test (relative path) is returning itself in NormalisePath")
	}

	path = NormalisePath("test" + consts.PathSlash)
	if path == "test"+consts.PathSlash {
		t.Error("test" + consts.PathSlash + " (relative path) is returning itself in NormalisePath")
	}

	windows := runtime.GOOS == "windows"

	path = NormalisePath("c:/test")
	if windows && path != "c:/test" {
		t.Error("c:/test (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path == "c:/test" {
		t.Error("c:/test (absolute path) is returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath("c:\\test")
	if windows && path != "c:\\test" {
		t.Error("c:\\test (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path == "c:\\test" {
		t.Error("c:\\test (absolute path) is returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath("c:/test/")
	if windows && path != "c:/test/" {
		t.Error("c:/test/ (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path == "c:/test/" {
		t.Error("c:/test/ (absolute path) is returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath("c:\\test\\")
	if windows && path != "c:\\test\\" {
		t.Error("c:\\test\\ (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path == "c:\\test\\" {
		t.Error("c:\\test\\ (absolute path) is returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath(consts.PathSlash + "c:/test/")
	if windows && path != consts.PathSlash+"c:/test/" {
		t.Error(consts.PathSlash + "c:/test/ (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path != consts.PathSlash+"c:/test" {
		t.Error(path + "|" + consts.PathSlash + "c:/test/ (absolute path) not returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath(consts.PathSlash + "c:\\test\\")
	if windows && path != consts.PathSlash+"c:\\test\\" {
		t.Error(consts.PathSlash + "c:\\test\\ (absolute path) not returning itself in NormalisePath in Windows")
	}
	if !windows && path != consts.PathSlash+"c:\\test\\" {
		t.Error(consts.PathSlash + "c:\\test\\ (absolute path) not returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath("test" + consts.PathSlash + "c:/test/")
	if windows && path == "test"+consts.PathSlash+"c:/test/" {
		t.Error("test" + consts.PathSlash + "c:/test/ (relative path) is returning itself in NormalisePath in Windows")
	}
	if !windows && path == "test"+consts.PathSlash+"c:/test" {
		t.Error("test" + consts.PathSlash + "c:/test/ (relative path) is returning itself in NormalisePath on non-Windows")
	}

	path = NormalisePath("test" + consts.PathSlash + "c:\\test\\")
	if windows && path == "test"+consts.PathSlash+"c:\\test\\" {
		t.Error("test" + consts.PathSlash + "c:\\test\\ (relative path) is returning itself in NormalisePath in Windows")
	}
	if !windows && path == "test"+consts.PathSlash+"c:\\test\\" {
		t.Error("test" + consts.PathSlash + "c:\\test\\ (relative path) is returning itself in NormalisePath on non-Windows")
	}

}
