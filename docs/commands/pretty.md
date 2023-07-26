# `pretty` - Command Reference

> Prettifies JSON to make it human readable

## Description

Takes JSON from the STDIN and reformats it to make it human readable, then
outputs that to STDOUT.

## Usage

```
<stdin> -> pretty -> <stdout>
```

## Examples

```
» tout: json {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} -> pretty 
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

## See Also

* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`tout`](../commands/tout.md):
  Print a string to the STDOUT and set it's data-type