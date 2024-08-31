# Next Iteration (`continue`)

> Terminate process of a block within a caller function

## Description

`continue` will terminate execution of a block (eg `function`, `private`,
`foreach`, `if`, etc) right up until the caller function. In iteration loops
like `foreach` and `formap` this will result in behavior similar to the
`continue` statement in other programming languages.

## Usage

```
continue block-name
```

## Examples

```
%[1..10] -> foreach i {
    if { $i == 5 } then {
        out "continue"
        continue foreach
        out "skip this code"
    }
    out $i
}
```

Running the above code would output:

```
» foo
1
2
3
4
continue
6
7
8
9
10
```

## Detail

`continue` cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling `continue
bar` (which is a different function) inside of the function `foo`:

```
function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            continue bar
            out "ended"
        }
    }
}

function bar {
    foo
}
```

Regardless of whether we run `foo` or `bar`, both of those functions will
raise the following error:

```
Error in `continue` (7,17): no block found named `bar` within the scope of `foo`
```

## See Also

* [Exit Block (`break`)](../commands/break.md):
  Terminate execution of a block within your processes scope
* [Exit Function (`return`)](../commands/return.md):
  Exits current function scope
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [For Each In List (`foreach`)](../commands/foreach.md):
  Iterate through an array
* [For Each In Map (`formap`)](../commands/formap.md):
  Iterate through a map or other collection of data
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block

<hr/>

This document was generated from [builtins/core/structs/break_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/break_doc.yaml).