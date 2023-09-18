# `!` (not)

> Reads the STDIN and exit number from previous process and not's it's condition

## Description

Reads the STDIN and exit number from previous process and not's it's condition.

## Usage

```
<stdin> -> ! -> <stdout>
```## Examples

```
» echo "Hello, world!" -> !
false
```

```
» false -> !
true
```

## Synonyms

* `!`


## See Also

* [and](../parser/and.md):
  
* [false](../parser/false.md):
  
* [if](../parser/if.md):
  
* [or](../parser/or.md):
  
* [true](../parser/true.md):
  

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).