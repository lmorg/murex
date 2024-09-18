# Left Sub-String (`left`)

> Left substring every item in a list

## Description

Takes a list from stdin and returns a left substring of that same list.

One parameter is required and that is the number of characters to return. If
the parameter is a negative then `left` counts from the right.

## Usage

```
<stdin> -> left int -> <stdout>
```

## Examples

### Count from the left

```
» ja [Monday..Wednesday] -> left 2
[
    "Mo",
    "Tu",
    "We"
]
```

### Count from the right

```
» ja [Monday..Wednesday] -> left -3
[
    "Mon",
    "Tues",
    "Wednes"
]
```

## Detail

Supported data types can queried via `runtime`

```
runtime --marshallers
runtime --unmarshallers
```

## Synonyms

* `left`
* `list.left`


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