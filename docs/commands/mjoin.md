# Join Array To String (`mjoin`)

> Joins a list or array into a single string

## Description

`mjoin` will read an array from either stdin or it's command line parameters,
and joins those elements together to form a single string.

The string will be delimited by the separator defined as the first command line
parameter.

## Usage

```
<stdin> -> mjoin separator                           -> <stdout>
           mjoin separator item1 [ item2 item3 ... ] -> <stdout>
```

## Examples

### As a method

```
» %[Monday..Friday] -> mjoin !
Monday!Tuesday!Wednesday!Thursday!Friday
```

### As a function

```
» mjoin ! @{ %[Monday..Friday] }
Monday!Tuesday!Wednesday!Thursday!Friday
```

## Synonyms

* `mjoin`
* `list.join`


## See Also

* [Split String (`jsplit`)](../commands/jsplit.md):
  Splits stdin into a JSON array based on a regex parameter
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`@Array` Sigil](../parser/array.md):
  Expand values as an array

<hr/>

This document was generated from [builtins/core/lists/mjoin_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/lists/mjoin_doc.yaml).