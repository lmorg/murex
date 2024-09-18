# Logic Or Statements (`or`)

> Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.

## Description

Returns a boolean results (`true` or `false`) depending on whether any of the
code-blocks included as parameters are successful or not.

## Usage

```
or { code-block } { code-block } -> <stdout>

!or { code-block } { code-block } -> <stdout>
```

`or` supports as many or as few code-blocks as you wish.

## Examples

```
if { or { = 1+1==2 } { = 2+2==5 } { = 3+3==6 } } then {
    out At least one of those equations are correct
}
```

## Detail

`or` does not set the exit number on failure so it is safe to use inside a `try`
or `trypipe` block.

If `or` is prefixed by a bang (`!or`) then it returns `true` when one or more
code-blocks are unsuccessful (ie the opposite of `or`).

### Code-Block Testing

* `or` only executes code-blocks up until one of the code-blocks is successful
  then it exits the function and returns `true`.

* `!or` only executes code-blocks while the code-blocks are successful. Once one
  is unsuccessful `!or` exits and returns `true` (ie it `not`s every code-block).

## Synonyms

* `or`
* `!or`


## See Also

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Logic And Statements (`and`)](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Not (`!`)](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [Pipe Fail (`trypipe`)](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [True (`true`)](../commands/true.md):
  Returns a `true` value
* [Try Block (`try`)](../commands/try.md):
  Handles non-zero exits inside a block of code
* [`&&` And Logical Operator](../parser/logical-and.md):
  Continues next operation if previous operation passes
* [`||` Or Logical Operator](../parser/logical-or.md):
  Continues next operation only if previous operation fails

<hr/>

This document was generated from [builtins/core/structs/andor_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/andor_doc.yaml).