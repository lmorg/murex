# Prettify JSON

> Prettifies JSON to make it human readable

## Description

Takes JSON from the stdin and reformats it to make it human readable, then
outputs that to stdout.

## Usage

```
<stdin> -> pretty -> <stdout>
```

## Examples

```
» tout json {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} -> pretty 
{
    "Array": [
        1,
        2,
        3
    ],
    "Map": {
        "String": "Foobar",
        "Number": 123.456
    }
}
```

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