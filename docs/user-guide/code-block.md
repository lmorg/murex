# Code Block Parsing

> Overview of how code blocks are parsed

The murex parser creates ASTs ahead of interpreting each block of code. However
the AST is only generated for a block at a time. Take this sample code:

```
function example {
    # An example function
    if { $ENVVAR } then {
        out 'foobar'
    }
    out 'Finished!'
}
```

When that code is run `function` is executed with the parameters `example` and
`{ ... }` but the contents of `{ ... }` isn't converted into ASTs until someone
calls `example` elsewhere in the shell.

When `example` (the Murex function defined above) is executed the parser will
then generate AST of the commands inside said function but not any blocks that
are associated with those functions. eg the AST would look something like this:

```
[
    {
        "Command": "if",
        "Parameters": [
            "{ $ENVVAR }",
            "then",
            "{\n        out 'foobar'\n    }"
        ]
    },
    {
        "Command": "out",
        "Parameters": [
            "Finished!"
        ]
    }
]
```

> Please note this is a mock JSON structure rather than a representation of the
> actual AST that would be created. Parameters are stored differently to allow
> infixing of variables; and there also needs to be data shared about how
> pipelining (eg STDOUT et al) is chained. What is being captured above is only
> the command name and parameters.

So when `if` executes, the conditional (the first parameter) is then parsed and
turned into ASTs and executed. Then the last parameter (the **then** block) is
parsed and turned into ASTs, if the first conditional is true.

This sequence of parsing is defined within the `if` builtin rather than
Murex's parser. That means any code blocks are parsed only when a builtin
specifically requests that they are executed.

With murex, there's no distinction between text and code. It's up to commands
to determine if they want to execute a parameter as code or not (eg a curly
brace block might be JSON).

## See Also

* [ANSI Constants](../user-guide/ansi.md):
  Infixed constants that return ANSI escape sequences
* [Pipeline](../user-guide/pipeline.md):
  Overview of what a "pipeline" is
* [Schedulers](../user-guide/schedulers.md):
  Overview of the different schedulers (or 'run modes') in Murex
* [`%(Brace Quote)`](../parser/brace-quote.md):
  Initiates or terminates a string (variables expanded)
* [`%[]` Create array](../parser/create-array.md):
  Quickly generate arrays
* [`%{}` Create object](../parser/create-object.md):
  Quickly generate objects and maps
* [`{ Curly Brace }`](../parser/curly-brace.md):
  Initiates or terminates a code block

<hr/>

This document was generated from [gen/parser/codeblock_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/codeblock_doc.yaml).