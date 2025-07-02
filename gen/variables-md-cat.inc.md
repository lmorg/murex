Variables are typed.

They can be a _primitive_ like `int` or `str`. They can also be a structured
document like `json`, `csv` or `sexpr`.

Variables can also have a string representation, for compatibility with older
POSIX idioms, as well as a native object format.

## Glossary Of Terms

To help better understand how variables work under the hood, blow is a glossary
of terms:

* _primitive_: this refers to the atomic component of a _data-type_. In other
  words, the smallest possible format for a piece of data. Where a JSON file
  might arrays and maps, the values for those objects cannot be divided any
  smaller than numbers, strings or a small number of constants like `true`,
  `false`, and `null`.

* _scope_: this is how far outside the code block that a particular variable
  can be written to, or read from.

* _local_ (scope): this refers to variables that cannot be read nor modified
  outside of the current function. Thus one function cannot read nor write to a
  variable in another function.

* _module_ (scope): these variables are accessible by any function or routine
  from within the same module. You'll only need _module scoped_ variables if
  you're writing modules -- and even then, only if you want that variable
  available to all functions within that module.

* _global_ (scope): these are variables which are accessible from any function,
  anywhere within Murex.

* _environmental variables_: sometimes written as _env vars_ for short, these
  are system variables. They can be passed from one process to another, so
  careful what secrets you store and what software you run while you have
  sensitive _env vars_ defined.

* _reserved variables_: this refers to variables that are read only. Some
  reserved variables are dynamic and thus can change their value depending on
  contextual circumstances.

* _POSIX_: this is a specification that Linux, Apple macOS, FreeBSD and its ilk
  follow. It defines a lot of the commonality between these environments.
  Windows and Plan 9 are not POSIX-compatible out-of-the-box but can support
  POSIX (eg via WSL).