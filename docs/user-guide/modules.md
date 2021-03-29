# _murex_ Shell Docs

## User Guide: Modules and packages

> An introduction to _murex_ modules and packages

_murex_ has it's own module system with namespacing and a package manager. But
why should a shell need all this?

The answer comes from years of me using Bash and wishing my Bash environment
could be consistent across multiple machines. So this document is authored from
the perspective of my personal usage ("me" being Laurence Morgan, the original
author of _murex_).

What _murex_'s package system provides is:
1. A way to ensure consistency across multiple platforms
2. An easy way to extend _murex_
3. An easy way to share what you've extended with others
4. An easy way to ensure your extensions are kept up-to-date
5. An easy way to track what code is running in your shell and from where it
   was loaded

Before I address those points in more detail, a bit of background into what
modules and packages are:

#### What Are Packages And Modules?

_murex_ comes with it's own package manager to make managing plugins easier.

The format of the packages is a directory, typically located at `~/.murex_modules`,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.

The way packages and modules are represented is as a path:
    
    package/module
        
`murex-package` is a package management tool for administrating murex modules
and packages.

### Using Packages And Modules

#### Consistency

Package database are stored locally at `~/.murex_modules/packages.json`. This
file is portable so any new machine can have `packages.json` imported. The
easiest way of doing this is using `murex-package` which can import from a
local path or HTTP(S) URI and automatically download any packages described in
the database.

For example the command I run on any new dev machine to import all of my DevOps
tools and terminal preferences is the following:

    murex-package: import https://gist.githubusercontent.com/lmorg/770c71786935b44ba6667eaa9d470888/raw/fb7b79d592672d90ecb733944e144d722f77fdee/packages.json
    
#### Extendability

Namespacing allows for `private` functions which allows you to write smaller
functions. Smaller functions are easier to write tests against (_murex_ also
has an inbuilt testing and debugging tools).

#### Sharing Code

Packages can be hosted via HTTP(S) or git. Anyone can import anyone elses
packages using `murex-package`. 

    murex-package: install https://github.com/lmorg/murex-module-murex-dev.git
    
#### Updating Packages

Updating packages is easy:

    murex-package: update
    
#### Tracking Code

All code loaded in _murex_, every function, variable and event (etc) is stored
in memory with metadata about where it was loaded from; which package, file and
at what time. This is called `FileRef`.

For more information on FileRef see the link below.

## See Also

* [user-guide/Modules and packages](../user-guide/fileref.md):
  How to track what code was loaded and from where
* [commands/`murex-package`](../commands/murex-package.md):
  _murex_'s package manager
* [commands/`private`](../commands/private.md):
  Define a private function block
* [commands/`test`](../commands/test.md):
  _murex_'s test framework - define tests, run tests and debug shell scripts