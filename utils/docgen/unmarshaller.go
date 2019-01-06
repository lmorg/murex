package main

import (
	yaml "gopkg.in/yaml.v2"
)

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
