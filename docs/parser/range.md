# Filter By Range `[ ..Range ]`

> Outputs a ranged subset of data from stdin

## Description

This will read from stdin and output a subset of data in a defined range.

The range can be defined as a number of different range types - such as the
content of the array or it's index / row number. You can also omit either
the start or the end of the search criteria to cover all items before or
after the remaining search criteria.

**Please note that `@[` syntax has been deprecated in favour of `[` syntax
instead**

## Usage

```
<stdin> -> [start..end]flags -> <stdout>
```

## Examples

### Include everything after string match:

```
» a [January..December] -> [March..]se
April
May
June
July
August
September
October
November
December
```

### Range from the 6th to the 10th index

By default, ranges start from one, `1`:

```
» a [January..December] -> [5..9]
May
June
July
August
September
```

### Return the first 3

This usage is similar to `head -n3`:

```
» a [January..December] -> [..3]
October
November
December
```

### Return the last 3

This usage is similar to `tail -n3`:

```
» a [January..December] -> [-3..]
October
November
December
```

## Flags

* `8`
    handles backspace characters (char 8) instead of treating it like a printable character
* `b`
    removes blank (empty) lines from source
* `e`
    exclude the start and end search criteria from the range
* `n`
    numeric offset (indexed from 0)
* `r`
    regexp match
* `s`
    exact string match
* `t`
    trims whitespace from source

## Synonyms

* `@[`


## See Also

* [Alter Data Structure (`alter`)](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [Append To List (`append`)](../commands/append.md):
  Add data to the end of an array
* [Count (`count`)](../commands/count.md):
  Count items in a map, list or array
* [Create JSON Array (`ja`)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
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

This document was generated from [builtins/core/ranges/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/ranges/ranges_doc.yaml).