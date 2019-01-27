package modules

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/which"
)

func gitSupported() error {
	if which.Which("git") == "" {
		return errors.New("`git` was not found in your $PATH")
	}

	return nil
}

func gitUpdate(p *lang.Process, _ *packageDb) error {
	if err := gitSupported(); err != nil {
		return err
	}

	return gitExec(p, "pull")
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
	if err := gitSupported(); err != nil {
		return "", err
	}

	path, err := gitGetPath(uri)
	if err != nil {
		return "", err
	}

	err = gitExec(p, "clone", uri)
	if err != nil {
		return "", err
	}

	return path, nil
}

func gitExec(p *lang.Process, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = p.Stderr
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
