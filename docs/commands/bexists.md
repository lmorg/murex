# Check Builtin Exists (`bexists`)

> Check which builtins exist

## Description

`bexists` takes an array of parameters and returns which commands have been
compiled into Murex. The 'b' in `bexists` stands for 'builtins'

## Usage

```
bexists command... -> <stdout>
```

## Examples

```
» bexists qr gzip runtime config
{
    "Installed": [
        "runtime",
        "config"
    ],
    "Missing": [
        "qr",
        "gzip"
    ]
}
```

## Detail

This builtin dates back to the start of Murex when all of the builtins were
considered optional. This was intended to be a way for scripts to determine
which builtins were compiled. Since then `runtime` has absorbed and centralized
a number of similar commands which have since been deprecated. The same fate
might also happen to `bexists` however it is in use by a few modules and for
that reason alone it has been spared from the axe.

## Synonyms

* `bexists`


## See Also

* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Modules And Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Shell Runtime (`runtime`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/management/functions_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/functions_doc.yaml).