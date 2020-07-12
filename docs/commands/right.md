# _murex_ Shell Docs

## Command Reference: `right`

> Right substring every item in a list

## Description

Takes a list from STDIN and returns a right substring of that same list.

One parameter is required and that is the number of characters to return. If
the parameter is a negative then `right` counts from the left.

## Usage

    <stdin> -> right int -> <stdout>

## Examples

Count from the right

    » ja: [Monday..Wednesday] -> right 4
    [
        "nday",
        "sday",
        "sday"
    ]
    
Count from the left

    » ja: [Monday..Wednesday] -> left -3
    [
        "day",
        "sday",
        "nesday"
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
* [commands/`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [commands/`right`](../commands/right.md):
  Right substring every item in a list
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`suffix`](../commands/suffix.md):
  Prefix a string to every item in a list
* [commands/length](../commands/length.md):
  
* [api/marshaldata](../api/marshaldata.md):
  
* [api/unmarshaldata](../api/unmarshaldata.md):
  