# Match String (`match`)

> Match an exact value in an array

## Description

`match` takes input from stdin and returns any array items / lines which
contain an exact match of the parameters supplied.

When multiple parameters are supplied they are concatenated into the search
string and white space delimited. eg all three of the below are the same:

```
match "a b c"
match a\sb\sc
match a b c
match a    b    c
```

If you want to return everything except the search string then use `!match`

## Usage

Match every occurrence of search string

```
<stdin> -> match search string -> <stdout>
```

Match everything except search string

```
<stdin> -> !match search string -> <stdout>
```

## Examples

### Return matched

Match **Wed**

```
» ja [Monday..Friday] -> match Wed
[
    "Wednesday"
]
```

### Ignore matched

Match everything except **Wed**

```
» ja [Monday..Friday] -> !match Wed
[
    "Monday",
    "Tuesday",
    "Thursday",
    "Friday"
] 
```

## Detail

`match` is data-type aware so will work against lists or arrays of whichever
Murex data-type is passed to it via stdin and return the output in the
same data-type.

## Synonyms

* `match`
* `!match`
* `list.str`
* `!list.str`


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
* [Regex Operations (`regexp`)](../commands/regexp.md):
  Regexp tools for arrays / lists of strings
* [Sort Array (`msort`)](../commands/msort.md):
  Sorts an array - data type agnostic
* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [Stream New List (`a`)](../commands/a.md):
  A sophisticated yet simple way to stream an array or list (mkarray)

<hr/>

This document was generated from [builtins/core/lists/regexp_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/regexp_doc.yaml).