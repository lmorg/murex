package main

import "text/template"

func readTemplate(path string) *template.Template {
	if path == "" {
		return nil
	}

	f := fileReader(path)
	tmpl := string(readAll(f))
	return template.Must(template.New(path).Funcs(funcMap).Parse(tmpl))
}
