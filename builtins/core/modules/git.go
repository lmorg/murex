package modules

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"

	profilepaths "github.com/lmorg/murex/config/profile/paths"
	"github.com/lmorg/murex/utils/ansi"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/which"
)

func gitSupported() error {
	if which.Which("git") == "" {
		return errors.New("`git` was not found in your $PATH")
	}

	return nil
}

func gitUpdate(p *lang.Process, pack *packageDb) error {
	if err := gitSupported(); err != nil {
		return err
	}

	return gitExec(p, "-C", profilepaths.ModulePath()+pack.Package, "pull", "-q")
}

func gitStatus(p *lang.Process, pack *packageDb) error {
	if err := gitSupported(); err != nil {
		return err
	}

	// git fetch

	params := []string{"-C", profilepaths.ModulePath() + pack.Package, "fetch"}

	if ansi.IsAllowed() {
		params = append([]string{"-c", "color.status=always"}, params...)
	}

	if err := gitExec(p, params...); err != nil {
		return err
	}

	// git status

	params = []string{"-C", profilepaths.ModulePath() + pack.Package, "status", "-sb"}

	if ansi.IsAllowed() {
		params = append([]string{"-c", "color.status=always"}, params...)
	}

	return gitExec(p, params...)
}

var (
	gitRxPath            = regexp.MustCompile(`^.*/(.*?)(\.git)?$`)
	errCannotParseGitURI = errors.New("cannot parse git URI to extract clone destination")
)

func gitGetPath(uri string) (string, error) {
	match := gitRxPath.FindAllStringSubmatch(uri, -1)
	if len(match) == 0 || len(match[0]) < 2 {
		return "", errCannotParseGitURI
	}

	if match[0][1] == "" {
		return "", errCannotParseGitURI
	}

	return match[0][1], nil
}

func gitGet(p *lang.Process, uri string) (string, error) {
	modulePath := profilepaths.ModulePath()

	if err := gitSupported(); err != nil {
		return "", err
	}

	path, err := gitGetPath(uri)
	if err != nil {
		return "", err
	}

	err = gitExec(p, "clone", uri, modulePath+path)
	if err != nil {
		return "", err
	}

	return mvPackagePath(modulePath + path)
}

func gitExec(p *lang.Process, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = p.Stdout
	cmd.Stderr = p.Stderr

	if err := cmd.Start(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		if !strings.HasPrefix(err.Error(), "signal:") {
			return err
		}
	}

	return nil
}
