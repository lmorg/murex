### Module Strings For Non-Module Code

#### Source

A common shell idiom is to load shell script files via `source` / `.`. When
this is done the module string (as seen in the `FileRef` structures described
above) will be `source/hash` where **hash** will be a unique hash of the file
path and load time.

Thus no two sourced files will share the same module string. Even the same file
but modified and sourced twice (before and after the edit) will have different
module strings due to the load time being part of the hashed data.

#### REPL

Any functions, variables, events, auto-completions, etc created manually,
directly, in the interactive shell will have a module string of `murex` and an
empty Filename string.