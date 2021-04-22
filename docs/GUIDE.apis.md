# _murex_ Shell Docs

## API Reference

This section is a glossary of APIs.

These APIs are provided for reference for any developers wishing to write
their own builtins. However some APIs are still worth being aware of even
when just writing _murex_ scripts because they provide a background into
the internal logic of _murex_'s runtime.

## Pages

* [`Marshal()` (type)](apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`ReadArray()` (type)](apis/ReadArray.md):
  Read from a data type one array element at a time
* [`ReadArrayWithType()` (type)](apis/ReadArrayWithType.md):
  Read from a data type one array element at a time and return the elements contents and data type
* [`ReadIndex()` (type)](apis/ReadIndex.md):
  Data type handler for the index, `[`, builtin
* [`ReadMap()` (type)](apis/ReadMap.md):
  Treat data type as a key/value structure and read its contents
* [`ReadNotIndex()` (type)](apis/ReadNotIndex.md):
  Data type handler for the bang-prefixed index, `![`, builtin
* [`Unmarshal()` (type)](apis/Unmarshal.md):
  Converts a structured file format into structured memory
* [`WriteArray()` (type)](apis/WriteArray.md):
  Write a data type, one array element at a time
* [`lang.ArrayTemplate()` (template API)](apis/lang.ArrayTemplate.md):
  Unmarshals a data type into a Go struct and returns the results as an array
* [`lang.ArrayTemplateWithType()` (template API)](apis/lang.ArrayTemplateWithType.md):
  Unmarshals a data type into a Go struct and returns the results as an array with data type included
* [`lang.IndexTemplateObject()` (template API)](apis/lang.IndexTemplateObject.md):
  Returns element(s) from a data structure
* [`lang.IndexTemplateTable()` (template API)](apis/lang.IndexTemplateTable.md):
  Returns element(s) from a table
* [`lang.MarshalData()` (system API)](apis/lang.MarshalData.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [`lang.UnmarshalData()` (system API)](apis/lang.UnmarshalData.md):
  Converts a _murex_ data-type into structured memory