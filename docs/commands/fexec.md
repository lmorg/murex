# Execute Shell Function or Builtin (`fexec`)

> Execute a command or function, bypassing the usual order of precedence.

## Description

`fexec` allows you to execute a command or function, bypassing the usual order
of precedence.

## Usage

```
fexec flag command [ parameters... ] -> <stdout>
```

## Examples

```
fexec private /source/builtin/autocomplete.alias
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
* `exec.builtin`
* `exec.function`
* `exec.private`


## See Also

* [Alias Pointer (`alias`)](../commands/alias.md):
  Create an alias for a command
* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Public Function (`function`)](../commands/function.md):
  Define a function block
* [Shell Runtime (`builtins`)](../commands/runtime.md):
  Returns runtime information on the internal state of Murex
* [Tab Autocompletion (`autocomplete`)](../commands/autocomplete.md):
  Set definitions for tab-completion in the command line
* [`event`](../commands/event.md):
  Event driven programming for shell scripts

<hr/>

This document was generated from [builtins/core/management/fexec_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/fexec_doc.yaml).