package main

import "text/template"

func readCategoryTemplates() {
	for catID := range Config.Categories {
		documentTemplates[catID] = readTemplate(Config.Categories[catID].DocumentTemplate)
		categoryTemplates[catID] = readTemplate(Config.Categories[catID].CategoryTemplate)
		makePath(Config.OutputPath + Config.Categories[catID].OutputDirName)
	}
}

func readTemplate(path string) *template.Template {
	f := fileReader(path)
	tmpl := string(readAll(f))
	return template.Must(template.New(path).Parse(tmpl))
}
