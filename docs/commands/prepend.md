# _murex_ Shell Docs

## Command Reference: `prepend` 

> Add data to the start of an array

## Description

`prepend` a data to the start of an array.

## Usage

    <stdin> -> prepend: value -> <stdout>

## Examples

    » a: [January..December] -> prepend: 'New Year'
    New Year
    January
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December

## Detail

`prepend` and `append` are data type aware:

    » tout json [1,2,3] -> append 4 5 6 bob
    Error in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing "bob": invalid syntax

## Synonyms

* `prepend`
* `list.prepend`


## See Also

* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`addheading` ](../commands/addheading.md):
  Adds headings to a table
* [commands/`append`](../commands/append.md):
  Add data to the end of an array
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`count`](../commands/count.md):
  Count items in a map, list or array
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings