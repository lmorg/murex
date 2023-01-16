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

// DumpNotIndex returns an array of compiled builtins supporting deserialization by !index
func DumpNotIndex() (dump []string) {
	for name := range ReadNotIndexes {
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
