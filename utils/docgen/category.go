package main

import (
	"sort"
	"text/template"
)

// category is the groupings for documents
type category struct {
	ID string `yaml:"ID"`

	// Name of the category
	Title string `yaml:"Title"`

	// Description of the category
	Description string `yaml:"Description"`

	Templates []templates `yaml:"Templates"`
}

type templates struct {
	// OutputPath to write the rendered documents
	OutputPath string `yaml:"OutputPath"`

	// CategoryFile is the file name (and path relative to OutputPath) of the
	// category "index" file
	CategoryFile string `yaml:"CategoryFile"`

	// OutputExt is the file extension of the rendered documents
	// (this is not applied to the category file)
	OutputExt string `yaml:"OutputExt"`

	// OutputFormat, applied to the fields before rendering
	//   * markdown (replace ``` with indents)
	//   * quote (markdown + escaped and quoted chars for embedding in JS/Go source)
	//   * html (replace markdown with HTML tags)
	//   * any other value (no format conversion)
	OutputFormat string `yaml:"OutputFormat"`

	// Document template for the documents
	DocumentTemplate string `yaml:"DocumentTemplate"`

	// Category template for the category (like an index.html type page)
	CategoryTemplate string `yaml:"CategoryTemplate"`

	docTemplate *template.Template
	catTemplate *template.Template
	ref         *category
}

// CategoryPath is the file name and path to write the category index file to
func (t templates) CategoryFilePath() string {
	return t.OutputPath + t.CategoryFile
}

func (t templates) CategoryValues(docs documents) *categoryValues {
	var dv sortableDocumentValues

	for i := range docs {
		if docs[i].CategoryID == t.ref.ID {
			dv = append(dv, t.DocumentValues(&docs[i], docs, true))
		}
	}

	sort.Sort(dv)

	return &categoryValues{
		Title:       t.ref.Title,
		Description: t.ref.Description,
		Documents:   dv,
	}
}

type categoryValues struct {
	Title       string
	Description string
	Documents   []*documentValues
}
