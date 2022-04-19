# _murex_ Shell Docs

## Command Reference: `map` 

> Creates a map from two data sources

## Description

This takes two parameters - which are code blocks - and combines them to output a key/value map in JSON.

The first block is the key and the second is the value.

## Usage

    map { code-block } { code-block } -> <stdout>

## Examples

    » map { tout: json (["key 1", "key 2", "key 3"]) } { tout: json (["value 1", "value 2", "value 3"]) } 
    {
        "key 1": "value 1",
        "key 2": "value 2",
        "key 3": "value 3"
    }

## See Also

* [commands/`@[` (range) ](../commands/range.md):
  Outputs a ranged subset of data from STDIN
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [commands/`a` (mkarray)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [commands/`alter`](../commands/alter.md):
  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input
* [commands/`append`](../commands/append.md):
  Add data to the end of an array
* [commands/`count`](../commands/count.md):
  Count items in a map, list or array
* [commands/`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [commands/`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [commands/`prepend` ](../commands/prepend.md):
  Add data to the start of an array