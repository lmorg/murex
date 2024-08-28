# struct.json.pretty: `pretty`

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
* `struct.json.pretty`


## See Also

* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [io.out.type: `tout`](../commands/tout.md):
  Print a string to the stdout and set it's data-type
* [io.out: `out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character

<hr/>

This document was generated from [builtins/core/pretty/pretty_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/pretty/pretty_doc.yaml).