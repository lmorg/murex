# Reformat Data type (`format`)

> Reformat one data-type into another data-type

## Description

`format` takes a data from stdin and returns that data reformated in another
specified data-type

## Usage

```
<stdin> -> format data-type -> <stdout>
```

## Examples

```
» tout json { "One": 1, "Two": 2, "Three": 3 } -> format yaml
One: 1
Three: 3
Two: 2
```

## See Also

* [Define Type (`cast`)](../commands/cast.md):
  Alters the data-type of the previous function without altering its output
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [`Marshal()` (type)](../apis/Marshal.md):
  Converts structured memory into a structured file format (eg for stdio)
* [`Unmarshal()` (type)](../apis/Unmarshal.md):
  Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/core/typemgmt/format_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/format_doc.yaml).