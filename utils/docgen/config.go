package main

import "strings"

type config struct {
	// Path to scan for documentation to render
	SourcePath string `yaml:"SourcePath"`

	// File extension of the source documents
	// (anything with this extension will be read by docgen)
	SourceExt string `yaml:"SourceExt"`

	// Path to write the rendered documents
	OutputPath string `yaml:"OutputPath"`

	// File extension of the rendered documents
	OutputExt string `yaml:"OutputExt"`

	// Categories
	Categories map[string]category `yaml:"Categories"`

	// OutputFormat, applied to the fields before rendering
	//   * markdown (replace ``` with indents)
	//   * html (replace markdown with HTML tags)
	//   * any other value (no format conversion)
	OutputFormat string `yaml:"OutputFormat"`

	// PostProcessing, to the entire document after it has been rendered
	//   * gzip (base64 gzipped archive)
	//   * base64 (converted to base64 with no additional compression)
	//   * thing1+thing2 (chain thing1 (eg markdown) wth thing2 (eg gzip))
	//   * any other value (no post processing)
	PostProcessing string `yaml:"PostProcessing"`
}

// Config is the global configuration for docgen
var Config = new(config)

func readConfig(path string) {
	parseSourceFile(path, Config)
	if !strings.HasSuffix(Config.OutputPath, "/") {
		Config.OutputPath += "/"
	}

	formatCategories()
}
