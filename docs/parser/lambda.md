# `[{lambda}]`

> Iterate through structured data

## Description

Lambdas, in Murex, are a concise way of performing various actions against
structured data. They're a convenience tool to range over arrays and objects,
similar to `foreach` and `formap`.

Code running inside a lambda inherit a special variable, named **meta values**,
which hold the state for each iteration.



## Detail

### Meta values

Meta values are a JSON object stored as the variable `$.`. The meta variable
will get overwritten by any other block which invokes meta values. So if you
wish to persist meta values across blocks you will need to reassign `$.`, eg

```
%[1..3] -> foreach {
    meta_parent = $.
    %[7..9] -> foreach {
        out "$(meta_parent.i): $.i"
    }
}
```

The following meta values are defined:

* `i`: iteration number (counts from one)
* `k`: key name (for arrays / lists this will count from zero)
* `v`: item value of map / object or array / list

## See Also

* [`foreach`](../commands/foreach.md):
  Iterate through an array
* [`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [`while`](../commands/while.md):
  Loop until condition false

<hr/>

This document was generated from [gen/parser/lambda_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/lambda_doc.yaml).