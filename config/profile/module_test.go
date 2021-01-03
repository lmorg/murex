package profile

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/test/count"
)

func testCreateModuleStruct() (posix, plan9, windows Module, err error) {
	var pwd string

	pwd, err = os.Getwd()
	if err != nil {
		return
	}

	if runtime.GOOS == "windows" {
		pwd = strings.Replace(pwd, `\`, "/", -1)
	}

	// this is a terrible kludge!
	source := "../../../../../../../../../../.." + pwd + "/module_test.mx"

	posix = Module{
		Name:     "test-module-posix",
		Summary:  "A test module",
		Version:  "1.0",
		Source:   source,
		Package:  "/..",
		Disabled: false,
		Dependencies: Dependencies{
			Optional: []string{"foo", "bar"},
			Required: []string{"sh"},
			Platform: []string{"posix"},
		},
	}

	plan9 = Module{
		Name:     "test-module-plan9",
		Summary:  "A test module",
		Version:  "1.0",
		Source:   source,
		Package:  "/..",
		Disabled: false,
		Dependencies: Dependencies{
			Optional: []string{"foo", "bar"},
			Required: []string{"rc"},
			Platform: []string{"plan9"},
		},
	}

	windows = Module{
		Name:     "test-module-windows",
		Summary:  "A test module",
		Version:  "1.0",
		Source:   source,
		Package:  "/..",
		Disabled: false,
		Dependencies: Dependencies{
			Optional: []string{"foo", "bar"},
			Required: []string{"cmd.exe"},
			Platform: []string{"windows"},
		},
	}

	return
}

func TestIsDisabled(t *testing.T) {
	count.Tests(t, 3)

	disabled = []string{
		"foo",
		"bar",
	}

	test := isDisabled("test")
	foo := isDisabled("foo")
	bar := isDisabled("bar")

	if test {
		t.Errorf("isDisabled true for value not in []string")
	}

	if !foo || !bar {
		t.Errorf("isDisabled false for value in []string")
	}
}

func TestPath(t *testing.T) {
	m, _, _, err := testCreateModuleStruct()
	if err != nil {
		t.Skipf("Unable to get current working directory: %s", err)
	}

	count.Tests(t, 2)

	path := m.Path()

	if strings.Contains(path, `\`) {
		t.Errorf("`\\` found if path (this could break Windows compatibility): %s", path)
		t.Log("(for the sake of standardization, all paths in Murex should be separated by `/` - even on Windows)")
	}

	if strings.Contains(path, "//") {
		t.Errorf("Possible missing path value in %s", path)
	}
}

func TestValidate(t *testing.T) {
	posix, plan9, windows, err := testCreateModuleStruct()
	if err != nil {
		t.Skipf("Unable to get current working directory: %s", err)
	}

	if runtime.GOOS == "windows" {
		t.Fatal("This cannot be tested on Windows because drive letter prefixes")
	}

	_, err = os.Stat(posix.Path())
	if err != nil {
		t.Skip("Unable to stat path. Skipping this test until murex is run for the first time")
	}

	count.Tests(t, 6)

	autocomplete.GlobalExes = map[string]bool{
		"sh":      true,
		"rc":      true,
		"cmd.exe": true,
	}

	errPosix := posix.validate()
	errPlan9 := plan9.validate()
	errWindows := windows.validate()

	if runtime.GOOS != "plan9" && runtime.GOOS != "windows" && errPosix != nil {
		t.Errorf("Failed to validate: %s", err)
	}

	if runtime.GOOS == "plan9" && errPlan9 != nil {
		t.Errorf("Failed to validate: %s", err)
	}

	if runtime.GOOS == "windows" && errWindows != nil {
		t.Errorf("Failed to validate: %s", err)
	}

	if (runtime.GOOS == "plan9" || runtime.GOOS == "windows") && errPosix == nil {
		t.Errorf("posix dependency ignored on non-posix OS")
	}

	if runtime.GOOS != "plan9" && errPlan9 == nil {
		t.Errorf("plan9 dependency ignored on non-plan9 OS")
	}

	if runtime.GOOS != "windows" && errWindows == nil {
		t.Errorf("windows dependency ignored on non-windows OS")
	}
}
