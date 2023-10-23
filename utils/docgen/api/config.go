package docgen

import (
	"fmt"
	"strings"
)

type config struct {
	// Path to scan for documentation to render
	SourcePath string `yaml:"SourcePath"`

	// File extension of the source documents
	// (anything with this extension will be read by docgen)
	SourceExt string `yaml:"SourceExt"`

	// Categories, templates, etc
	Categories categories `yaml:"Categories"`

	renderedCategories map[string][]string
	renderedDocuments  map[string]map[string][]string
}

// Config is the global configuration for docgen
var Config = new(config)

type categories []category

func (c *categories) ByID(catID string) *category {
	for i := range *c {
		if (*c)[i].ID == catID {
			return &(*c)[i]
		}
	}
	panic(fmt.Sprintf("no category found with ID '%s'", catID))
}

// ReadConfig loads a config file from disk
func ReadConfig(path string) (err error) {
	defer func() {
		// Write a stack trace on error
		if !Panic {
			if r := recover(); r != nil {
				err = fmt.Errorf("%s", r)
			}
		}
	}()

	parseSourceFile(path, Config)

	for cat := range Config.Categories {
		for i := range Config.Categories[cat].Templates {
			updateConfig(&Config.Categories[cat].Templates[i], cat, i)
		}
	}

	Config.renderedCategories = make(map[string][]string)
	Config.renderedDocuments = make(map[string]map[string][]string)

	return
}

func updateConfig(t *templates, cat int, i int) {
	t.ref = &Config.Categories[cat]
	t.index = i

	if !strings.HasSuffix(t.OutputPath, "/") {
		t.OutputPath += "/"
	}

	t.docTemplate = readTemplate(t.DocumentTemplate)
	t.catTemplate = readTemplate(t.CategoryTemplate)
}
