# _murex_ Shell Docs

## Command Reference: `murex-package`

> _murex_'s package manager

### Description

_murex_ comes with it's own package manager to make managing plugins easier.

### Usage

Install a new package

    murex-package: install uri
    
Update all packages

    murex-package: update
    
Enable a package or module which had been disabled

    murex-package: enable package
    
    murex-package: enable package/module
    
Disable a package

    murex-package: disable package
    
    murex-package: disable package/module
    
Import packages from another package database

    murex-package: import [ uri/ | local/path/ ]packages.json
    
Check status of murex packages

    murex-package: status -> <stdout>

### Flags

* `disable`
    Disables a previously enabled package or module
* `enable`
    Enables a previously disabled package or module
* `import`
    Import packages described in a backup package DB from user defined URI or local path
* `install`
    Installs a package from a user defined URI
* `status`
    Returns the version status of locally installed packages
* `update`
    Updates all installed packages

### See Also

* [user-guide/_murex_ profile files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [commands/`alias`](../commands/alias.md):
  Create an alias for a command
* [commands/`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* [commands/`function`](../commands/function.md):
  Define a function block
* [commands/`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [commands/`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [user-guide/modules](../user-guide/modules.md):
  