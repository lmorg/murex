# `len`

> Outputs the length of an array

## Description

This will read an array from STDIN and outputs the length for that array

## Usage

    `<stdin>` -> len -> `<stdout>`

## Examples

```
» tout: json (["a", "b", "c"]) -> len
3
```

## Detail

Please note that this returns the length of the _array_ rather than string.
For example `out "foobar" -> len` would return `1` because an array in the
`str` data type would be new line separated (eg `out "foo\nbar" -> len`
would return `2`). If you need to count characters in a string and are
running POSIX (eg Linux / BSD / OSX) then it is recommended to use `wc`
instead. But be mindful that `wc` will also count new line characters

```
» out: "foobar" -> len
1

» out: "foo\nbar" -> len
2

» out: "foobar" -> wc: -c
7

» out: "foo\nbar" -> wc: -c
8

» printf: "foobar" -> wc: -c
6
# (printf does not print a trailing new line)
```

## See Also

- [commands/`@[` (range) ](./range.md):
  Outputs a ranged subset of data from STDIN
- [commands/`[[` (element)](./element.md):
  Outputs an element from a nested structure
- [commands/`[` (index)](./index2.md):
  Outputs an element from an array, map or table
- [commands/`a` (mkarray)](./a.md):
  A sophisticated yet simple way to build an array or list
- [commands/`append`](./append.md):
  Add data to the end of an array
- [commands/`ja` (mkarray)](./ja.md):
  A sophisticated yet simply way to build a JSON array
- [commands/`jsplit` ](./jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
- [commands/`map` ](./map.md):
  Creates a map from two data sources
- [commands/`msort` ](./msort.md):
  Sorts an array - data type agnostic
- [commands/`mtac`](./mtac.md):
  Reverse the order of an array
- [commands/`prepend` ](./prepend.md):
  Add data to the start of an array
