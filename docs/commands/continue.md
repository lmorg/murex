# _murex_ Shell Docs

## Command Reference: `continue`

> terminate process of a block within a caller function

## Description

`continue` will terminate execution of a block (eg `function`, `private`,
`foreach`, `if`, etc) right up until the caller function. In iteration loops
like `foreach` and `formap` this will result in behavior similar to the
`continue` statement in other programming languages.

## Usage

    continue block-name

## Examples

    %[1..10] -> foreach i {
        if { $i == 5 } then {
            out "continue"
            continue foreach
            out "skip this code"
        }
        out $i
    }
    
Running the above code would output:

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

## Detail

`continue` cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling `continue
bar` (which is a different function) inside of the function `foo`:

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
    
Regardless of whether we run `foo` or `bar`, both of those functions will
raise the following error:

    Error in `continue` (7,17): no block found named `bar` within the scope of `foo`

## See Also

* [commands/`break`](../commands/break.md):
  terminate execution of a block within your processes scope
* [commands/`exit`](../commands/exit.md):
  Exit murex
* [commands/`foreach`](../commands/foreach.md):
  Iterate through an array
* [commands/`formap`](../commands/formap.md):
  Iterate through a map or other collection of data
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [commands/`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [commands/`private`](../commands/private.md):
  Define a private function block