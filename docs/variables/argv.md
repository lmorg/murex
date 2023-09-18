# `ARGV` (json)

> Array of the command name and parameters within a given scope

## Description

`ARGV` returns an array of the command name and parameters within a given
scope. eg `function`, `private`, `autocomplete` or shell script.

Unlike `$PARAMS`, `$ARGV` includes the function name.

This is a reserved variable so it cannot be changed.



## Examples

```
» function example { $ARGV }
» example abc 1 2 3
[
    "example",
    "abc",
    "1",
    "2",
    "3"
]
```

## Detail

### Deprecation of `ARGS`

In Murex versions 4.x and below, this variable was named `ARGS` (with an 'S').
However in Murex 5.x and above it was renamed to `ARGV` (with a 'V') to unify
the name with other languages.

`ARGS` will remain available for compatibility reasons but is considered
deprecated and may be removed from future releases.

## Synonyms

* `argv`
* `ARGV`
* `ARGS`


## See Also

* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [array](../variables/array.md):
  
* [autocomplete](../variables/autocomplete.md):
  
* [function](../variables/function.md):
  
* [json](../variables/json.md):
  
* [modules](../variables/modules.md):
  
* [out](../variables/out.md):
  
* [pipeline](../variables/pipeline.md):
  
* [private](../variables/private.md):
  
* [reserved-vars](../variables/reserved-vars.md):
  
* [scoping](../variables/scoping.md):
  
* [set](../variables/set.md):
  
* [string](../variables/string.md):
  

<hr/>

This document was generated from [gen/variables/ARGV_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/ARGV_doc.yaml).