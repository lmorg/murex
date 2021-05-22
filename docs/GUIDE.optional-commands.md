# _murex_ Shell Docs

## Optional Command Reference

This section is a glossary of _murex_ optional builtins.

These builtins likely wont be compiled with _murex_ unless you specifically
request them to be. This might be because they duplicate functionality
already available on POSIX systems or introduce more complex dependencies.
It might also be because that specific builtin is in an alpha stage and thus
not ready to ship with _murex_.

There are also optional some builtins which are automatically be built on
Windows to include a few coreutils commands missing Windows.

## Other Reference Material

### Language Guides

1. [GUIDE.builtin-functions](./GUIDE.builtin-functions.md), for docs
on the core builtins.

2. [GUIDE.control-structures](./GUIDE.control-structures.md), which
contains builtins required for building logic.

### _murex_'s Source Code

In _murex_'s source under the `lang/builtins` path of the project files
is several directories, each hosting different categories of _murex_
builtins. From core commands through to data-types and methods.

Each package will include a README.md file with a basic summary of what
that package is used for and all you to enable or disable builtins, should
you decide to compile the shell from source.

### Shell Commands For Querying Builtins

From the shell itself: run `builtins` to list the builtin command.

If you require a manual on any of those commands, you can run `murex-docs`
to return the same markdown-formatted document as those listed below. eg

    murex-docs trypipe

## Pages

* [`!bz2` ](optional/bz2.md):
  Decompress a bz2 file
* [`base64` ](optional/base64.md):
  Encode or decode a base64 string
* [`gz` ](optional/gz.md):
  Compress or decompress a gzip file
* [`qr` ](optional/qr.md):
  Creates a QR code from STDIN
* [`select` ](optional/select.md):
  Inlining SQL into shell pipelines
* [`sleep` ](optional/sleep.md):
  Suspends the shell for a number of seconds