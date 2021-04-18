# _murex_ Shell Docs

## Command Reference: `2darray` 

> Create a 2D JSON array from multiple input sources

## Description

`2darray` merges multiple input sources to create a two dimensional array in JSON

## Usage

    2darray: { code-block } { code-block } -> <stdout>

## Examples

    » ps: -fe -> head: -n 10 -> set: ps 
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

## Detail

`2darray` can have as many or as few code blocks as you wish.

## See Also

* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`append`](../commands/append.md):
  Add data to the end of an array
* [commands/`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [types/`json` ](../types/json.md):
  JavaScript Object Notation (JSON) (primitive)
* [commands/`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`map` ](../commands/map.md):
  Creates a map from two data sources
* [commands/`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`prepend` ](../commands/prepend.md):
  Add data to the start of an array