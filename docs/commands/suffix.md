# _murex_ Shell Docs

## Command Reference: `suffix`

> Prefix a string to every item in a list

## Description

Takes a list from STDIN and returns that same list with each element suffixed.

## Usage

    <stdin> -> suffix str -> <stdout>

## Examples

    » ja: [Monday..Wednesday] -> suffix foobar
    [
        "Mondayfoobar",
        "Tuesdayfoobar",
        "Wednesdayfoobar"
    ]

## Detail

Supported data types can queried via `runtime`

    runtime: --marshallers
    runtime: --unmarshallers

## See Also

* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`left`](../commands/left.md):
  Left substring every item in a list
* [commands/`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [commands/`right`](../commands/right.md):
  Right substring every item in a list
* [commands/length](../commands/length.md):
  
* [api/marshaldata](../api/marshaldata.md):
  
* [api/unmarshaldata](../api/unmarshaldata.md):
  