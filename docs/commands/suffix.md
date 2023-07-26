# `suffix`

> Prefix a string to every item in a list

## Description

Takes a list from STDIN and returns that same list with each element suffixed.

## Usage

    <stdin> -> suffix str -> <stdout>

## Examples

```
» ja: [Monday..Wednesday] -> suffix foobar
[
    "Mondayfoobar",
    "Tuesdayfoobar",
    "Wednesdayfoobar"
]
```

## Detail

Supported data types can queried via `runtime`

```
runtime: --marshallers
runtime: --unmarshallers
```

## Synonyms

- `suffix`
- `list.suffix`

## See Also

- [`a` (mkarray)](./a.md):
  A sophisticated yet simple way to build an array or list
- [`count`](./count.md):
  Count items in a map, list or array
- [`ja` (mkarray)](./ja.md):
  A sophisticated yet simply way to build a JSON array
- [`lang.MarshalData()` (system API)](/apis/lang.MarshalData.md):
  Converts structured memory into a Murex data-type (eg for stdio)
- [`lang.UnmarshalData()` (system API)](/apis/lang.UnmarshalData.md):
  Converts a Murex data-type into structured memory
- [`left`](./left.md):
  Left substring every item in a list
- [`prefix`](./prefix.md):
  Prefix a string to every item in a list
- [`right`](./right.md):
  Right substring every item in a list
- [`runtime`](./runtime.md):
  Returns runtime information on the internal state of Murex
