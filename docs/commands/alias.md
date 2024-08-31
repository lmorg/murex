# Alias Pointer (`alias`)

> Create an alias for a command

## Description

`alias` allows you to create a shortcut or abbreviation for a longer command.

IMPORTANT: aliases in Murex are not macros and are therefore different than
 other shells. if the shortcut requires any dynamics such as `piping`,
 `command sequencing`, `variable evaluations` or `scripting`...
 Prefer the **`function`** builtin.

## Usage

```
alias alias=command parameter parameter

!alias command
```

## Examples

Because aliases are parsed into an array of parameters, you cannot put the
entire alias within quotes. For example:

```
# bad :(
» alias hw="out Hello, World!"
» hw
exec "out\\ Hello,\\ World!": executable file not found in $PATH

# good :)
» alias hw=out "Hello, World!"
» hw
Hello, World!
```

Notice how only the command `out "Hello, World!"` is quoted in `alias` the
same way you would have done if you'd run that command "naked" in the command
line? This is how `alias` expects it's parameters and where `alias` on Murex
differs from `alias` in POSIX shells.

To materialize those differences, pay attention to the examples below:

```
# bad : the following statements generate errors,
#  prefer function builtin to implent them

» alias myalias=out "Hello, World!" | wc
» alias myalias=out $myvariable | wc
» alias myalias=out ${vmstat} | wc
» alias myalias=out "hello" && out "world"
» alias myalias=out "hello" ; out "world"
» alias myalias="out hello; out world"
```

In some ways this makes `alias` a little less flexible than it might
otherwise be. However the design of this is to keep `alias` focused on it's
core objective. To implement the above aliasing, you can use `function`
instead.

## Detail

### Allowed characters

Alias names can only include alpha-numeric characters, hyphen and underscore.
The following regex is used to validate the `alias`'s parameters:
`^([-_a-zA-Z0-9]+)=(.*?)$`

### Undefining an alias

Like all other definable states in Murex, you can delete an alias with the
bang prefix:

```
» alias hw=out "Hello, World!"
» hw
Hello, World!

» !alias hw
» hw
exec "hw": executable file not found in $PATH
```

### Order of preference

There is an order of precedence for which commands are looked up:

1. `runmode`: this is executed before the rest of the script. It is invoked by
   the pre-compiler forking process and is required to sit at the top of any
   scripts.

1. `test` and `pipe` functions also alter the behavior of the compiler and thus
   are executed ahead of any scripts.

4. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function directly from the interactive command line
   (however you can force an indirect call via `fexec`).

2. Aliases - defined via `alias`. All aliases are global.

3. Murex functions - defined via `function`. All functions are global.

5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.
   Also environmental variables too, declared via `export`.

6. globbing: however this only applies for commands executed in the interactive
   shell.

7. Murex builtins.

8. External executable files

You can override this order of precedence via the `fexec` and `exec` builtins.

## Synonyms

* `alias`
* `!alias`


## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Define Global (`global`)](../commands/global.md):
  Define a global variable and set it's value
* [Define Method Relationships (`method`)](../commands/method.md):
  Define a methods supported data-types
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Globbing (`g`)](../commands/g.md):
  Glob pattern matching for file system objects (eg `*.txt`)
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [`let`](../commands/let.md):
  Evaluate a mathematical function and assign to variable (deprecated)

<hr/>

This document was generated from [builtins/core/structs/function_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/function_doc.yaml).