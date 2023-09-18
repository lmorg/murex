# `PWDHIST` (json)

> History of each change to the sessions working directory

## Description

`PWDHIST` is a JSON array containing the history of all the working directories
within the current shell session.

It is updated via `cd` however you can overwrite its value manually via `set`.



## Examples

```
» cd ~bob
» cd /tmp
» $PWDHIST
[
    "/Users/bob",
    "/private/tmp"
]
```

## Synonyms

* `pwdhist`
* `PWDHIST`


## See Also

* [`PWD` (str)](../variables/pwd.md):
  Current working directory
* [array](../variables/array.md):
  
* [cd](../variables/cd.md):
  
* [json](../variables/json.md):
  
* [modules](../variables/modules.md):
  
* [path](../variables/path.md):
  
* [pipeline](../variables/pipeline.md):
  
* [reserved-vars](../variables/reserved-vars.md):
  
* [scoping](../variables/scoping.md):
  
* [set](../variables/set.md):
  
* [string](../variables/string.md):
  

<hr/>

This document was generated from [gen/variables/PWDHIST_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/PWDHIST_doc.yaml).