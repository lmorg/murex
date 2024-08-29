# `type`

> Command type (function, builtin, alias, etc)

## Description

`type` returns information about the type of the command. This is a POSIX
requirement and not to be confused with Murex data types. 

## Usage

```
type command -> <stdout>
```

## Examples

### TTY output

```
» type murex-docs
`murex-docs` is a shell function:

    # Wrapper around builtin to pipe to less

    config: set proc strict-arrays false
    fexec: builtin murex-docs @PARAMS | less
```

### Piped output

```
» type murex-docs -> cat
function
```

## Detail

There are a few different types of commands:

### alias

This will be represented in `which` and `type` by the term **alias** and, when
stdout is a TTY, `which` will follow the alias to print what command the alias
points to.

### function

This is a Murex function (defined via `function`) and will be represented in
`which` and `type` by the term **function**.

### private

This is a private function (defined via `private`) and will be represented in
`which` and `type` by the term **private**.

### builtin

This is a shell builtin, like `out` and `exit`. It will be represented in
`which` and `type` by the term **builtin**.

### external executable

This is any other external command, such as `systemctl` and `python`. This
will be represented in `which` by the path to the executable. When stdout is a
TTY, `which` will also print the destination path of any symlinks too.

In `type`, it is represented by the term **executable**.

## See Also

* [`exit`](../commands/exit.md):
  Exit murex
* [`which`](../commands/which.md):
  Locate command origin
* [exec.* (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [exec.file: `exec`](../commands/exec.md):
  Runs an executable
* [io.out](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [shell.alias](../commands/alias.md):
  Create an alias for a command
* [shell.function](../commands/function.md):
  Define a function block

<hr/>

This document was generated from [builtins/core/management/type_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/type_doc.yaml).