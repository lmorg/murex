package modules

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins/core/structs"
	"github.com/lmorg/murex/config/profile"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/consts"
)

var (
	// Test Package

	testPackage = "TestPackage"

	testJsonPackage = profile.Package{
		Name:    testPackage,
		Version: "0.0",
	}

	// Test Module 2

	testModule1 = "TestModule1"

	testJsonModule1 = profile.Module{
		Name:    testModule1,
		Summary: "ūnus",
		Version: "1.0",
		Source:  "one.mx",
		Dependencies: profile.Dependencies{
			Required: []string{
				"go",
			},
			Platform: []string{
				"any",
			},
		},
	}

	testFunction1 = "modules.testMxSource1"
	testMxSource1 = "function " + testFunction1 + " {}"

	// Test Module 2

	testModule2 = "TestModule2"

	testJsonModule2 = profile.Module{
		Name:    testModule2,
		Summary: "duo",
		Version: "2.0",
		Source:  "two.mx",
		Dependencies: profile.Dependencies{
			Required: []string{
				"go",
			},
			Platform: []string{
				"any",
			},
		},
	}

	testFunction2 = "modules.testMxSource2"
	testMxSource2 = "function " + testFunction2 + " {}"

	testJsonModules = []profile.Module{
		testJsonModule1,
		testJsonModule2,
	}
)

func testModulesWriteBytes(t *testing.T, name, path string, contents []byte) {
	t.Helper()

	file, err := os.OpenFile(path+name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", name, err.Error())
	}

	_, err = file.Write(contents)
	if err != nil {
		t.Fatalf("Error initializing %s: %s", name, err.Error())
	}

	if file.Close() != nil {
		t.Fatalf("Error closing %s: %s", name, err.Error())
	}
}

func testModulesWriteStruct(t *testing.T, name, path string, v interface{}) {
	t.Helper()

	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		t.Fatalf("Error initializing %s: %s", name, err.Error())
	}

	testModulesWriteBytes(t, name, path, b)
}

func vToString(t *testing.T, v interface{}) string {
	t.Helper()

	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		t.Fatalf("Error creating JSON from %T: %s", v, err.Error())
	}
	return string(b)
}

// TestModulesAndCustomPaths tests a range of functionality from the env var
// custom paths to a lot of the code surrounding loading, enabling and disabling
// modules and packages
func TestModulesAndCustomPaths(t *testing.T) {
	var (
		preloadFileName = "preload_TestModulesAndCustomPaths.mx"
		modulesPathName = "modules_TestModulesAndCustomPaths.d/"
		profileFileName = "profile_TestModulesAndCustomPaths.mx"
	)

	path, err := test.TempDir()
	if err != nil {
		t.Fatalf(err.Error())
	}

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

	if os.Setenv(profile.PreloadEnvVar, path+preloadFileName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.PreloadEnvVar, err.Error())
	}

	if os.Setenv(profile.ModuleEnvVar, path+modulesPathName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ModuleEnvVar, err.Error())
	}

	if os.Setenv(profile.ProfileEnvVar, path+profileFileName) != nil {
		t.Errorf("Unable to set env var %s: %s", profile.ProfileEnvVar, err.Error())
	}

	// initialize empty directory structures

	lang.InitEnv()
	profile.Execute(profile.F_MODULES)

	// initialize test package

	packagePath := path + modulesPathName + consts.PathSlash + "TestPackage" + consts.PathSlash
	if err = os.Mkdir(packagePath, 0777); err != nil && !strings.Contains(err.Error(), "file exists") {
		t.Fatalf("Unable to initialize test package: Cannot create dir: %s", err.Error())
	}

	testModulesWriteStruct(t, "package.json", packagePath, testJsonPackage)
	testModulesWriteStruct(t, "module.json", packagePath, testJsonModules)
	testModulesWriteBytes(t, "one.mx", packagePath, []byte(testMxSource1))
	testModulesWriteBytes(t, "two.mx", packagePath, []byte(testMxSource2))

	// import new packages

	count.Tests(t, 1) // importing from non-standard location
	profile.Execute(profile.F_MODULES)

	if !lang.MxFunctions.Exists(testFunction1) || !lang.MxFunctions.Exists(testFunction2) {
		t.Fatalf("test functions were not imported from test package. Reason: unknown\n%s\nTry deleting '%s' and then rerun",
			vToString(t, lang.MxFunctions.Dump()), path+modulesPathName+consts.PathSlash)
	}

	// run tests

	count.Tests(t, 2)
	list, err := listModulesLoadNotLoad(lang.ShellProcess, true)
	if err != nil {
		t.Fatalf("Error in listModulesLoadNotLoad(true): %s", err.Error())
	}
	if len(list) != 2 || list[testPackage+"/"+testModule1] == "" || list[testPackage+"/"+testModule2] == "" {
		t.Fatalf("listModulesLoadNotLoad(true) has returned an unexpected list:\n%s", vToString(t, list))
	}

	count.Tests(t, 2)
	list, err = listModulesEnDis(lang.ShellProcess, true)
	if err != nil {
		t.Fatalf("Error in listModulesLoadNotLoad(true): %s", err.Error())
	}
	if len(list) != 3 || list[testPackage] == "" ||
		list[testPackage+"/"+testModule1] == "" ||
		list[testPackage+"/"+testModule2] == "" {
		t.Fatalf("listModulesLoadNotLoad(true) has returned an unexpected list:\n%s", vToString(t, list))
	}

	count.Tests(t, 3)
	var disabled []string
	err = profile.ReadJson(profile.ModulePath()+profile.DisabledFile, &disabled)
	if err != nil {
		t.Fatalf("profile.ReadJson() err: %s", err.Error())
	}

	err = disableMod(testPackage+"/"+testModule1, &disabled)
	if err != nil {
		t.Fatalf("disableMod() err: %s", err.Error())
	}
	err = writeDisabled(&disabled)
	if err != nil {
		t.Fatalf("writeDisabled() err: %s", err.Error())
	}

	count.Tests(t, 2)
	list, err = listModulesEnDis(lang.ShellProcess, false)
	if err != nil {
		t.Fatalf("Error in listModulesEnDis(true): %s", err.Error())
	}
	if len(list) != 1 || list[testPackage] != "" ||
		list[testPackage+"/"+testModule1] == "" ||
		list[testPackage+"/"+testModule2] != "" {
		t.Fatalf("listModulesEnDis(true) has returned an unexpected list:\n%s", vToString(t, list))
	}

	count.Tests(t, 3)
	err = profile.ReadJson(profile.ModulePath()+profile.DisabledFile, &disabled)
	if err != nil {
		t.Fatalf("profile.ReadJson() err: %s", err.Error())
	}

	disabled, err = enableMod(testPackage+"/"+testModule1, disabled)
	if err != nil {
		t.Fatalf("enableMod() err: %s", err.Error())
	}
	err = writeDisabled(&disabled)
	if err != nil {
		t.Fatalf("writeDisabled() err: %s", err.Error())
	}

	count.Tests(t, 2)
	list, err = listModulesEnDis(lang.ShellProcess, true)
	if err != nil {
		t.Fatalf("Error in listModulesEnDis(true): %s", err.Error())
	}
	if len(list) != 3 || list[testPackage] == "" ||
		list[testPackage+"/"+testModule1] == "" ||
		list[testPackage+"/"+testModule2] == "" {
		t.Fatalf("listModulesEnDis(true) has returned an unexpected list:\n%s", vToString(t, list))
	}

}
