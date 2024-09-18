# Special Variables

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
    outside of the current function. Thus one function cannot read nor write to
    a variable in another function.

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

## Pages

* [Numeric (str)](../variables/numeric.md):
  Variables who's name is a positive integer, eg `0`, `1`, `2`, `3` and above
* [`$.`, Meta Values (json)](../variables/meta-values.md):
  State information for iteration blocks
* [`ARGV` (json)](../variables/argv.md):
  Array of the command name and parameters within a given scope
* [`COLUMNS` (int)](../variables/columns.md):
  Character width of terminal
* [`EVENT_RETURN` (json)](../variables/event_return.md):
  Return values for events
* [`HOME` (path)](../variables/home.md):
  Return the home directory for the current session user
* [`HOSTNAME` (str)](../variables/hostname.md):
  Hostname of the current machine
* [`LINES` (int)](../variables/lines.md):
  Character height of terminal
* [`LOGNAME` (str)](../variables/logname.md):
  Username for the current session (historic)
* [`MUREX_ARGV` (json)](../variables/murex_argv.md):
  Array of the command name and parameters passed to the current shell
* [`MUREX_EXE` (path)](../variables/murex_exe.md):
  Absolute path to running shell
* [`OLDPWD` (path)](../variables/oldpwd.md):
  Return the home directory for the current session user
* [`PARAMS` (json)](../variables/params.md):
  Array of the parameters within a given scope
* [`PWDHIST` (json)](../variables/pwdhist.md):
  History of each change to the sessions working directory
* [`PWD` (path)](../variables/pwd.md):
  Current working directory
* [`RANDOM` (int)](../variables/random.md):
  Return a random 32-bit integer (historical)
* [`SELF` (json)](../variables/self.md):
  Meta information about the running scope.
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`TMPDIR` (path)](../variables/tmpdir.md):
  Return the temporary directory
* [`USER` (str)](../variables/user.md):
  Username for the current session