# `$.`, Meta Values (json)

> State information for iteration blocks

## Description

Meta Values, `$.`, provides state information for blocks like `foreach`,
`formap`, `while` and lambdas.

Meta Values are a specific to the block, so you will need to refer to each
iteration structure's documentation to check what information is exposed via
`$.`

## Examples

```
» %[Monday..Friday] -> foreach day { out "$.i: $day" }
1: Monday
2: Tuesday
3: Wednesday
4: Thursday
5: Friday
```

## See Also

* [For Each In Map: `formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [For Each In array: `foreach`](../commands/foreach.md):
  Iterate through an array
* [Loop While: `while`](../commands/while.md):
  Loop until condition false
* [`[{ Lambda }]`](../parser/lambda.md):
  Iterate through structured data

<hr/>

This document was generated from [gen/variables/meta-values_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/meta-values_doc.yaml).