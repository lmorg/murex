# Deprecated Features

Murex is committed to backwards compatibility. While we do want to continue to grow and improve the shell, this will not come at the expense of long term usability. However sometimes features do need to be removed to keep the shell focused and well maintained.

Please read our [compatibility commitment](https://murex.rocks/compatibility.html) for more information on Murex's
compatibility commitment

## Pages


### Builtins

Builtin commands which have been removed from Murex

* [Arithmetic Evaluation: `=`](../deprecated/equ.md):
  Evaluate a mathematical function (removed 7.0)
* [Exit after error: `die`](../deprecated/die.md):
  Terminate murex with an exit number of 1 (removed 7.0)
* [Integer Operations: `let`](../deprecated/let.md):
  Evaluate a mathematical function and assign to variable (removed 7.0)
* [Read With Type: `tread`](../deprecated/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (removed 7.0)

### Tokens And Operators

Tokens, operators, and other syntactic features which have been removed from Murex

* [`?` stderr Pipe](../deprecated/pipe-err.md):
  Pipes stderr from the left hand command to stdin of the right hand command (removed 8.0)


