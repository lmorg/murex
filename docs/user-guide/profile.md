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

## Overriding The Default Paths (XDG)

Some individuals, particularly those running Linux, follow a standard called
[XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html).
While Murex does not adhere to this standard, instead conforming to the
_de facto_ standard defined by the past precedents of previous shells, in
order to offer flexibility for those who do prefer the XDG specification
Murex does support overriding its own default paths via special environmental
variables.

- `MUREX_PRELOAD` defines the preload path (and file name)
- `MUREX_MODULES` defines the module path (only)
- `MUREX_PROFILE` defines the profile path (and file name)

Where `MUREX_PRELOAD` and/or `MUREX_PROFILE` are directories rather than
absolute file names, the path is appended with the default file names as
named above.

For people who wish to use XDG paths, in many instances you can get away
with setting the follow prior to launching Murex (eg in `/etc/profile.d`):

```
MUREX_PRELOAD="$XDG_CONFIG_HOME/murex/"
MUREX_MODULES="$XDG_CONFIG_HOME/murex/"
MUREX_PROFILE="$XDG_CONFIG_HOME/murex/"
```

This, however, depends on `$XDG_CONFIG_HOME` pointing to a single path rather
than an array of paths (like `$PATH`). In that instance you can still use
custom paths in Murex but you might need to get a little more creative in
how you define that value.

## See Also

* [Define Environmental Variable (`export`)](../commands/export.md):
  Define an environmental variable and set it's value
* [Modules And Packages](../user-guide/modules.md):
  An introduction to Murex modules and packages
* [Murex Package Management (`murex-package`)](../commands/murex-package.md):
  Murex's package manager

<hr/>

This document was generated from [gen/user-guide/profile_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/profile_doc.yaml).