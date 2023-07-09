//go:build !windows
// +build !windows

package man

import (
	"bufio"
	"compress/gzip"
	"os"
	"strings"
)

// ParseSummary runs the parser to locate a summary
func ParseSummary(paths []string) string {
	for i := range paths {
		if invalidMan(paths[i]) {
			continue
		}
		desc := SummaryCache.Get(paths[i])
		if desc != "" {
			return desc
		}
		desc = parseSummary(paths[i])
		if desc != "" {
			SummaryCache.Set(paths[i], desc)
			return desc
		}
	}

	return ""
}

func parseSummary(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return ""
		}
		defer gz.Close()

		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(file)
	}

	var (
		read bool
		desc string
	)

	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "SYNOPSIS") {
			if len(desc) > 0 && desc[len(desc)-1] == '-' {
				desc = desc[:len(desc)-1]
			}
			return strings.TrimSpace(desc)
		}

		if read {
			// Tidy up man pages generated from reStructuredText
			if strings.HasPrefix(s, `\\n[rst2man-indent`) ||
				strings.HasPrefix(s, `\\$1 \\n`) ||
				strings.HasPrefix(s, `level \\n`) ||
				strings.HasPrefix(s, `level margin: \\n`) {
				continue
			}

			s = strings.Replace(s, ".Nd ", " - ", -1)
			s = strings.Replace(s, "\\(em ", " - ", -1)
			s = strings.Replace(s, " , ", ", ", -1)
			s = strings.Replace(s, "\\fB", "", -1)
			s = strings.Replace(s, "\\fR", "", -1)
			if strings.HasSuffix(s, " ,") {
				s = s[:len(s)-2] + ", "
			}
			s = rxReplaceMarkup.ReplaceAllString(s, "")
			s = strings.Replace(s, "\\", "", -1)

			if strings.HasPrefix(s, `.`) {
				continue
			}

			desc += s
		}

		if strings.Contains(s, "NAME") {
			read = true
		}
	}

	return ""
}
