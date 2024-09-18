# Right Sub-String (`right`)

> Right substring every item in a list

## Description

Takes a list from stdin and returns a right substring of that same list.

One parameter is required and that is the number of characters to return. If
the parameter is a negative then `right` counts from the left.

## Usage

```
<stdin> -> right int -> <stdout>
```

## Examples

### Count from the right

```
» ja [Monday..Wednesday] -> right 4
[
    "nday",
    "sday",
    "sday"
]
```

### Count from the left

```
» ja [Monday..Wednesday] -> left -3
[
    "day",
    "sday",
    "nesday"
]
```

## Detail

Supported data types can queried via `runtime`

```
runtime --marshallers
runtime --unmarshallers
```

## Synonyms

* `right`
* `list.right`


## See Also

* [Add Prefix (`prefix`)](../commands/prefix.md):
  Prefix a string to every item in a list
* [Add Suffix (`suffix`)](../commands/suffix.md):
  Prefix a string to every item in a list
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Right Sub-String (`right`)](../commands/right.md):
  Right substring every item in a list
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)
* [`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)
* [`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory

<hr/>

This document was generated from [builtins/core/lists/push_pop_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/push_pop_doc.yaml).