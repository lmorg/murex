# `list.join` 

> Joins a list or array into a single string

## Description

`mjoin` will read an array from either stdin or it's command line parameters,
and joins those elements together to form a single string.

The string will be delimited by the separator defined as the first command line
parameter.

## Usage

```
<stdin> -> list.join separator                      -> <stdout>
           list.join separator item1 item2 item3... -> <stdout>
```

## Examples

### As a method

```
» %[Monday..Friday] -> list.join !
Monday!Tuesday!Wednesday!Thursday!Friday
```

### As a function

```
» list.join ! @{ %[Monday..Friday] }
Monday!Tuesday!Wednesday!Thursday!Friday
```

## Synonyms

* `list.join`
* `mjoin`


## See Also

* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array
* [`jsplit` ](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter

<hr/>

This document was generated from [builtins/core/lists/mjoin_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/mjoin_doc.yaml).