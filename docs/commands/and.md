# Logic And Statements (`and`)

> Returns `true` or `false` depending on whether multiple conditions are met

## Description

Returns a boolean results (`true` or `false`) depending on whether all of the
code-blocks included as parameters are successful or not.

## Usage

```
and { code-block } { code-block } -> <stdout>

!and { code-block } { code-block } -> <stdout>
```

`and` supports as many or as few code-blocks as you wish.

## Examples

```
if { and { = 1+1==2 } { = 2+2==4 } { = 3+3==6 } } then {
    out The laws of mathematics still exist in this universe.
}
```

## Detail

`and` does not set the exit number on failure so it is safe to use inside a `try`
or `trypipe` block.

If `and` is prefixed by a bang then it returns `true` only when all code-blocks
are unsuccessful.

### Code-Block Testing

* `and` tests all code-blocks up until one of the code-blocks is unsuccessful,
  then `and` exits and returns `false`.

* `!and` tests all code-blocks up until one of the code-blocks is successful,
  then `!and` exits and returns `false` (ie `!and` is `not`ing every code-block).

## Synonyms

* `and`
* `!and`


## See Also

* [Caught Error Block (`catch`)](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [False (`false`)](../commands/false.md):
  Returns a `false` value
* [If Conditional (`if`)](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [Logic Or Statements (`or`)](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
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