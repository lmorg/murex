# Location Of Command (`which`)

> Locate command origin

## Description

`which` locates a command's origin. If stdout is a TTY, then it's output will be
human readable. If stdout is a pipe then it's output will be a simple list.

`which` can take multiple parameters, each representing a different command you
want looked up.

## Usage

```
which command... -> <stdout>
```

## Examples

### TTY output

```
» which cat dog jobs git dug
cat => (/bin/cat) cat - concatenate and print files
dog => unknown
jobs => (alias) fid-list --jobs => (builtin) Lists all running functions within the current Murex session
git => (/opt/homebrew/bin/git -> ../Cellar/git/2.41.0/bin/git) git - the stupid content tracker
dug => (murex function) A bit like dig but which outputs JSON
```

### Piped output

```
» which cat dog jobs git dug -> cat
/bin/cat
unknown
alias
/opt/homebrew/bin/git
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

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Display Command Type (`type`)](../commands/type.md):
  Command type (function, builtin, alias, etc)
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Exit Murex (`exit`)](../commands/exit.md):
  Exit murex
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Public Function (`function`)](../commands/function.md):
  Define a function block

<hr/>

This document was generated from [builtins/core/management/which_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/which_doc.yaml).