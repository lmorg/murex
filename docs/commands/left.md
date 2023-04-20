# `left` - Command Reference

> Left substring every item in a list

## Description

Takes a list from STDIN and returns a left substring of that same list.

One parameter is required and that is the number of characters to return. If
the parameter is a negative then `left` counts from the right.

## Usage

    <stdin> -> left int -> <stdout>

## Examples

Count from the left

    » ja: [Monday..Wednesday] -> left 2
    [
        "Mo",
        "Tu",
        "We"
    ]
    
Count from the right

    » ja: [Monday..Wednesday] -> left -3
    [
        "Mon",
        "Tues",
        "Wednes"
    ]

## Detail

Supported data types can queried via `runtime`

    runtime: --marshallers
    runtime: --unmarshallers

## Synonyms

* `left`
* `list.left`


## See Also

* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a _murex_ data-type (eg for stdio)
* [`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a _murex_ data-type into structured memory
* [`prefix`](../commands/prefix.md):
  Prefix a string to every item in a list
* [`right`](../commands/right.md):
  Right substring every item in a list
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`suffix`](../commands/suffix.md):
  Prefix a string to every item in a list