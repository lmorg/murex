# list.suffix

> Prefix a string to every item in a list

## Description

Takes a list from stdin and returns that same list with each element suffixed.

## Usage

```
<stdin> -> suffix str -> <stdout>
```

## Examples

```
» ja [Monday..Wednesday] -> suffix foobar
[
    "Mondayfoobar",
    "Tuesdayfoobar",
    "Wednesdayfoobar"
]
```

## Detail

Supported data types can queried via `runtime`

```
runtime --marshallers
runtime --unmarshallers
```

## Synonyms

* `suffix`
* `list.suffix`


## See Also

* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`lang.MarshalData()` (system API)](../apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)
* [`lang.UnmarshalData()` (system API)](../apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory
* [list.left](../commands/left.md):
  Left substring every item in a list
* [list.new.str (`a`)](../commands/a.md):
  A sophisticated yet simple way to build an array or list (mkarray)
* [list.prefix](../commands/prefix.md):
  Prefix a string to every item in a list
* [list.right](../commands/right.md):
  Right substring every item in a list
* [shell.runtime](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [struct.count](../commands/count.md):
  Count items in a map, list or array

<hr/>

This document was generated from [builtins/core/lists/push_pop_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/push_pop_doc.yaml).