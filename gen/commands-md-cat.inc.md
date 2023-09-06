This section is a glossary of Murex builtin commands.

Because Murex is loosely modelled on the functional paradigm, it means
all language constructs are exposed via functions and those are typically
builtins because they can share the Murex runtime virtual machine.
However any executable command can also be called from within Murex;
be that either via the `exec` builtin or natively like you would from any
Linux, UNIX, or even Windows command prompt.

## Other Reference Material

### Language Guides

1. [Language Tour]({{if env "DOCGEN_TARGET="}}/docs{{end}}/tour.md), which is an introduction into
    the Murex language.

2. [Rosetta Stone]({{if env "DOCGEN_TARGET="}}/docs{{end}}/user-guide/rosetta-stone.md), which is a reference
    table comparing Bash syntax to Murex's.

### Murex's Source Code

The source for each of these builtins can be found on [Github](https://github.com/lmorg/murex/tree/master/builtins/core).

### Shell Commands For Querying Builtins

From the shell itself: run `builtins` to list the builtin command.

If you require a manual on any of those commands, you can run `murex-docs`
to return the same markdown-formatted document as those listed below. eg

```
murex-docs trypipe
```