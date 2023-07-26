# `fexec`  - Command Reference

> Execute a command or function, bypassing the usual order of precedence.

## Description

`fexec` allows you to execute a command or function, bypassing the usual order
of precedence.

## Usage

```
fexec: flag command [ parameters... ] -> <stdout>
```

## Examples

```
fexec: private /source/builtin/autocomplete.alias
```

## Flags

* `builtin`
    Execute a Murex builtin
* `function`
    Execute a Murex public function
* `help`
    Display help message
* `private`
    Execute a Murex private function

## Detail

### Order of precedence

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

### Compatibility with POSIX

For compatibility with traditional shells like Bash and Zsh, `builtin` is an
alias to `fexec builtin`

## Synonyms

* `fexec`
* `builtin`


## See Also

* [`alias`](../commands/alias.md):
  Create an alias for a command
* [`autocomplete`](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`bg`](../commands/bg.md):
  Run processes in the background
* [`builtins`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [`event`](../commands/event.md):
  Event driven programming for shell scripts
* [`exec`](../commands/exec.md):
  Runs an executable
* [`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [`function`](../commands/function.md):
  Define a function block
* [`jobs`](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [`open`](../commands/open.md):
  Open a file with a preferred handler
* [`private`](../commands/private.md):
  Define a private function block
* [`source` ](../commands/source.md):
  Import Murex code from another file of code block