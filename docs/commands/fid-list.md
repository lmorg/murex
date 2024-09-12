# Display Running Functions (`fid-list`)

> Lists all running functions within the current Murex session

## Description

`fid-list` is a tool for outputting all the functions currently managed by that
Murex session. Those functions could be Murex functions, builtins or any
external executables launched from Murex.

Conceptually `fid-list` is a little like `ps` (on POSIX systems) however
`fid-list` was not written to be POSIX compliant.

Multiple flags cannot be used with each other.

## Usage

```
fid-list [ flag ] -> <stdout>
```

`jobs` is an alias for `fid-list: --jobs`:
```
jobs -> <stdout>
```

## Flags

* `--background`
    Returns a `json` map of background jobs
* `--csv`
    Output table in a `csv` format
* `--help`
    Outputs a list of parameters and a descriptions
* `--jobs`
    Show background and stopped jobs
* `--jsonl`
    Output table in a jsonlines (`jsonl`) format (defaulted to when piped)
* `--stopped`
    Returns a `json` map of stopped jobs
* `--tty`
    Force default TTY output even when piped

## Detail

Because Murex is a multi-threaded shell, builtins are not forked processes
like in a traditional / POSIX shell. This means that you cannot use the
operating systems default process viewer (eg `ps`) to list Murex functions.
This is where `fid-list` comes into play. It is used to view all the functions
and processes that are managed by the current Murex session. That would
include:

* any aliases within Murex
* public and private Murex functions
* builtins (eg `fid-list` is a builtin command)
* any external processes that were launched from within this shell session
* any background functions or processes of any of the above

## Synonyms

* `fid-list`
* `jobs`


## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Check Builtin Exists (`bexists`)](../commands/bexists.md):
  Check which builtins exist
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
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
* [`*` (generic)](../types/generic.md):
  generic (primitive)
* [`csv`](../types/csv.md):
  CSV files (and other character delimited tables)
* [`jsonl`](../types/jsonl.md):
  JSON Lines

<hr/>

This document was generated from [builtins/core/processes/fid-list_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/processes/fid-list_doc.yaml).