# Profile Files

> A breakdown of the different files loaded on start up

## Description

Murex has several profile files which are loaded in the following order of
execution:

1. `~/.murex_preload`
2. `~/.murex_modules/*/`
3. `~/.murex_profile`

### `.murex_preload`

This file should only used to define any environmental variables that might
need to be set before the modules are loaded (eg including directories in
`$PATH` if you have anything installed in non-standard locations).

Most of the time this file will be empty bar the standard warning message:

    # This file is loaded before any murex modules. It should only contain
    # environmental variables required for the modules to work eg:
    #
    #     export PATH=...
    #
    # Any other profile config belongs in your profile script instead:
    # /home/$USER/.murex_profile

This file is created upon the first run of Murex.

### `.murex_modules/`

Murex's module directory - where all the modules are installed
to. This directory is managed by `murex-package` builtin.

### `.murex_profile`

This file is comparable to `.bash_profile`, `.bashrc` and `.zshrc` etc. It
is the standard place to put all user and/or machine specific config in.

`.murex_profile` is only read from the users home directory. Unlike bash et
al, profiles will not be read from `/etc/profile.d` nor similar. Modules
should be used in its place.

## Overriding The Default Paths

Each of the config paths can be manually overridden:

### Specific Entities

- `MUREX_PRELOAD` defines the preload path and file name
- `MUREX_MODULES` defines the module path
- `MUREX_PROFILE` defines the profile path and file name
- `MUREX_HISTORY` defines the history path and file name

Where `MUREX_PRELOAD`, `MUREX_PROFILE` and/or `MUREX_HISTORY` are directories
rather than absolute file names, the path is appended with the default file
names. ie the file name listed above.

### Root Config Directory

You can also specify one directory for all the above entities via the following
environmental variable:

- `MUREX_CONFIG_DIR`

This path can be superseded for specific entities by using (for example)
`MUREX_PRELOAD`. This would mean those entities would follow their specific
environmental variable path while unnamed entities would fall back to
`MUREX_CONFIG_DIR`.

If the path does not exist, then it is created automatically by Murex.

### XDG 

Some individuals, particularly those running Linux, follow a standard called
[XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html).
While Murex does not adhere to this standard, instead conforming to the
_de facto_ standard defined by the past precedents of previous shells.

For people who wish to use XDG paths, in many instances you can get away
with setting the follow prior to launching Murex (eg in `/etc/profile.d`):

```
export MUREX_CONFIG_DIR="$XDG_CONFIG_HOME/murex/"
```

This, however, depends on `$XDG_CONFIG_HOME` pointing to a single path rather
than an array of paths (like `$PATH`). In that instance you can still use
custom paths in Murex but you might need to get a little more creative in
how you define that value.

## See Also

* [Command Line History (`history`)](../commands/history.md):
  Outputs murex's command history
* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Modules And Packages](../user-guide/modules.md):
  Modules and packages: An Introduction
* [Murex Package Management (`murex-package`)](../commands/murex-package.md):
  Murex's package manager

<hr/>

This document was generated from [gen/user-guide/profile_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/profile_doc.yaml).