# _murex_ Shell Docs

## Command Reference: `fid-list`

> Lists all running functions within the current _murex_ session

## Description

`fid-list` is a tool for outputting all the functions currently managed by that
_murex_ session. Those functions could be _murex_ functions, builtins or any
external executables launched from _murex_.

Conceptually `fid-list` is a little like `ps` (on POSIX systems) however
`fid-list` was not written to be POSIX compliant.

Multiple flags cannot be used with each other.

## Usage

    fid-list [ flag ] -> <stdout>
    
`jobs` is an alias for `fid-list: --jobs`:
    jobs -> <stdout>

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

Because _murex_ is a multi-threaded shell, builtins are not forked processes
like in a traditional / POSIX shell. This means that you cannot use the
operating systems default process viewer (eg `ps`) to list _murex_ functions.
This is where `fid-list` comes into play. It is used to view all the functions
and processes that are managed by the current _murex_ session. That would
include:
* any aliases within _murex_
* public and private _murex_ functions
* builtins (eg `fid-list` is a builtin command)
* any external processes that were launched from within this shell session
* any background functions or processes of any of the above

## Synonyms

* `fid-list`
* `jobs`


## See Also

* [types/`*` (generic) ](../types/generic.md):
  generic (primitive)
* [commands/`bexists`](../commands/bexists.md):
  Check which builtins exist
* [commands/`bg`](../commands/bg.md):
  Run processes in the background
* [commands/`builtins`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [types/`csv` ](../types/csv.md):
  CSV files (and other character delimited tables)
* [commands/`exec`](../commands/exec.md):
  Runs an executable
* [commands/`fexec` ](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [commands/`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [commands/`fid-kill`](../commands/fid-kill.md):
  Terminate a running _murex_ function
* [commands/`fid-killall`](../commands/fid-killall.md):
  Terminate _all_ running _murex_ functions
* [commands/`jobs`](../commands/fid-list.md):
  Lists all running functions within the current _murex_ session
* [types/`jsonl` ](../types/jsonl.md):
  JSON Lines (primitive)
* [commands/`murex-update-exe-list`](../commands/murex-update-exe-list.md):
  Forces _murex_ to rescan $PATH looking for exectables