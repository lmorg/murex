package main

import (
	"os"
	"path/filepath"
	"strings"
)

func walkSourcePath(path string) {
	err := filepath.Walk(path, walkCallback)
	if err != nil {
		panic(err.Error())
	}
}

func walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		panic(err.Error())
	}

	// We are not interested in anything that isn't a source file
	if !strings.HasSuffix(f.Name(), Config.SourceExt) {
		return nil
	}

	log("Reading", path)

	var src []document
	parseSourceFile(path, &src)
	formatDocuments(src)
	Documents = append(Documents, src...)

	return nil
}
