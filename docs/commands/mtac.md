# Reverse Array (`mtac`)

> Reverse the order of an array

## Description

`mtac` takes input from stdin and reverses the order of it.

It's name is derived from a program called `tac` - a tool that functions
like `cat` but returns the contents in the reverse order. The difference
with the `mtac` builtin is that it is data-type aware. So it doesn't just
function as a replacement for `tac` but it also works on JSON arrays,
s-expressions, and any other data-type supporting arrays compiled into
Murex.

## Usage

```
<stdin> -> mtac -> <stdout>
```

## Examples

```
» ja [Monday..Friday] -> mtac
[
    "Friday",
    "Thursday",
    "Wednesday",
    "Tuesday",
    "Monday"
]

# Normal output (without mtac)
» ja [Monday..Friday]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday"
]
```

## Detail

Please bare in mind that while Murex is optimised with concurrency and
streaming in mind, it's impossible to reverse an incomplete array. Thus all
all of stdin must have been read and that file closed before `mtac` can
output.

In practical terms you shouldn't notice any difference except for when
stdin is a long running process or non-standard stream (eg network pipe).

## Synonyms

* `mtac`
* `list.reverse`


## See Also

* [Add Prefix (`prefix`)](../commands/prefix.md):
  Prefix a string to every item in a list
* [Add Suffix (`suffix`)](../commands/suffix.md):
  Prefix a string to every item in a list
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create 2d Array (`2darray`)](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Create Map (`map`)](../commands/map.md):
  Creates a map from two data sources
* [Create New Array (`ta`)](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Prettify JSON](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/lists/mtac_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/mtac_doc.yaml).