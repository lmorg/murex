package profilepaths

import (
	"fmt"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/home"
)

const (
	// default locations
	PreloadFileName = ".murex_preload"
	moduleDirName   = ".murex_modules/"
	ProfileFileName = ".murex_profile"
	historyFileName = ".murex_history"
	fileNameCrop    = len(".murex_")

	// PreloadEnvVar environmental variable name for the preload profile override path. It's stored as a constant so typos are caught by the compiler
	PreloadEnvVar = "MUREX_PRELOAD"

	// ModuleEnvVar environmental variable name for the module override path. It's stored as a constant so typos are caught by the compiler
	ModuleEnvVar = "MUREX_MODULES"

	// ProfileEnvVar environmental variable name for the profile override path. It's stored as a constant so typos are caught by the compiler
	ProfileEnvVar = "MUREX_PROFILE"

	HistoryEnvVar = "MUREX_HISTORY"

	ConfigEnvVar = "MUREX_CONFIG_DIR"
)

var (
	_pathPreload string
	_pathModules string
	_pathProfile string
	_pathHistory string
)

// PreloadPath returns the path of the preload profile
func PreloadPath() string {
	if _pathPreload == "" {
		_pathPreload = PreloadPathTestable()
	}
	return _pathPreload
}

func PreloadPathTestable() string {
	return validateProfilePath(PreloadEnvVar, PreloadFileName, false)
}

// ProfilePath returns the path of your murex profile
func ProfilePath() string {
	if _pathProfile == "" {
		_pathProfile = ProfilePathTestable()
	}
	return _pathProfile
}

func ProfilePathTestable() string {
	return validateProfilePath(ProfileEnvVar, ProfileFileName, false)
}

// HistoryPath returns the path of your shell's history file
func HistoryPath() string {
	if _pathHistory == "" {
		_pathHistory = validateProfilePath(HistoryEnvVar, historyFileName, false)
	}
	return _pathHistory
}

func validateProfilePath(envvar, defaultFileName string, isDir bool) string {
	profilePath := strings.TrimSpace(os.Getenv(envvar))
	configPath := strings.TrimSpace(os.Getenv(ConfigEnvVar))
	return _validateProfilePath(envvar, defaultFileName, profilePath, configPath, isDir)
}

func _validateProfilePath(envvar, defaultFileName, profilePath, configPath string, isDir bool) string {
	if profilePath != "" {
		return _validateProfilePathWithProfileEnvVar(envvar, defaultFileName, profilePath, isDir)
	}

	if configPath != "" {
		return _validateProfilePathWithConfigEnvVar(defaultFileName, configPath)
	}

	return _validateProfilePathWithHome(defaultFileName)
}

func _validateProfilePathWithProfileEnvVar(envvar, defaultFileName, path string, isDir bool) string {
	fi, err := os.Stat(path)
	if err != nil {
		if isDir {
			return _makeDirectory(path, envvar, defaultFileName)
		}
		return path
	}

	if fi.IsDir() && !isDir {
		path += consts.PathSlash + defaultFileName
	} else if isDir {
		path += consts.PathSlash
	}

	return path
}

func _validateProfilePathWithConfigEnvVar(defaultFileName, path string) string {
	pathFile := path + consts.PathSlash + defaultFileName[fileNameCrop:]

	fi, err := os.Stat(path)
	if err != nil {
		return _makeDirectory(pathFile, ConfigEnvVar, defaultFileName)
	}

	if fi.IsDir() {
		return pathFile
	}

	home := _validateProfilePathWithHome(defaultFileName)
	fmt.Fprintf(os.Stderr,
		"Override path specified in %s exists and is not a directory: '%s'!\n  This path cannot be used so defaulting to $HOME: %s\n",
		ConfigEnvVar, path, home)

	return home
}

func _makeDirectory(path, envvar, defaultFileName string) string {
	fmt.Fprintf(os.Stderr,
		"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new file will be created.\n",
		envvar, path)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		home := _validateProfilePathWithHome(defaultFileName)
		fmt.Fprintf(os.Stderr, "!!! ERROR: %s\n!!! This path cannot be used so defaulting to $HOME: %s\n",
			err.Error(), home)
		return home
	}
	return path
}

func _validateProfilePathWithHome(defaultFileName string) string {
	return home.MyDir + consts.PathSlash + defaultFileName
}

// ModulePath returns the install path of the murex modules / packages
func ModulePath() string {
	if _pathModules == "" {
		_pathModules = ModulePathTestable()
	}

	return _pathModules
}

func ModulePathTestable() string {
	path := validateProfilePath(ModuleEnvVar, moduleDirName, true)

	/*fi, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s does not exist: '%s'!\n  Assuming this is intentional, a new directory will be created.\n",
			ModuleEnvVar, path)
		return addTrailingSlash(path)
	}

	if !fi.IsDir() {
		fmt.Fprintf(os.Stderr,
			"Override path specified in %s is a file, directory expected: '%s'!\n  Falling back to default path.\n",
			ModuleEnvVar, path)
		return _validateProfilePathWithHome(moduleDirName)
	}*/

	return addTrailingSlash(path)
}

func addTrailingSlash(path string) string {
	if strings.HasSuffix(path, consts.PathSlash) {
		return path
	}
	return path + consts.PathSlash
}
