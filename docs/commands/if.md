# _murex_ Shell Guide

## Command Reference: `if`

> Conditional statement to execute different blocks of code depending on the result of the condition

### Description

Conditional control flow

`if` can be utilized both as a method as well as a standalone function. As a
method, the conditional state is derived from the calling function (eg if the
previous function succeeds then the condition is `true`).

### Usage

#### Function `if`:

    if { code-block } then {
        # true
    } else {
        # false
    }
    
#### Method `if`:

    command -> if {
        # true
    } else {
        # false
    }
    
#### Negative Function `if`:

    !if { code-block } then {
        # false
    }
    
#### Negative Method `if`:

    command -> !if {
        # false
    }
    
#### Please Note:
the `then` and `else` statements are optional. So the first usage could
also be written as:

    if { code-block } {
        # true
    } {
        # false
    }
    
However the practice of omitting those statements isn't recommended beyond
writing short one liners in the interactive command prompt.

### Examples

Check if a file exists:

    if { g somefile.txt } then {
        out "File exists"
    }
    
...or does not exist (both ways are valid):

    !if { g somefile.txt } then {
        out "File does not exist"
    }
    
    if { g somefile.txt } else {
        out "File does not exist"
    }

### Detail

The conditional block can contain entire pipelines - even multiple lines of code
let alone a single pipeline - as well as solitary commands as demonstrated in
the examples above. However the conditional block does not output STDOUT nor
STDERR to the rest of the pipeline so you don't have to worry about redirecting
the output streams to `null`.

If you require output from the conditional blocks STDOUT then you will need to
use either a _murex_ named pipe to redirect the output, or test or debug flags
(depending on your use case) if you only need to occasionally inspect the
conditionals output.

### Synonyms

* `if`
* `!if`


### See Also

* [commands/`!` (not)](../commands/not.md):
  Reads the STDIN and exit number from previous process and not's it's condition
* [commands/`and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [commands/`catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe` 
* [commands/`debug`](../commands/debug.md):
  Debugging information
* [commands/`false`](../commands/false.md):
  Returns a `false` value
* [commands/`or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [commands/`true`](../commands/true.md):
  Returns a `true` value
* [commands/`try`](../commands/try.md):
  Handles errors inside a block of code
* [commands/`trypipe`](../commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [commands/test](../commands/test.md):
  