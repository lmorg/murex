package lang

import "sort"

// DumpIndex returns an array of compiled builtins supporting deserialization by index
func DumpIndex() (dump []string) {
	for name := range ReadIndexes {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpUnmarshaller returns an array of compiled builtins supporting unmarshalling
func DumpUnmarshaller() (dump []string) {
	for name := range Unmarshallers {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

// DumpMarshaller returns an array of compiled builtins supporting marshalling
func DumpMarshaller() (dump []string) {
	for name := range Marshallers {
		dump = append(dump, name)
	}
	sort.Strings(dump)
	return
}

/*// DumpMime returns a map of MIME-types and their associated murex data type
func DumpMime() map[string]string {
	return mimes
}

// DumpFileExts returns a map of file extensions and their associated murex data type
func DumpFileExts() map[string]string {
	return fileExts
}*/
