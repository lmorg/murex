# Modules And Packages

> An introduction to Murex modules and packages

## Description

Murex has it's own module system with namespacing and a package manager. But
why should a shell need all this?

The answer comes from years of me using Bash and wishing my Bash environment
could be consistent across multiple machines. So this document is authored from
the perspective of my personal usage ("me" being Laurence Morgan, the original
author of Murex).

What Murex's package system provides is:

1. A way to ensure consistency across multiple platforms
2. An easy way to extend Murex
3. An easy way to share what you've extended with others
4. An easy way to ensure your extensions are kept up-to-date
5. An easy way to track what code is running in your shell and from where it
   was loaded

Before I address those points in more detail, a bit of background into what
modules and packages are:

### What Are Packages And Modules?

Murex comes with it's own package manager to make managing plugins easier.

The format of the packages is a directory, typically located at `~/.murex_modules`,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.

The way packages and modules are represented is as a path:
    package/module

`murex-package` is a package management tool for administrating murex modules
and packages.

| Name                                                         | Summary                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| https://github.com/lmorg/murex-module-jump                   | Enables autojump functionalities                             |
| https://github.com/orefalo/murex-module-starship             | starship - The minimal, blazing-fast, and infinitely customizable prompt |
| and [many more](https://github.com/search?q=murex-module-&type=repositories) | Murex modules typically follow the `murex-module-*` naming convention |

## Using Packages And Modules

### Consistency

Package database are stored locally at `~/.murex_modules/packages.json`. This
file is portable so any new machine can have `packages.json` imported. The
easiest way of doing this is using `murex-package` which can import from a
local path or HTTP(S) URI and automatically download any packages described in
the database.

For example the command I run on any new dev machine to import all of my DevOps
tools and terminal preferences is the following:

```
murex-package import https://gist.githubusercontent.com/lmorg/770c71786935b44ba6667eaa9d470888/raw/fb7b79d592672d90ecb733944e144d722f77fdee/packages.json
```

### Extendability

Namespacing allows for `private` functions which allows you to write smaller
functions. Smaller functions are easier to write tests against (Murex also
has an inbuilt testing and debugging tools).

### Sharing Code

Packages can be hosted via HTTP(S) or git. Anyone can import anyone elses
packages using `murex-package`. 

```
murex-package install https://github.com/lmorg/murex-module-murex-dev.git
```

### Updating Packages

Updating packages is easy:

```
murex-package update
```

### Tracking Code

All code loaded in Murex, every function, variable and event (etc) is stored
in memory with metadata about where it was loaded from; which package, file and
at what time. This is called `FileRef`.

For more information on `FileRef` see the link below.

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

* [FileRef](../user-guide/fileref.md):
  How to track what code was loaded and from where
* [Include / Evaluate Murex Code (`source`)](../commands/source.md):
  Import Murex code from another file or code block
* [Murex Package Management (`murex-package`)](../commands/murex-package.md):
  Murex's package manager
* [Private Function (`private`)](../commands/private.md):
  Define a private function block
* [Shell Script Tests (`test`)](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts

<hr/>

This document was generated from [gen/user-guide/modules_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/modules_doc.yaml).