package mail

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"regexp"

	"github.com/lmorg/murex/app"
)

var rxGetDomain = regexp.MustCompile(`@([\.\-a-zA-Z0-9]+)$`)

func getDomain(email string) (string, error) {
	match := rxGetDomain.FindAllStringSubmatch(email, 1)

	if len(match) != 1 {
		return "", fmt.Errorf("Could not extract domain from %s: Parsed as %v", email, match)
	}

	if len(match[0]) != 2 {
		return "", fmt.Errorf("Could not extract domain from %s: Parsed as %v", email, match)
	}

	return match[0][1], nil
}

func hostname() string {
	s, err := os.Hostname()
	if err != nil {
		return "localhost"
	}

	return s
}

func username() string {
	u, err := user.Current()
	if err != nil {
		return app.Name
	}

	return u.Username
}

func setSubject(w io.Writer, subject string) error {
	_, err := fmt.Fprintf(w, "Subject: %s\nMIME-version: 1.0;\nContent-Type: text/text\n\n", subject)
	if err != nil {
		return fmt.Errorf("Unable to write Subject header: %s", err.Error())
	}
	return nil
}
