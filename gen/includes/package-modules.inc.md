Murex comes with it's own package manager to make managing plugins easier.

The format of the packages is a directory, typically located at `~/.murex_modules`,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.

The way packages and modules are represented is as a path:
    
```
package/module
```
    
`murex-package` is a package management tool for administrating murex modules
and packages.

You can find existing modules on [GitHub](https://github.com/search?q=murex-module-&type=repositories), they typically follow a `murex-module-*` naming convention.

