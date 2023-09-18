# `&&` And Logical Operator

> Continues next operation if previous operation passes

## Description

When in the **normal** run mode (see "schedulers" link below) this will only
run the command on the right hand side if the command on the left hand side
does not error. Neither STDOUT nor STDERR are piped.

This has no effect in `try` nor `trypipe` run modes because they automatically
apply stricter error handling.

## Examples

Second command runs because the first command doesn't error:

```
» out one && out two
one
two
```

Second command does not run because the first command produces an error:

```
» err one && out two
one
```

## Detail

This is equivelent to a `try` block:

```
try {
    err one
    out two
}
```

## See Also

* [`?` STDERR Pipe](../parser/pipe-err.md):
  Pipes STDERR from the left hand command to STDIN of the right hand command
* [err](../parser/err.md):
  
* [out](../parser/out.md):
  
* [pipeline](../parser/pipeline.md):
  
* [schedulers](../parser/schedulers.md):
  
* [try](../parser/try.md):
  
* [trypipe](../parser/trypipe.md):
  

<hr/>

This document was generated from [gen/parser/logical_ops_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/logical_ops_doc.yaml).