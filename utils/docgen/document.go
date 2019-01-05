package main

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

// Path is the path to write documents to
func (d document) Path() string {
	file := d.DocumentID + Config.OutputExt
	path := Config.OutputPath + Config.Categories[d.CategoryID].OutputDirName
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return path + file
}

func (d document) TemplateValues(documents documents, nest bool) *documentValues {
	t := &documentValues{
		Title:         d.Title,
		Path:          Config.Categories[d.CategoryID].OutputDirName + "/" + d.DocumentID + Config.OutputExt,
		CategoryID:    d.CategoryID,
		CategoryTitle: Config.Categories[d.CategoryID].Title,
		Summary:       d.Summary,
		Description:   d.Description,
		Usage:         d.Usage,
		Examples:      d.Examples,
		Detail:        d.Detail,
		Synonyms:      d.Synonyms,
	}

	if !nest {
		return t
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

		t.Related = append(t.Related, documents.ByID(relCatID, relDocID).TemplateValues(documents, false))
		//sort.Sort(t.Related)
	}

	for flag, desc := range d.Flags {
		t.Flags = append(t.Flags, &flagValues{
			Flag:        flag,
			Description: desc,
		})
	}

	sort.Sort(t.Flags)
	sort.Sort(t.Related)

	return t
}

type documentValues struct {
	Title         string
	Path          string
	CategoryID    string
	CategoryTitle string
	Summary       string
	Description   string
	Usage         string
	Examples      string
	Flags         sortableFlagValues
	Detail        string
	Synonyms      []string
	Related       sortableDocumentValues
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

// GetTitle returns the title of a document from it's CategoryID and DocumentID
func (d documents) GetTitle(categoryID, documentID string) string {
	for i := range d {
		if d[i].DocumentID == documentID && d[i].CategoryID == categoryID {
			return d[i].Title
		}
	}

	warning(fmt.Sprintf("Cannot find document with the ID `%s/%s`", categoryID, documentID))
	return documentID
}

// GetTitle returns the title of a document from it's CategoryID and DocumentID
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
