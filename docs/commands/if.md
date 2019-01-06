# _murex_ Language Guide

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
    
...or does not exist:

    !if { g somefile.txt } then {
        out "File does not exist"
    }

### Detail

The conditional block can contain entire pipelines - even multiple lines of code
let alone a single pipeline - as well as solitary commands as demonstrated in
the examples above. However the conditional block does not output STDOUT nor
STDERR to the rest of the pipeline so you don't have to worry about redirecting
the output streams to `null`.

### Synonyms

* `if`
* `!if`


### See Also

* [`and`](../docs/commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [`catch`](../docs/commands/catch.md):
  Handles the exception code raised by `try` or `trypipe
* [`or`](../docs/commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [`try`](../docs/commands/try.md):
  Handles errors inside a block of code
* [`trypipe`](../docs/commands/trypipe.md):
  Checks state of each function in a pipeline and exits block on error
* [false](../docs/commands/commands/false.md):
  
* [not](../docs/commands/commands/not.md):
  
* [true](../docs/commands/commands/true.md):
  