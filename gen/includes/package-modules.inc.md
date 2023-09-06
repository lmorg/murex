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
| https://github.com/lmorg/murex-module-jump                   | Adds `j` function which calls `jump` and pupulates the suggestion(s) as autocompelete suggestions Adds autocomplete suggestions for `jump` executable Adds event to watch for directory changes |
| https://github.com/orefalo/murex-module-starship             | starship - The minimal, blazing-fast, and infinitely customizable prompt |
| and [many more](https://github.com/search?q=murex-module-&type=repositories) | Murex modules typically follow the `murex-module-*` naming convention |

