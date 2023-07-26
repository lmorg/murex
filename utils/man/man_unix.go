//go:build !windows
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

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/rmbs"
)

const errPrefix = "error parsing man page: "

var (
	rxMatchManSection   = regexp.MustCompile(`/man[1678]/`)
	rxMatchFlagsEscaped = regexp.MustCompile(`\\f[BI]((\\-|-)[a-zA-Z0-9]|(\\-\\-|--)[\\\-a-zA-Z0-9]+).*?\\f[RP]`)
	rxMatchFlagsQuoted  = regexp.MustCompile(`\.IP "(.*?)"`)
	rxMatchFlagsDarwin  = regexp.MustCompile(`\.It Fl ([a-zA-Z0-9])`)
	rxMatchFlagsOther   = regexp.MustCompile(`\.B (.*?)`)
	//rxMatchFlagsOther   = regexp.MustCompile(`\.B (.*?)\\fR`)
	rxMatchFlagsNoFmt = regexp.MustCompile(`(--[\-a-zA-Z0-9]+)=([_\-a-zA-Z0-9]+)`)

	rxMatchGetFlag = regexp.MustCompile(`(--[\-a-zA-Z0-9]+)`)

	rxReplaceMarkup = regexp.MustCompile(`\.[a-zA-Z]+(\s|)`)
)

// GetManPages executes `man -w` to locate the manual files
func GetManPages(exe string) []string {
	paths := Paths.Get(exe)
	if paths != nil {
		return paths
	}

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

	paths = strings.Split(s, ":")
	Paths.Set(exe, paths)
	return paths
}

// ParseByPaths runs the parser to locate any flags with hyphen prefixes
func ParseByPaths(command string, paths []string) ([]string, map[string]string) {
	f := Flags.Get(command)
	if f != nil {
		return f.Flags, f.Descriptions
	}

	fMap := make(map[string]string)

	for i := range paths {
		if !rxMatchManSection.MatchString(paths[i]) {
			continue
		}

		scanner, closer, err := createScanner(paths[i])
		switch {
		case err != nil:
			return []string{errPrefix + err.Error()}, map[string]string{}
		case scanner == nil:
			return []string{errPrefix + "scanner is undefined"}, map[string]string{}
		default:
			parseFlags(&fMap, scanner)
			closer()
		}
	}

	parseDescriptions(command, &fMap)

	flags := make([]string, len(fMap))
	var i int
	for f := range fMap {
		flags[i] = f
		i++
	}
	sort.Strings(flags)

	Flags.Set(command, flags, fMap)
	return flags, fMap
}

func createScanner(filename string) (*bufio.Scanner, func() error, error) {
	var scanner *bufio.Scanner

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}

	closer := file.Close

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return nil, closer, err
		}

		closer = func() error {
			gz.Close()
			file.Close()
			return nil
		}
		scanner = bufio.NewScanner(gz)

	} else {
		scanner = bufio.NewScanner(file)
	}

	return scanner, closer, err
}

// ParseByStdio runs the parser to locate any flags with hyphen prefixes
func ParseByStdio(io stdio.Io) ([]string, map[string]string) {
	//scanner := bufio.NewScanner(io)

	fMap := make(map[string]string)
	//parseFlags(&fMap, scanner)

	parseDescriptionsLines(io, &fMap)

	flags := make([]string, len(fMap))
	var i int
	for f := range fMap {
		flags[i] = f
		i++
	}
	sort.Strings(flags)

	return flags, fMap
}

func parseFlags(flags *map[string]string, scanner *bufio.Scanner) {
	for scanner.Scan() {
		s := rmbs.Remove(scanner.Text())

		match := rxMatchFlagsEscaped.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			s := strings.Replace(match[i][1], `\`, "", -1)
			if strings.HasSuffix(s, "fR") || strings.HasSuffix(s, "fP") {
				s = s[:len(s)-2]
			}
			(*flags)[s] = ""
		}

		match = rxMatchFlagsQuoted.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			flag := rxMatchGetFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = ""
			}
		}

		match = rxMatchFlagsDarwin.FindAllStringSubmatch(s, -1) // eg `cat` on OSX
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			(*flags)["-"+match[i][1]] = ""
		}

		match = rxMatchFlagsOther.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) == 0 {
				continue
			}

			//// Fix \^ seen on some OSX man pages
			//match[i][1] = strings.Replace(match[i][1], `\^`, "", -1)

			flag := rxMatchGetFlag.FindAllStringSubmatch(match[i][1], -1)
			for j := range flag {
				if len(flag[j]) == 0 {
					continue
				}

				(*flags)[flag[j][1]] = ""
			}
		}

		match = rxMatchFlagsNoFmt.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) != 3 {
				continue
			}

			(*flags)[match[i][1]] = ""
		}

		match = rxMatchGetFlag.FindAllStringSubmatch(s, -1)
		for i := range match {
			if len(match[i]) != 2 {
				continue
			}
			if strings.HasPrefix(match[i][1], "---") {
				continue
			}

			(*flags)[match[i][1]] = ""
		}
	}

	if scanner.Err() != nil {
		panic(errPrefix + scanner.Err().Error())
	}
}
