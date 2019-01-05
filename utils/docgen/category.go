package main

import (
	"sort"
	"text/template"
)

// category is the groupings for documents
type category struct {
	// Directory name to write the categorized contents to
	// (this is also used with an extension for the category
	// "index" file)
	OutputDirName string `yaml:"OutputDirName"`

	// Name of the category
	Title string `yaml:"Title"`

	// Description of the category
	Description string `yaml:"Description"`

	// Template for the documents
	DocumentTemplate string `yaml:"DocumentTemplate"`

	// Template for the category (like an index.html type page)
	CategoryTemplate string `yaml:"CategoryTemplate"`
}

var (
	documentTemplates = make(map[string]*template.Template)
	categoryTemplates = make(map[string]*template.Template)
)

func (c category) TemplateValues(id string, documents documents) *categoryValues {
	var dc sortableDocumentValues

	for i := range documents {
		if documents[i].CategoryID == id {
			dc = append(dc, documents[i].TemplateValues(documents, true))
		}
	}

	sort.Sort(dc)

	return &categoryValues{
		Title:       c.Title,
		Description: c.Description,
		Documents:   dc,
	}
}

type categoryValues struct {
	Title       string
	Description string
	Documents   []*documentValues
}
