package docgen

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

	// Document template for the documents
	DocumentTemplate string `yaml:"DocumentTemplate"`

	// Category template for the category (like an index.html type page)
	CategoryTemplate string `yaml:"CategoryTemplate"`

	docTemplate *template.Template
	catTemplate *template.Template
	ref         *category
	index       int
}

// CategoryPath is the file name and path to write the category index file to
func (t templates) CategoryFilePath() string {
	return t.OutputPath + t.CategoryFile
}

func (t templates) CategoryValues(docs documents) *categoryValues {
	var (
		dv sortableDocumentValues
		dt sortableDocumentDateTime
	)

	for i := range docs {
		if docs[i].CategoryID == t.ref.ID {
			dv = append(dv, t.DocumentValues(&docs[i], docs, true))
			dt = append(dt, t.DocumentValues(&docs[i], docs, true))
		}
	}

	sort.Sort(dv)
	sort.Sort(dt)

	return &categoryValues{
		ID:          t.ref.ID,
		Title:       t.ref.Title,
		Description: t.ref.Description,
		Documents:   dv,
		DateTime:    dt,
	}
}

type categoryValues struct {
	ID          string
	Title       string
	Description string
	Documents   []*documentValues
	DateTime    []*documentValues
}
