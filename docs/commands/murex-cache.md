# `murex-cache`

> Management interface for Murex's cache database

## Description

`murex-cache` provides a management interface for Murex's cache database.

## Usage

```
murex-cache flags -> <stdout>
```

## Flags

* `clear`
    Empties the cache database
* `get`
    Returns the cached item from database (requires two additional parameters for namespace and key)
* `namespaces`
    List namespaces in cache database
* `trim`
    Removes expired cache items from database

## See Also

* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/cache/cache_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/cache/cache_doc.yaml).