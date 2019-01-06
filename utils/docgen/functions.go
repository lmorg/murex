package main

import (
	"errors"
	"strconv"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"quote": funcQuote,
	"html":  funcHTML,
	"md":    funcMarkdown,
	"trim":  strings.TrimSpace,
	"doc":   funcRenderedDocuments,
	"cat":   funcRenderedCategories,
	"file":  funcFile,
}

/************
 * MARKDOWN *
 ************/

func funcMarkdown(s string) string {
	var (
		new          []rune
		backtick     int
		code         bool
		skipNextCrLf bool
	)

	for _, c := range s {
		switch c {
		case '`':
			backtick++
			if backtick == 3 {
				backtick = 0
				code = !code
				skipNextCrLf = true
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
 *  Quote   *
 ************/

func funcQuote(s string) string {
	return strconv.Quote(funcMarkdown(s))
}

/************
 *   HTML   *
 ************/

func funcHTML(s string) string {
	panic("HTML output not yet written")
}

func funcRenderedDocuments(cat, doc string, index int) (string, error) {
	if len(Config.renderedDocuments[cat]) == 0 {
		return "", errors.New("Category does not exist, doesn't have any documents, or hasn't yet been parsed")
	}

	if len(Config.renderedDocuments[cat][doc]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that document")
	}

	return Config.renderedDocuments[cat][doc][index], nil
}

func funcRenderedCategories(cat string, index int) (string, error) {
	if len(Config.renderedCategories[cat]) == 0 {
		return "", errors.New("Category does not exist or hasn't yet been parsed")
	}

	if len(Config.renderedCategories[cat]) <= index {
		return "", errors.New("index greater than number of templates currently parsed for that category")
	}

	return Config.renderedCategories[cat][index], nil
}

func funcFile(path ...string) string {
	f := fileReader(strings.Join(path, ""))
	b := readAll(f)
	return string(b)
}
