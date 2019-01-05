package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func fileReader(path string) *os.File {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func readAll(f *os.File) []byte {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	return b
}

func unmarshal(b []byte, v interface{}) {
	err := yaml.UnmarshalStrict(b, v)
	if err != nil {
		panic(err.Error())
	}
}

func parseSourceFile(path string, structure interface{}) {
	f := fileReader(path)
	b := readAll(f)
	unmarshal(b, structure)
}

func structuredMessage(message string, v interface{}) string {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err.Error())
	}

	return message + "\n" + string(b)
}
