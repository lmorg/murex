package escape

import "strings"

// CommandLine takes in command line arguments as a slice and escapes the parameters
func CommandLine(s []string) {
	for i := range s {
		s[i] = strings.Replace(s[i], `\`, `\\`, -1)
		s[i] = strings.Replace(s[i], `$`, `\$`, -1)
		s[i] = strings.Replace(s[i], `@`, `\@`, -1)
		s[i] = strings.Replace(s[i], `|`, `\|`, -1)
		s[i] = strings.Replace(s[i], `?`, `\?`, -1)
		s[i] = strings.Replace(s[i], `'`, `\'`, -1)
		s[i] = strings.Replace(s[i], `"`, `\"`, -1)
		s[i] = strings.Replace(s[i], `(`, `\(`, -1)
		s[i] = strings.Replace(s[i], `)`, `\)`, -1)
		s[i] = strings.Replace(s[i], `<`, `\<`, -1)
		s[i] = strings.Replace(s[i], `>`, `\>`, -1)
		s[i] = strings.Replace(s[i], `#`, `\#`, -1)
		s[i] = strings.Replace(s[i], ` `, `\ `, -1)
		s[i] = strings.Replace(s[i], "\t", `\t`, -1)
		s[i] = strings.Replace(s[i], "\r", `\r`, -1)
		s[i] = strings.Replace(s[i], "\n", `\n`, -1)
	}
}

// Table takes in terminal-rendered tables cells and escapes the contents
func Table(s []string) {
	for i := range s {
		s[i] = strings.Replace(s[i], `\`, `\\`, -1)
		s[i] = strings.Replace(s[i], `$`, `\$`, -1)
		s[i] = strings.Replace(s[i], `@`, `\@`, -1)
		s[i] = strings.Replace(s[i], `"`, `\"`, -1)
		s[i] = strings.Replace(s[i], `<`, `\<`, -1)
		s[i] = strings.Replace(s[i], `>`, `\>`, -1)
		s[i] = strings.Replace(s[i], "\t", `\t`, -1)
		s[i] = strings.Replace(s[i], "\r", `\r`, -1)
		s[i] = strings.Replace(s[i], "\n", `\n`, -1)
	}
}
