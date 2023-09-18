# `[..range]`

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
```## Examples

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

* [`[[ element ]]`](../parser/element.md):
  Outputs an element from a nested structure
* [`[index]`](../parser/item-index.md):
  Outputs an element from an array, map or table
* [a](../parser/a.md):
  
* [alter](../parser/alter.md):
  
* [append](../parser/append.md):
  
* [count](../parser/count.md):
  
* [ja](../parser/ja.md):
  
* [jsplit](../parser/jsplit.md):
  
* [prepend](../parser/prepend.md):
  

<hr/>

This document was generated from [builtins/core/ranges/ranges_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/ranges/ranges_doc.yaml).