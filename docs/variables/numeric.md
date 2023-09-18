# Numeric (str)

> Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above

## Description

Variables named `0` and above are the equivalent index value of `@ARGV`.

These are reserved variable so they cannot be changed.



## Examples

```
» function example { out $0 $2 }
» example 1 2 3
example 2
```

## Detail

### `0` (str)

This returns the name of the executable (like `$ARGS[0]`)

### `1`, `2`, `3`... (str)

This returns parameter _n_ (like `$ARGS[n]`). If there is no parameter _n_
then the variable will not be set thus the upper limit variable is determined
by how many parameters are set. For example if you have 19 parameters passed
then variables `$1` through to `$19` (inclusive) will all be set.

## See Also

* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [autocomplete](../variables/autocomplete.md):
  
* [function](../variables/function.md):
  
* [out](../variables/out.md):
  
* [private](../variables/private.md):
  
* [set](../variables/set.md):
  
* [string](../variables/string.md):
  

<hr/>

This document was generated from [gen/variables/numeric_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/numeric_doc.yaml).