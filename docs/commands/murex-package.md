# _murex_ Shell Guide

## Command Reference: `murex-package`

> _murex_'s package manager

### Description

_murex_ comes with it's own package manager to make managing plugins easier.

### Usage

    murex-package install uri
    
    murex-package update
    
    murex-package enable package
    murex-package enable package/module
    
    murex-package disable package
    murex-package disable package/module
    
    murex-package import [ uri | local/path ] packages.json
    
    murex-package status -> <stdout>

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

* commands/[`alias`](../commands/alias.md):
  Create an alias for a command
* commands/[`config`](../commands/config.md):
  Query or define _murex_ runtime settings
* commands/[`function`](../commands/function.md):
  Define a function block
* commands/[`private`](../commands/private.md):
  Define a private function block
* commands/[`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* commands/[murex-man](../commands/murex-man.md):
  
* commands/[status](../commands/status.md):
  