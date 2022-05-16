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
* [commands/`count`](../commands/count.md):
  Count items in a map, list or array
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [apis/`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [apis/`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a _murex_ data-type into structured memory
* [commands/`left`](../commands/left.md):
  Left substring every item in a list
* [commands/`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [commands/`right`](../commands/right.md):
  Right substring every item in a list
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_