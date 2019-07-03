package escape

import "strings"

// CommandLine takes in command line arguments as a slice and escapes the parameters
func CommandLine(s []string) {
	for i := range s {
		s[i] = strings.ReplaceAll(s[i], `\`, `\\`)
		s[i] = strings.ReplaceAll(s[i], `$`, `\$`)
		s[i] = strings.ReplaceAll(s[i], `@`, `\@`)
		s[i] = strings.ReplaceAll(s[i], `|`, `\|`)
		s[i] = strings.ReplaceAll(s[i], `?`, `\?`)
		s[i] = strings.ReplaceAll(s[i], `'`, `\'`)
		s[i] = strings.ReplaceAll(s[i], `"`, `\"`)
		s[i] = strings.ReplaceAll(s[i], `(`, `\(`)
		s[i] = strings.ReplaceAll(s[i], `)`, `\)`)
		s[i] = strings.ReplaceAll(s[i], `<`, `\<`)
		s[i] = strings.ReplaceAll(s[i], `>`, `\>`)
		s[i] = strings.ReplaceAll(s[i], `#`, `\#`)
		s[i] = strings.ReplaceAll(s[i], ` `, `\ `)
		s[i] = strings.ReplaceAll(s[i], "\t", `\t`)
		s[i] = strings.ReplaceAll(s[i], "\r", `\r`)
		s[i] = strings.ReplaceAll(s[i], "\n", `\n`)
	}
}
