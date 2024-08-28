# struct.new.2darray: `2darray` 

> Create a 2D JSON array from multiple input sources

## Description

`2darray` merges multiple input sources to create a two dimensional array in JSON

## Usage

```
2darray { code-block } { code-block } ... -> <stdout>
```

## Examples

```
» ps -fe -> head -n 10 -> set ps 
» 2darray { $ps[UID] } { $ps[PID] } { $ps[TTY] } { $ps[TIME] }
[
    [
        "",
        "",
        "",
        ""
    ],
    [
        "UID",
        "PID",
        "TTY",
        "TIME"
    ],
    [
        "root",
        "1",
        "?",
        "00:00:02"
    ],
    [
        "root",
        "2",
        "?",
        "00:00:00"
    ],
    [
        "root",
        "3",
        "?",
        "00:00:00"
    ],
    [
        "root",
        "4",
        "?",
        "00:00:00"
    ],
    [
        "root",
        "6",
        "?",
        "00:00:00"
    ],
    [
        "root",
        "8",
        "?",
        "00:00:00"
    ],
    [
        "root",
        "9",
        "?",
        "00:00:03"
    ],
    [
        "root",
        "10",
        "?",
        "00:00:19"
    ],
    [
        "root",
        "11",
        "?",
        "00:00:01"
    ]
]
```

## Detail

`2darray` can have as many or as few code blocks as you wish.

## Synonyms

* `2darray`
* `struct.new.2darray`


## See Also

* [`[ ..Range ]`](../parser/range.md):
  Outputs a ranged subset of data from stdin
* [`[ Index ]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`json`](../types/json.md):
  JavaScript Object Notation (JSON)
* [list.append: `append`](../commands/append.md):
  Add data to the end of an array
* [list.new.str: `a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [list.prepend: `prepend`](../commands/prepend.md):
  Add data to the start of an array
* [list.reverse: `mtac`](../commands/mtac.md):
  Reverse the order of an array
* [list.sort: `msort`](../commands/msort.md):
  Sorts an array - data type agnostic
* [str.split: `jsplit` ](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [struct.count: `count`](../commands/count.md):
  Count items in a map, list or array
* [struct.new.map: `map`](../commands/map.md):
  Creates a map from two data sources

<hr/>

This document was generated from [builtins/core/arraytools/2darray_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/arraytools/2darray_doc.yaml).