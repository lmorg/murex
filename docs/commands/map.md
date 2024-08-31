# Create Map (`map`)

> Creates a map from two data sources

## Description

This takes two parameters - which are code blocks - and combines them to output a key/value map in JSON.

The first block is the key and the second is the value.

## Usage

```
map { code-block } { code-block } -> <stdout>
```

## Examples

```
» map { tout json (["key 1", "key 2", "key 3"]) } { tout json (["value 1", "value 2", "value 3"]) } 
{
    "key 1": "value 1",
    "key 2": "value 2",
    "key 3": "value 3"
}
```

## Synonyms

* `map`


## See Also

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [Filter By Range `[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [Get Item (`[ Index ]`)](../parser/item-index.md):
  Outputs an element from an array, map or table
* [Get Nested Element (`[[ Element ]]`)](../parser/element.md):
  Outputs an element from a nested structure
* [Prepend To List (`prepend`)](../commands/prepend.md):
  Add data to the start of an array
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/arraytools/map_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/map_doc.yaml).