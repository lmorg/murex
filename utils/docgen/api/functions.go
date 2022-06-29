package docgen

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var funcMap = template.FuncMap{
	"quote":      funcQuote,
	"html":       funcHTML,
	"md":         funcMarkdown,
	"trim":       strings.TrimSpace,
	"doc":        funcRenderedDocuments,
	"cat":        funcRenderedCategories,
	"file":       funcFile,
	"include":    funcInclude,
	"notanindex": funcNotAnIndex,
	"date":       funcDate,
	"time":       funcTime,
	"otherdocs":  funcOtherDocs,
}

/************
 * Markdown *
 ************/

// Takes: string (contents as read from YAML in a machine readable subset of markdown)
// Returns: markdown contents cleaned up for printing
func funcMarkdown(s string) string {
	var (
		new          []rune
		backtick     int
		code         bool
		skipNextBt   bool
		skipNextCrLf bool
	)

	for pos, c := range s {
		switch c {
		case '`':
			backtick++
			if backtick == 3 {
				backtick = 0
				switch {
				case skipNextBt:
					new = append(new, '`', '`', '`')
					skipNextBt = false
				case pos != len(s)-1 && s[pos+1] != '\r' && s[pos+1] != '\n':
					new = append(new, '`', '`', '`')
					skipNextBt = true
				default:
					code = !code
					skipNextCrLf = true
				}
			}

		case '\r':
			// strip carridge returns from output (even on Windows)

		case '\n':
			for i := 0; i < backtick; i++ {
				new = append(new, '`')
			}
			backtick = 0
			if !skipNextCrLf {
				new = append(new, c)
			}

			if code {
				new = append(new, ' ', ' ', ' ', ' ')
			}
			skipNextCrLf = false

		default:
			for i := 0; i < backtick; i++ {
				new = append(new, '`')
			}
			backtick = 0
			skipNextCrLf = false
			new = append(new, c)
		}
	}

	if skipNextCrLf {
		new = new[:len(new)-5]
	}

	s = strings.TrimSuffix(string(new), "\n")
	return strings.TrimSuffix(s, "\r")
}

/************
 *   Quote  *
 ************/

// Takes: strings (contents)
// Returns: contents with some characters escaped for printing in source code (eg \")
func funcQuote(s string) string {
	return strconv.Quote(funcMarkdown(s))
}

/************
 *   HTML   *
 ************/

// Takes: string (contents in markdown)
// Returns: HTML rendered contents
func funcHTML(s string) string {
	panic("HTML output not yet written")
}

/************
 *    Doc   *
 ************/

// Takes: string (category name), string (document name), int (index of rendered document template)
// Returns: contents of rendered document template
func funcRenderedDocuments(cat, doc string, index int) (string, error) {
	if len(Config.renderedDocuments[cat]) == 0 {
		return "", errors.New("Category does not exist, doesn't have any documents, or hasn't yet been parsed")
	}

	if len(Config.renderedDocuments[cat][doc]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that document")
	}

	return Config.renderedDocuments[cat][doc][index], nil
}

/************
 *    Cat   *
 ************/

// Takes: string (category name) and int (index of rendered category template)
// Returns: contents of rendered category template
func funcRenderedCategories(cat string, index int) (string, error) {
	if len(Config.renderedCategories[cat]) == 0 {
		return "", errors.New("Category does not exist or hasn't yet been parsed")
	}

	if len(Config.renderedCategories[cat]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that category")
	}

	return Config.renderedCategories[cat][index], nil
}

/************
 *   File   *
 ************/

// Takes: slice of strings (file path)
// Returns: contents of file based on a concatenation of the slice
func funcFile(path ...string) string {
	f := fileReader(strings.Join(path, ""))
	b := readAll(f)
	return string(b)
}

/************
 *  Include *
 ************/

// Takes: string (contents)
// Looks for: {{ include "filename" }}
// Returns: document with search string replaced with contents of filename
func funcInclude(s string) string {
	rx := regexp.MustCompile(`\{\{ include "([-_/.a-zA-Z0-9]+)" \}\}`)
	if !rx.MatchString(s) {
		return s
	}

	match := rx.FindAllStringSubmatch(s, -1)
	for i := range match {
		f := fileReader(match[i][1])
		b := bytes.TrimSpace(readAll(f))
		s = strings.Replace(s, match[i][0], string(b), -1)
	}

	return s
}

/************
 *NotAnIndex*
 ************/

// Takes: number (eg variable)
// Returns: number incremented by 1
func funcNotAnIndex(i int) int {
	return i + 1
}

/************
 *   Date   *
 ************/

// Takes: DateTime (time.Time)
// Returns: Date as string
func funcDate(dt time.Time) string {
	return dt.Format("02.01.2006")
}

/************
 *   Time   *
 ************/

// Takes: DateTime (time.Time)
// Returns: Time as string
func funcTime(dt time.Time) string {
	return dt.Format("15:04")
}

/************
 * OtherDocs*
 ************/

// Takes: string (category ID)
// Returns: list of documents in that category
func funcOtherDocs(id string) (d documents) {
	for i := range Documents {
		if Documents[i].CategoryID == id {
			d = append(d, Documents[i])
		}
	}
	return
}
