package docgen

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Document is the catalogue of config files found in the search path
type document struct {
	// DocumentID is the identifier for the document and name to write the document to disk as (excluding extension)
	DocumentID string `yaml:"DocumentID"`

	// Title of the document
	Title string `yaml:"Title"`

	// CategoryID as per the Categories map (see below)
	CategoryID string `yaml:"CategoryID"`

	// SubCategory is an optional field for grouping documents in a particular category
	SubCategoryIDs []string `yaml:"SubCategoryIDs"`

	// Summary is a one line summary
	Summary string `yaml:"Summary"`

	// Usage - this is more intended as an API reference than example code
	Usage string `yaml:"Usage"`

	// Example code to accompany the usage reference
	Examples string `yaml:"Examples"`

	// Description is the contents of the document
	Description string `yaml:"Description"`

	// Payload is a document describing an APIs payload
	Payload string `yaml:"Payload"`

	// EventReturn is a document describing an events callback values
	EventReturn string `yaml:"EventReturn"`

	// MetaValues is a document describing the meta values supported by a command or API
	MetaValues string `yaml:"MetaValues"`

	// Flags is a map of supported flags
	Flags map[string]string `yaml:"Flags"`

	// Like flags but where parameters are numerically defined
	Parameters []string `yaml:"Parameters"`

	// Associations is for murex data-types
	Associations AssociationValues `yaml:"Associations"`

	// API hooks for murex data-types
	Hooks map[string]string `yaml:"Hooks"`

	// Detail is for misc details
	Detail string `yaml:"Detail"`

	// Synonyms or aliases (if applicable)
	Synonyms []string `yaml:"Synonyms"`

	// Related documents (these should be in the format of `Category/FileName`)
	Related []string `yaml:"Related"`

	// Date article was published
	DateTime string `yaml:"DateTime"`

	// WriteTo is the path to write to, if different from the category path
	WriteTo string `yaml:"WriteTo"`

	// Automatically pulled from file location
	SourcePath string `yaml:"-"`
}

// AssociationValues are associations registered by murex data-types
type AssociationValues struct {
	Extensions []string `yaml:"Extensions"`
	Mimes      []string `yaml:"Mimes"`
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
	path := d.WriteTo
	if path == "" {
		path = t.OutputPath
	}
	return path + t.DocumentFileName(d)
}

const dateTimeParse = `2006-01-02 15:04`

func (t templates) DocumentValues(d *document, docs documents, nest bool) *documentValues {
	var (
		dateTime time.Time
		err      error
	)

	if d.DateTime == "" {
		dateTime = time.Now()
	} else {
		dateTime, err = time.Parse(dateTimeParse, d.DateTime)
		if err != nil {
			panic(fmt.Sprintf("Cannot parse DateTime as `%s` on %s: %s", dateTimeParse, d.DocumentID, err.Error()))
		}
	}

	dv := &documentValues{
		ID:                  d.DocumentID,
		Title:               d.Title,
		FileName:            t.DocumentFileName(d),
		FilePath:            t.DocumentFilePath(d),
		WriteTo:             d.WriteTo,
		SourcePath:          d.SourcePath,
		Hierarchy:           d.Hierarchy(),
		CategoryID:          d.CategoryID,
		CategoryTitle:       t.ref.Title,
		CategoryDescription: t.ref.Description,
		SubCategories:       documentSubCategories(t.ref),
		Summary:             d.Summary,
		Description:         d.Description,
		Usage:               d.Usage,
		Payload:             d.Payload,
		EventReturn:         d.EventReturn,
		MetaValues:          d.MetaValues,
		Examples:            d.Examples,
		Detail:              d.Detail,
		Synonyms:            d.Synonyms,
		Parameters:          d.Parameters,
		Associations:        d.Associations,
		DateTime:            dateTime,
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
			relCatID = "???" + d.CategoryID // any category
			relDocID = val
		}

		dv.Related = append(
			dv.Related,
			t.DocumentValues(docs.ByID(d.DocumentID, relCatID, relDocID), docs, false),
		)
	}

	for flag, desc := range d.Flags {
		dv.Flags = append(dv.Flags, &flagValues{
			Flag:        flag,
			Description: desc,
		})
	}

	for hook, comment := range d.Hooks {
		dv.Hooks = append(dv.Hooks, &hookValues{
			Hook:    hook,
			Comment: comment,
		})
	}

	sort.Sort(dv.Flags)
	sort.Sort(dv.Related)
	sort.Sort(dv.Hooks)
	sort.Strings(dv.Associations.Extensions)
	sort.Strings(dv.Associations.Mimes)

	return dv
}

type documentValuesSubCats struct {
	ID          string
	Title       string
	Description string
}

func documentSubCategories(cat *category) []documentValuesSubCats {
	var subs []documentValuesSubCats
	for _, c := range cat.SubCategories {
		subs = append(subs, documentValuesSubCats{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
		})
	}
	return subs
}

type documentValues struct {
	ID                  string
	Title               string
	FileName            string
	FilePath            string
	WriteTo             string
	SourcePath          string
	Hierarchy           string
	CategoryID          string
	CategoryTitle       string
	CategoryDescription string
	SubCategories       []documentValuesSubCats
	Summary             string
	Description         string
	Payload             string
	EventReturn         string
	MetaValues          string
	Usage               string
	Examples            string
	Flags               sortableFlagValues
	Hooks               sortableHookValues
	Parameters          []string
	Associations        AssociationValues
	Detail              string
	Synonyms            []string
	Related             sortableDocumentValues
	DateTime            time.Time
}

type sortableDocumentValues []*documentValues

func (v sortableDocumentValues) Len() int           { return len(v) }
func (v sortableDocumentValues) Less(i, j int) bool { return v[i].Title < v[j].Title }
func (v sortableDocumentValues) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type flagValues struct {
	Flag        string
	Description string
}

type sortableDocumentDateTime []*documentValues

func (v sortableDocumentDateTime) Len() int           { return len(v) }
func (v sortableDocumentDateTime) Less(i, j int) bool { return v[i].DateTime.After(v[j].DateTime) }
func (v sortableDocumentDateTime) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type sortableFlagValues []*flagValues

func (v sortableFlagValues) Len() int           { return len(v) }
func (v sortableFlagValues) Less(i, j int) bool { return v[i].Flag < v[j].Flag }
func (v sortableFlagValues) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type hookValues struct {
	Hook    string
	Comment string
}

type sortableHookValues []*hookValues

func (v sortableHookValues) Len() int           { return len(v) }
func (v sortableHookValues) Less(i, j int) bool { return v[i].Hook < v[j].Hook }
func (v sortableHookValues) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type documents []document

func (d documents) Len() int           { return len(d) }
func (d documents) Less(i, j int) bool { return d[i].Title < d[j].Title }
func (d documents) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

// ByID returns a document by it's CategoryID and DocumentID
func (d documents) ByID(requesterID, categoryID, documentID string) *document {
	var anyCat bool

	if strings.HasPrefix(categoryID, "???") {
		categoryID = categoryID[3:]
		anyCat = true
	}

	for i := range d {
		if d[i].CategoryID != categoryID {
			continue
		}

		if d[i].DocumentID == documentID {
			return &d[i]
		}
		for syn := range d[i].Synonyms {
			if d[i].Synonyms[syn] == documentID {
				copy := d[i]
				title := strings.Replace(d[i].Title, d[i].DocumentID, d[i].Synonyms[syn], -1)
				if title == d[i].Title {
					title = d[i].Synonyms[syn]
				}
				copy.Title = title
				return &copy
			}
		}
	}

	if anyCat {
		for i := range d {
			if d[i].DocumentID == documentID {
				return &d[i]
			}
			for syn := range d[i].Synonyms {
				if d[i].Synonyms[syn] == documentID {
					copy := d[i]
					title := strings.Replace(d[i].Title, d[i].DocumentID, d[i].Synonyms[syn], -1)
					if title == d[i].Title {
						title = d[i].Synonyms[syn]
					}
					copy.Title = title
					return &copy
				}
			}
		}
	}

	warning(requesterID, fmt.Sprintf("Cannot find document with the ID `%s/%s`", categoryID, documentID))
	return &document{
		Title:      documentID,
		DocumentID: categoryID + "/" + documentID,
		CategoryID: categoryID,
	}
}

// Documents is all of the collated documents pre-rendering
var Documents documents

func (d *document) IsInSubCategory(id string) bool {
	for i := range d.SubCategoryIDs {
		if d.SubCategoryIDs[i] == id {
			return true
		}
	}

	return false
}
