# `PARAMS` (json)

> Array of the parameters within a given scope

## Description

`PARAMS` returns an array of the parameters within a given scope.
eg `function`, `private`, `autocomplete` or shell script.

Unlike `$ARGV`, `$PARAMS` does not include the function name.

This is a reserved variable so it cannot be changed.



## Examples

```
» function example { $PARAMS }
» example abc 1 2 3
[
    "abc",
    "1",
    "2",
    "3"
]
```

## Synonyms

* `params`
* `PARAMS`


## See Also

* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
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

This document was generated from [gen/variables/PARAMS_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PARAMS_doc.yaml).