# Prettify JSON

> Prettifies data documents to make it human readable

## Description

Takes JSON or XML from the stdin and reformats it to make it human readable, then
outputs that to stdout.

## Usage

```
<stdin> -> [ --strict | --type (XML|JSON) ] -> <stdout>
```

## Examples

```
» %{Array:[1,2,3],Map:{String:Foobar,Number:123.456}} -> pretty 
{
    "Array": [
        1,
        2,
        3
    ],
    "Map": {
        "Number": 123.456,
        "String": "Foobar"
    }
}
```

## Flags

* `--strict`
    If data type doesn't have a pretty parser, then just output stdin (default behaviour is to try every parser until one works)
* `--type`
    Specify a pretty parser (supported values: "json", "xml")

## Synonyms

* `pretty`


## See Also

* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Output With Type Annotation (`tout`)](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [Reformat Data type (`format`)](../commands/format.md):
  Reformat one data-type into another data-type

<hr/>

This document was generated from [builtins/core/pretty/pretty_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/pretty/pretty_doc.yaml).