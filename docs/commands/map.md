# _murex_ Language Guide

## Command Reference: `map` 

> Creates a map from two data sources

### Description

This takes two parameters - which are code blocks - and combines them to output a key/value map in JSON.

The first block is the key and the second is the value.

### Usage

    map { code-block } { code-block } -> <stdout>

### Examples

    » { tout: json (["key 1", "key 2", "key 3"]) } { tout: json (["value 1", "value 2", "value 3"]) } 
    {
        "key 1": "value 1",
        "key 2": "value 2",
        "key 3": "value 3"
    }

### See Also

* [`append`](../commands/append.md):
  Add data to the end of an array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [a](../commands/a.md):
  
* [ja](../commands/ja.md):
  