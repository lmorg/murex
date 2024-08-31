# Optional Builtins Reference

This section is a glossary of Murex optional builtins.

These builtins likely wont be compiled with Murex unless you specifically
request them to be. This might be because they duplicate functionality
already available on POSIX systems or introduce more complex dependencies.
It might also be because that specific builtin is in an alpha stage and thus
not ready to ship with Murex.

## Other Reference Material

### Language Guides

1. [Core Builtins](/docs/commands/), for docs
    on the core builtins.

2. [Language Tour](/docs/tour.md), which is an introduction into
    the Murex language.

3. [Rosetta Stone](/docs/user-guide/rosetta-stone.md), which is a reference
    table comparing Bash syntax to Murex's.

### Murex's Source Code

The source for each of these builtins can be found on [Github](https://github.com/lmorg/murex/tree/master/builtins/optional).

### Shell Commands For Querying Builtins

From the shell itself: run `builtins` to list the builtin command.

If you require a manual on any of those commands, you can run `murex-docs`
to return the same markdown-formatted document as those listed below. eg

```
murex-docs trypipe
```

## Pages

* [Inline SQL (`select`)](../optional/select.md):
  Inlining SQL into shell pipelines
* [`!bz2`](../optional/bz2.md):
  Decompress a bz2 file
* [`base64` ](../optional/base64.md):
  Encode or decode a base64 string
* [`gz`](../optional/gz.md):
  Compress or decompress a gzip file
* [`qr`](../optional/qr.md):
  Creates a QR code from stdin
* [`sleep`](../optional/sleep.md):
  Suspends the shell for a number of seconds