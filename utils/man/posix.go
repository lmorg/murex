// +build !windows

package man

import (
	"bufio"
	"compress/gzip"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	rxMatchManSection   *regexp.Regexp = regexp.MustCompile(`/man[1678]/`)
	rxMatchFlagsEscaped *regexp.Regexp = regexp.MustCompile(`\\f[BI]((\\-|-)[a-zA-Z0-9]|(\\-\\-|--)[\\\-a-zA-Z0-9]+).*?\\f[RP]`)
	rxMatchFlagsQuoted  *regexp.Regexp = regexp.MustCompile(`\.IP "(.*?)"`)
	rxMatchFlagsOther   *regexp.Regexp = regexp.MustCompile(`\.B (.*?)\\fR`)
	rxMatchFlagsFlag    *regexp.Regexp = regexp.MustCompile(`(--[\-a-zA-Z0-9]+)`)
	rxReplaceMarkup     *regexp.Regexp = regexp.MustCompile(`\.[a-zA-Z]+(\s|)`)
)

/*
MANUAL SECTIONS (Linux)
    The standard sections of the manual include:

    1      User Commands
    2      System Calls
    3      C Library Functions
    4      Devices and Special Files
    5      File Formats and Conventions
    6      Games et. al.
    7      Miscellanea
    8      System Administration tools and Daemons

    Distributions customize the manual section to their specifics,
    which often include additional sections.
*/
/*
	OpenBSD
	1	General commands (tools and utilities).
	2	System calls and error numbers.
	3	Library functions.
	3p	perl(1) programmer's reference guide.
	4	Device drivers.
	5	File formats.
	6	Games.
	7	Miscellaneous information.
	8	System maintenance and operation commands.
	9	Kernel internals.
*/

// GetManPages executes `man -w` to locate the manual files
func GetManPages(exe string) []string {
	// Get paths
	cmd := exec.Command("man", "-w", exe)
	b, err := cmd.Output()
	if err != nil {
		return nil
	}

	s := strings.TrimSpace(string(b))
	if s == exe {
		return nil
	}

	return strings.Split(s, ":")
}

// ParseFlags runs the parser to locate any flags with hyphen prefixes
func ParseFlags(paths []string) (flags []string) {
	// Parse man pages
	fMap := make(map[string]bool)
	for i := range paths {
		if !rxMatchManSection.MatchString(paths[i]) {
			continue
		}
		parseFlags(&fMap, paths[i])
	}

	for f := range fMap {
		flags = append(flags, f)
	}
	sort.Strings(flags)
	return
}

func parseFlags(flags *map[string]bool, filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return
	}

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		defer gz.Close()
		if err != nil {
			return
		}

		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		s := scanner.Text()

		match := rxMatchFlagsEscaped.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			s := strings.Replace(match[i][1], `\`, "", -1)
			if strings.HasSuffix(s, "fR") || strings.HasSuffix(s, "fP") {
				s = s[:len(s)-2]
			}
			(*flags)[s] = true
		}

		match = rxMatchFlagsQuoted.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			flag := rxMatchFlagsFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = true
			}
		}

		match = rxMatchFlagsOther.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			flag := rxMatchFlagsFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = true
			}
		}
	}

	return
}

// ParseDescription runs the parser to locate a description
func ParseDescription(paths []string) string {
	for i := range paths {
		if !rxMatchManSection.MatchString(paths[i]) {
			continue
		}
		desc := parseDescription(paths[i])
		if desc != "" {
			return desc
		}
	}

	return ""
}

func parseDescription(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return ""
	}

	var scanner *bufio.Scanner

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		defer gz.Close()
		if err != nil {
			return ""
		}

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
			return strings.TrimSpace(desc)
		}

		if read {
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
			desc += s
		}

		if strings.Contains(s, "NAME") {
			read = true
		}
	}

	return ""
}
