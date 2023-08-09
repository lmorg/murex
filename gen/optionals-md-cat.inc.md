This section is a glossary of Murex optional builtins.

These builtins likely wont be compiled with Murex unless you specifically
request them to be. This might be because they duplicate functionality
already available on POSIX systems or introduce more complex dependencies.
It might also be because that specific builtin is in an alpha stage and thus
not ready to ship with Murex.

## Other Reference Material

### Language Guides

1. [Core Builtins]({{if env "DOCGEN_TARGET="}}/docs{{end}}/commands/), for docs
    on the core builtins.

2. [Language Tour]({{if env "DOCGEN_TARGET="}}/docs{{end}}/tour.md), which is an introduction into
    the Murex language.

3. [Rosetta Stone]({{if env "DOCGEN_TARGET="}}/docs{{end}}/user-guide/rosetta-stone.md), which is a reference
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