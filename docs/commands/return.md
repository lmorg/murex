# Exit Function (`return`)

> Exits current function scope

## Description

`return` will terminate execution of a block at the scope level (eg `function`,
`private`, etc)

Conceptually it is the same as `break` except it doesn't require the scope name
as a parameter and you can specify the exit number rather than defaulting to 0.

## Usage

```
return [ exit-number ]
```

## Examples

### Setting an exit number

```
function example {
    out foo
    return 13
    out bar
}
example
exitnum
```

Running the above code would output:

```
foo
13
```

### Returning withing an exit number

If we were to run the same code as above but with `return` written without any
parameters (ie instead of `return 13` it would be just `return`), then you
would see the following output:

```
foo
0
```

## Detail

Any process that has been initialised within a `return`ed scope will have their
exit number updated to the value specified in `return` (or `0` if no parameter
was passed).

## See Also

* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [Get Exit Code (`exitnum`)](../commands/exitnum.md):
  Output the exit number of the previous process
* [Next Iteration (`continue`)](../commands/continue.md):
  Terminate process of a block within a caller function
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block

<hr/>

This document was generated from [builtins/core/structs/break_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/break_doc.yaml).