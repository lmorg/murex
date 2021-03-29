# _murex_ Shell Docs

## User Guide: FileRef

> How to track what code was loaded and from where

Every function, event, autocompletion and even variable is stored with which
file it was sourced, when it was loaded and what module it was loaded from.
This makes it trivial to identify buggy 3rd party code or even malicious
libraries.

    » runtime: --functions -> [[/agent/FileRef/]]
    {
        "Column": 5,
        "Line": 5,
        "Source": {
            "DateTime": "2021-03-28T09:10:53.572197+01:00",
            "Filename": "/Users/laurencemorgan/.murex_modules/murex-dev/murex-dev.mx",
            "Module": "murex-dev/murex-dev"
        }
    }
    
    » runtime --globals -> [[/DEVOPSBIN/FileRef]]
    {
        "Column": 1,
        "Line": 0,
        "Source": {
            "DateTime": "2021-03-28T09:10:53.541952+01:00",
            "Filename": "/Users/laurencemorgan/.murex_modules/devops/global.mx",
            "Module": "devops/global"
        }
    }

## See Also

* [user-guide/Modules and packages](../user-guide/modules.md):
  An introduction to _murex_ modules and packages
* [commands/`[[` (element)](../commands/element.md):
  Outputs an element from a nested structure
* [commands/`murex-package`](../commands/murex-package.md):
  _murex_'s package manager
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_