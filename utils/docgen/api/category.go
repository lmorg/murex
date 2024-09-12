package docgen

import (
	"fmt"
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

	VueIcon string `yaml:"VueIcon"`

	SubCategories []*category `yaml:"SubCategories"`

	Templates []templates `yaml:"Templates"`
}

func (c *category) SubCategoryByID(id string) (*category, error) {
	for i := range c.SubCategories {
		if c.SubCategories[i].ID == id {
			return c.SubCategories[i], nil
		}
	}

	return nil, fmt.Errorf("cannot find a sub-category with the id '%s'", id)
}

func (c *category) getSubCategoryTitle(id string) string {
	sub, err := c.SubCategoryByID(id)
	if err != nil {
		return ""
	}

	return sub.Title
}

func (c *category) getSubCategoryDescription(id string) string {
	sub, err := c.SubCategoryByID(id)
	if err != nil {
		return ""
	}

	return sub.Description
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
		ID:            t.ref.ID,
		Title:         t.ref.Title,
		Description:   t.ref.Description,
		Documents:     dv,
		DateTime:      dt,
		SubCategories: t.SubCategoryValues(docs, t.ref),
		UncatDocs:     t.UncategorisedValues(docs, t.ref),
	}
}

func (t templates) UncategorisedValues(docs documents, cat *category) []*documentValues {
	var uncat []*documentValues

	for i := range docs {
		if docs[i].CategoryID != cat.ID || len(docs[i].SubCategoryIDs) > 0 {
			continue
		}
		uncat = append(uncat, t.DocumentValues(&docs[i], docs, true))
	}

	return uncat
}

func (t templates) SubCategoryValues(docs documents, cat *category) []*categoryValues {
	var subs []*categoryValues

	for i := range cat.SubCategories {
		subs = append(subs, t.subCategoryValues(docs, cat.SubCategories[i]))
	}

	return subs
}

func (t templates) subCategoryValues(docs documents, cat *category) *categoryValues {
	var (
		dv sortableDocumentValues
		dt sortableDocumentDateTime
	)

	for i := range docs {
		//if docs[i].SubCategoryID == cat.ID {
		if docs[i].IsInSubCategory(cat.ID) {
			dv = append(dv, t.DocumentValues(&docs[i], docs, true))
			dt = append(dt, t.DocumentValues(&docs[i], docs, true))
		}
	}

	sort.Sort(dv)
	sort.Sort(dt)

	return &categoryValues{
		ID:            cat.ID,
		Title:         cat.Title,
		Description:   cat.Description,
		Documents:     dv,
		DateTime:      dt,
		SubCategories: t.SubCategoryValues(docs, cat),
	}
}

type categoryValues struct {
	ID            string
	Title         string
	Description   string
	Documents     []*documentValues
	DateTime      []*documentValues
	SubCategories []*categoryValues
	UncatDocs     []*documentValues
}
