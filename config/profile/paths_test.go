package profile_test

import (
	"os"
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

func TestProfilePaths(t *testing.T) {
	count.Tests(t, 10)

	var path string
	home := home.MyDir
	temp := t.TempDir()

	// get running settings

	bakPreload := os.Getenv(profile.PreloadEnvVar)
	bakModule := os.Getenv(profile.ModuleEnvVar)
	bakProfile := os.Getenv(profile.ProfileEnvVar)

	defer func() {
		if err := os.Setenv(profile.PreloadEnvVar, bakPreload); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.PreloadEnvVar, bakPreload)
		}

		if err := os.Setenv(profile.ModuleEnvVar, bakModule); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.ModuleEnvVar, bakModule)
		}

		if err := os.Setenv(profile.ProfileEnvVar, bakProfile); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.ProfileEnvVar, bakProfile)
		}
	}()

	// unset env vars (default paths)

	os.Unsetenv(profile.PreloadEnvVar) // don't care about errors
	path = profile.PreloadPath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	os.Unsetenv(profile.ModuleEnvVar) // don't care about errors
	path = profile.ModulePath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected ModulePath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	os.Unsetenv(profile.ProfileEnvVar) // don't care about errors
	path = profile.ProfilePath()
	if !strings.HasPrefix(path, home) {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", home)
		t.Logf("Actual full path: '%s'", path)
	}

	// set env vars (custom paths)

	if err := os.Setenv(profile.PreloadEnvVar, temp); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if err := os.Setenv(profile.ModuleEnvVar, temp); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ModuleEnvVar, err.Error())
	}
	path = profile.ModulePath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected ModulePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if err := os.Setenv(profile.ProfileEnvVar, temp); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if !strings.HasPrefix(path, temp) {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	// set env vars (exact custom file names)

	if err := os.Setenv(profile.PreloadEnvVar, temp+"foobar"); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if path != temp+"foobar" {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected path:  '%s'", temp+"foobar")
		t.Logf("Actual path:    '%s'", path)
	}

	if err := os.Setenv(profile.ProfileEnvVar, temp+"foobar"); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if path != temp+"foobar" {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected path:  '%s'", temp+"foobar")
		t.Logf("Actual path:    '%s'", path)
	}

	// as above but negative test

	if err := os.Setenv(profile.PreloadEnvVar, temp); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}
	path = profile.PreloadPath()
	if path == temp {
		t.Error("Unexpected PreloadPath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}

	if err := os.Setenv(profile.ProfileEnvVar, temp); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}
	path = profile.ProfilePath()
	if path == temp+"foobar" {
		t.Error("Unexpected ProfilePath():")
		t.Logf("Expected prefix:  '%s'", temp)
		t.Logf("Actual full path: '%s'", path)
	}
}

func TestProfileAndCustomPaths(t *testing.T) {
	var (
		preloadFileName = "preload_TestProfileAndCustomPaths.mx"
		modulesPathName = "modules_TestProfileAndCustomPaths.d" // test needs to exclude trailing slash!
		profileFileName = "profile_TestProfileAndCustomPaths.mx"
	)

	path := t.TempDir()

	// get running settings

	bakPreload := os.Getenv(profile.PreloadEnvVar)
	bakModule := os.Getenv(profile.ModuleEnvVar)
	bakProfile := os.Getenv(profile.ProfileEnvVar)

	defer func() {
		if err := os.Setenv(profile.PreloadEnvVar, bakPreload); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.PreloadEnvVar, bakPreload)
		}

		if err := os.Setenv(profile.ModuleEnvVar, bakModule); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.ModuleEnvVar, bakModule)
		}

		if err := os.Setenv(profile.ProfileEnvVar, bakProfile); err != nil {
			t.Errorf("Unable to restore env var settings: '%s' to '%s'", profile.ProfileEnvVar, bakProfile)
		}
	}()

	// set env vars

	if err := os.Setenv(profile.PreloadEnvVar, path+preloadFileName); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}

	if err := os.Setenv(profile.ModuleEnvVar, path+modulesPathName); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ModuleEnvVar, err.Error())
	}

	if err := os.Setenv(profile.ProfileEnvVar, path+profileFileName); err != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}

	// initialize preload

	file, err := os.OpenFile(path+preloadFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", preloadFileName, err.Error())
	}

	_, err = file.WriteString("function: test_preload {}\n")
	if err != nil {
		t.Fatalf("Error initializing %s: %s", preloadFileName, err.Error())
	}

	if file.Close() != nil {
		t.Fatalf("Error closing %s: %s", preloadFileName, err.Error())
	}

	// initialize profile

	file, err = os.OpenFile(path+profileFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", profileFileName, err.Error())
	}

	_, err = file.WriteString("function: test_profile {}\n")
	if err != nil {
		t.Fatalf("Error initializing %s: %s", profileFileName, err.Error())
	}

	if file.Close() != nil {
		t.Fatalf("Error closing %s: %s", profileFileName, err.Error())
	}

	// run tests

	count.Tests(t, 5)

	lang.InitEnv()
	profile.Execute(profile.F_PRELOAD | profile.F_MODULES | profile.F_PROFILE)

	filename := path + modulesPathName
	fi, err := os.Stat(filename)
	if err != nil {
		t.Errorf("Unable to stat '%s': %s", filename, err.Error())
	}
	if !fi.IsDir() {
		t.Errorf("Modules path is not a directory: '%s'", filename)
	}

	filename = path + modulesPathName + consts.PathSlash + "packages.json"
	_, err = os.Stat(filename)
	if err != nil {
		t.Errorf("Unable to stat '%s': %s", filename, err.Error())
	}

	filename = path + modulesPathName + consts.PathSlash + "disabled.json"
	_, err = os.Stat(filename)
	if err != nil {
		t.Errorf("Unable to stat '%s': %s", filename, err.Error())
	}

	if !lang.MxFunctions.Exists("test_preload") {
		t.Errorf("test_preload failed to be defined. Reason: unknown")
		t.Logf("  %v", lang.MxFunctions.Dump())
	}

	if !lang.MxFunctions.Exists("test_profile") {
		t.Errorf("test_profile failed to be defined. Reason: unknown")
		t.Logf("  %v", lang.MxFunctions.Dump())
	}
}
