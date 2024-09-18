# Exit Block (`break`)

> Terminate execution of a block within your processes scope

## Description

`break` will terminate execution of a block (eg `function`, `private`,
`foreach`, `if`, etc).

`break` requires a parameter and that parameter is the name of the caller
block you wish to break out of. If it is a `function` or `private`, then it
will be the name of that function or private. If it is an `if` or `foreach`
loop, then it will be `if` or `foreach` (respectively).

## Usage

```
break block-name
```

## Examples

### Exiting an iteration block

```
function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            break foo
            out "ended"
        }
    }
}
```

Running the above code would output:

```
» foo
1
2
3
4
5
exit running function
```

### Exiting a function

`break` can be considered to exhibit the behavior of _return_ (from other
languages) too

```
function example {
    if { $USER == "root" } then {
        err "Don't run this as root"
        break example
    }
    
    # ... do something ...
}
```

Though in this particular use case it is recommended that you use `return`
instead, the above code does illustrate how `break` behaves.

## Detail

`break` cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling `break
bar` (which is a different function) inside of the function `foo`:

```
function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            break bar
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
Error in `break` (7,17): no block found named `bar` within the scope of `foo`
```

## See Also

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