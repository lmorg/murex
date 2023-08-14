# `[` (range)

> Outputs a ranged subset of data from STDIN

## Description

This will read from STDIN and output a subset of data in a defined range.

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

**Range over all months after March:**

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

**Range from the 6th to the 10th month:**

By default, ranges start from one, `1`

```
» a [January..December] -> [5..9]
May
June
July
August
September
```

**Return the first 3 months:**

This usage is similar to `head -n3`

```
» a [January..December] -> [..3]
October
November
December
```

**Return the last 3 months:**

This usage is similar to `tail -n3`

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

* [`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`alter`](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`count`](../commands/count.md):
  Count items in a map, list or array
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`prepend`](../commands/prepend.md):
  Add data to the start of an array
* [index](../commands/item-index.md):
  Outputs an element from an array, map or table