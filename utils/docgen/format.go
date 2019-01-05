package main

import "strings"

func formatDocuments(d []document) {
	var fn func(*string)

	switch Config.OutputFormat {
	case "markdown":
		fn = formatMarkdown

	case "html":
		fn = formatHTML

	default:
		return
	}

	for i := range d {
		fn(&d[i].Examples)
		fn(&d[i].Description)
		fn(&d[i].Detail)
		fn(&d[i].Summary)
		fn(&d[i].Title)
		d[i].Title = strings.TrimSpace(d[i].Title)
		fn(&d[i].Usage)
	}
}

func formatCategories() {
	var fn func(*string)

	format := func(s string) string {
		fn(&s)
		return s
	}

	switch Config.OutputFormat {
	case "markdown":
		fn = formatMarkdown

	case "html":
		fn = formatHTML

	default:
		return
	}

	for name := range Config.Categories {
		Config.Categories[name] = category{
			CategoryTemplate: Config.Categories[name].CategoryTemplate,
			DocumentTemplate: Config.Categories[name].DocumentTemplate,
			OutputDirName:    Config.Categories[name].OutputDirName,

			Title:       format(Config.Categories[name].Title),
			Description: format(Config.Categories[name].Description),
		}
	}
}

/************
 * MARKDOWN *
 ************/

func formatMarkdown(s *string) {
	var (
		new          []rune
		backtick     int
		code         bool
		skipNextCrLf bool
	)

	for _, c := range *s {
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

	*s = string(new)
}

/************
 *   HTML   *
 ************/

func formatHTML(s *string) {
	panic("HTML output not yet written")
}
