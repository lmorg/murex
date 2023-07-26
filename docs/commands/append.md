# `append`

> Add data to the end of an array

## Description

`append` data to the end of an array.

## Usage

    <stdin> -> append: value -> <stdout>

## Examples

```
» a: [Monday..Sunday] -> append: Funday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
Funday
```

## Detail

`prepend` and `append` are data type aware:

```
» tout json [1,2,3] -> append 4 5 6 bob
Error in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing "bob": invalid syntax
```

## Synonyms

- `append`
- `list.append`

## See Also

- [`[[` (element)](./element.md):
  Outputs an element from a nested structure
- [`[` (index)](./index2.md):
  Outputs an element from an array, map or table
- [`[` (index)](./index2.md):
  Outputs an element from an array, map or table
- [`[` (range) ](./range.md):
  Outputs a ranged subset of data from STDIN
- [`a` (mkarray)](./a.md):
  A sophisticated yet simple way to build an array or list
- [`addheading` ](./addheading.md):
  Adds headings to a table
- [`cast`](./cast.md):
  Alters the data type of the previous function without altering it's output
- [`count`](./count.md):
  Count items in a map, list or array
- [`ja` (mkarray)](./ja.md):
  A sophisticated yet simply way to build a JSON array
- [`match`](./match.md):
  Match an exact value in an array
- [`msort` ](./msort.md):
  Sorts an array - data type agnostic
- [`mtac`](./mtac.md):
  Reverse the order of an array
- [`prepend` ](./prepend.md):
  Add data to the start of an array
- [`regexp`](./regexp.md):
  Regexp tools for arrays / lists of strings
