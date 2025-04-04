package main

import (
	"compress/gzip"
	"io"
	"os"
)

func diskSource(filename string) ([]byte, error) {
	var b []byte

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			return nil, err
		}
		b, err = io.ReadAll(gz)

		file.Close()
		gz.Close()

		if err != nil {
			return nil, err
		}

	} else {
		b, err = io.ReadAll(file)
		file.Close()
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
