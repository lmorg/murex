# _murex_ Shell Docs

## Command Reference: `right`

> Right substring a list

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
* [commands/`right`](../commands/right.md):
  Right substring a list
* [api/marshaldata](../api/marshaldata.md):
  
* [commands/prefix](../commands/prefix.md):
  
* [commands/suffix](../commands/suffix.md):
  
* [api/unmarshaldata](../api/unmarshaldata.md):
  