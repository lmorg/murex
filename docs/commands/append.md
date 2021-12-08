# _murex_ Shell Docs

## Command Reference: `append`

> Add data to the end of an array

## Description

`append` data to the end of an array.

## Usage

    <stdin> -> append: value -> <stdout>

## Examples

    » a: [Monday..Sunday] -> append: Funday
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
    Funday

## Detail

It's worth noting that `prepend` and `append` are not data type aware. So 
any integers in data type aware structures will be converted into strings:

    » tout: json [1,2,3] -> append: new 
    [
        "1",
        "2",
        "3",
        "new"
    ]

## See Also

* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`len` ](../commands/len.md):
  Outputs the length of an array
* [commands/`match`](../commands/match.md):
  Match an exact value in an array
* [commands/`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [commands/`mtac`](../commands/mtac.md):
  Reverse the order of an array
* [commands/`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [commands/`regexp`](../commands/regexp.md):
  Regexp tools for arrays / lists of strings