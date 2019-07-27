package docgen

import (
	"fmt"
	"sort"
	"strings"
)

// Document is the catalogue of config files found in the search path
type document struct {
	// DocumentID is the identifier for the document and name to write the document to disk as (excluding extension)
	DocumentID string `yaml:"DocumentID"`

	// Title of the document
	Title string `yaml:"Title"`

	// CategoryID as per the Categories map (see below)
	CategoryID string `yaml:"CategoryID"`

	// Summary is a one line summary
	Summary string `yaml:"Summary"`

	// Usage - this is more intended as an API reference than example code
	Usage string `yaml:"Usage"`

	// Example code to accompany the usage reference
	Examples string `yaml:"Examples"`

	// Description is the contents of the document
	Description string `yaml:"Description"`

	// Flags is a map of supported flags
	Flags map[string]string `yaml:"Flags"`

	// Detail is for misc details
	Detail string `yaml:"Detail"`

	// Synonyms or aliases (if applicable)
	Synonyms []string `yaml:"Synonyms"`

	// Related documents (these should be in the format of `Category/FileName`)
	Related []string `yaml:"Related"`
}

// Hierarchy is the ID path
func (d document) Hierarchy() string {
	if strings.HasPrefix(d.DocumentID, d.CategoryID+"/") {
		return d.DocumentID
	}
	return d.CategoryID + "/" + d.DocumentID
}

// DocumentFileName is the file name of the written documents
func (t templates) DocumentFileName(d *document) string {
	return d.DocumentID + t.OutputExt
}

// DocumentFilePath is the file name and path to write documents to
func (t templates) DocumentFilePath(d *document) string {
	return t.OutputPath + t.DocumentFileName(d)
}

func (t templates) DocumentValues(d *document, docs documents, nest bool) *documentValues {
	dv := &documentValues{
		ID:                  d.DocumentID,
		Title:               d.Title,
		FileName:            t.DocumentFileName(d),
		FilePath:            t.DocumentFilePath(d),
		Hierarchy:           d.Hierarchy(),
		CategoryID:          d.CategoryID,
		CategoryTitle:       t.ref.Title,
		CategoryDescription: t.ref.Description,
		Summary:             d.Summary,
		Description:         d.Description,
		Usage:               d.Usage,
		Examples:            d.Examples,
		Detail:              d.Detail,
		Synonyms:            d.Synonyms,
	}

	if !nest {
		return dv
	}

	for _, val := range d.Related {
		var relCatID, relDocID string

		if strings.Contains(val, "/") {
			split := strings.Split(val, "/")
			if len(split) != 2 {
				panic("related value contains multiple slashes")
			}

			relCatID = split[0]
			relDocID = split[1]

		} else {
			relCatID = d.CategoryID
			relDocID = val
		}

		dv.Related = append(
			dv.Related,
			t.DocumentValues(docs.ByID(relCatID, relDocID), docs, false),
		)
	}

	for flag, desc := range d.Flags {
		dv.Flags = append(dv.Flags, &flagValues{
			Flag:        flag,
			Description: desc,
		})
	}

	sort.Sort(dv.Flags)
	sort.Sort(dv.Related)

	return dv
}

type documentValues struct {
	ID                  string
	Title               string
	FileName            string
	FilePath            string
	Hierarchy           string
	CategoryID          string
	CategoryTitle       string
	CategoryDescription string
	Summary             string
	Description         string
	Usage               string
	Examples            string
	Flags               sortableFlagValues
	Detail              string
	Synonyms            []string
	Related             sortableDocumentValues
}

type sortableDocumentValues []*documentValues

func (v sortableDocumentValues) Len() int           { return len(v) }
func (v sortableDocumentValues) Less(i, j int) bool { return v[i].Title < v[j].Title }
func (v sortableDocumentValues) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type flagValues struct {
	Flag        string
	Description string
}

type sortableFlagValues []*flagValues

func (v sortableFlagValues) Len() int           { return len(v) }
func (v sortableFlagValues) Less(i, j int) bool { return v[i].Flag < v[j].Flag }
func (v sortableFlagValues) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type documents []document

// ByID returns the from it's CategoryID and DocumentID
func (d documents) ByID(categoryID, documentID string) *document {
	for i := range d {
		if d[i].DocumentID == documentID && d[i].CategoryID == categoryID {
			return &d[i]
		}
	}

	warning(fmt.Sprintf("Cannot find document with the ID `%s/%s`", categoryID, documentID))
	return &document{
		Title:      documentID,
		DocumentID: categoryID + "/" + documentID,
		CategoryID: categoryID,
	}
}

// Documents is all of the collated documents pre-rendering
var Documents documents
