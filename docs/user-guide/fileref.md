# _murex_ Shell Docs

## User Guide: FileRef

> How to track what code was loaded and from where

Every function, event, autocompletion and even variable is stored with which
file it was sourced, when it was loaded and what module it was loaded from.
This makes it trivial to identify buggy 3rd party code or even malicious
libraries.

    » runtime: --functions -> [[ /agent/FileRef/ ]]
    {
        "Column": 5,
        "Line": 5,
        "Source": {
            "DateTime": "2021-03-28T09:10:53.572197+01:00",
            "Filename": "/Users/laurencemorgan/.murex_modules/murex-dev/murex-dev.mx",
            "Module": "murex-dev/murex-dev"
        }
    }
    
    » runtime --globals -> [[ /DEVOPSBIN/FileRef ]]
    {
        "Column": 1,
        "Line": 0,
        "Source": {
            "DateTime": "2021-03-28T09:10:53.541952+01:00",
            "Filename": "/Users/laurencemorgan/.murex_modules/devops/global.mx",
            "Module": "devops/global"
        }
    }
    
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

## See Also

* [user-guide/Modules and packages](../user-guide/modules.md):
  An introduction to _murex_ modules and packages
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`murex-package`](../commands/murex-package.md):
  _murex_'s package manager
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block