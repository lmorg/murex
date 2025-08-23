# Deprecated Features

Murex is committed to backwards compatibility. While we do want to continue to grow and improve the shell, this will not come at the expense of long term usability. However sometimes features do need to be removed to keep the shell focused and well supported.

## Pages


### Builtins

Builtin commands which have been removed from Murex

* [Arithmetic Evaluation: `=`](../deprecated/equ.md):
  Evaluate a mathematical function (removed 7.0)
* [Integer Operations: `let`](../deprecated/let.md):
  Evaluate a mathematical function and assign to variable (removed 7.0)

### Tokens And Operators

Tokens, operators, and other syntactic features which have been removed from Murex



### Uncategorised

* [Read With Type: `tread` (removed 7.x)](../deprecated/tread.md):
  `read` a line of input from the user and store as a user defined *typed* variable (deprecated)
* [`die`](../deprecated/die.md):
  Terminate murex with an exit number of 1 (deprecated)
