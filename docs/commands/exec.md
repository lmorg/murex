# Execute External Command (`exec`)

> Runs an executable

## Description

With Murex, like most other shells, you launch a process by calling the
name of that executable directly. While this is suitable 99% of the time,
occasionally you might run into an edge case where that wouldn't work. The
primary reason being if you needed to launch a process from a variable, eg

```
» set exe=uname
» $exe
uname
```

As you can see here, Murex's behavior here is to output the contents of
the variable rather then executing the contents of the variable. This is
done for safety reasons, however if you wanted to override that behavior
then you could prefix the variable with exec:

```
» set exe=uname
» exec $exe
Linux
```

## Usage

```
<stdin> -> exec

<stdin> -> exec -> <stdout>

           exec -> <stdout>
```

## Examples

### As a function

```
» exec printf "Hello, world!"
Hello, world!
```

### Working around aliases

If you have an alias like `alias ls=ls --color=auto` and you wanted to run `ls`
but without colour, you might run `exec ls`.

## Detail

If any command doesn't exist as a builtin, function nor alias, then Murex
will default to forking out to any command with this name (subject to an
absolute path or the order of precedence in `$PATH`). Any forked process will
show up in both the operating systems process viewer (eg `ps`) but also
Murex's own process viewer, `fid-list`. However inside `fid-list` you will
notice that all external processes are listed as `exec` with the process name
as part of `exec`'s parameters. That is because `exec` is handler for programs
that aren't native to Murex.

### Compatibility with POSIX

For compatibility with traditional shells like Bash and Zsh, `command` is an
alias for `exec`.

## Synonyms

* `exec`
* `command`
* `exec.file`


## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Define Variable (`set`)](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Kill All In Session (`fid-killall`)](../commands/fid-killall.md):
  Terminate all running Murex functions in current session
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function
* [Re-Scan $PATH For Executables](../commands/murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
* [Shell Runtime (`builtins`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex

<hr/>

This document was generated from [builtins/core/typemgmt/types_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/typemgmt/types_doc.yaml).