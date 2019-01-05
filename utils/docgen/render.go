package main

import (
	"fmt"
	"os"
	"text/template"
)

func fileWriter(path string) *os.File {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func renderAll() {
	for i := range Documents {
		renderDocument(documentTemplates[Documents[i].CategoryID], &Documents[i])
	}

	for cat := range Config.Categories {
		renderCategory(categoryTemplates[cat], cat, Config.Categories[cat])
	}
}

func renderDocument(t *template.Template, d *document) {
	if t == nil {
		if Config.Categories[d.CategoryID].DocumentTemplate == "" {
			panic(fmt.Sprintf("no document template for CategoryID `%s`", d.CategoryID))
		}
		panic(structuredMessage("nil template for unknown reasons (this is likely a software bug).", *d))
	}

	f := fileWriter(d.Path())
	log("Rendering document", d.DocumentID)
	err := t.Execute(f, d.TemplateValues(Documents, true))
	if err != nil {
		panic(err.Error())
	}
}

func renderCategory(t *template.Template, id string, c category) {
	f := fileWriter(Config.OutputPath + id + Config.OutputExt)
	log("Rendering category", id)
	err := t.Execute(f, c.TemplateValues(id, Documents))
	if err != nil {
		panic(err.Error())
	}
}
