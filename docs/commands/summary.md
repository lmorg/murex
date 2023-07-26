# `summary`

> Defines a summary help text for a command

## Description

`summary` define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.

Normally this text is populated from the `man` pages or `murex-docs`, however
if neither exist or if you wish to override their text, then you can use
`summary` to define that text.

## Usage

Define a commands summary

```
summary command description
```

Undefine a summary

```
!summary command
```

## Examples

Define a commands summary

```
» summary: foobar "Hello, world!"
» runtime: --summaries -> [ foobar ]
Hello, world!
```

Undefine a summary

```
» !summary: foobar
```

## Synonyms

- `summary`
- `!summary`

## See Also

- [`bexists`](./bexists.md):
  Check which builtins exist
- [`builtins`](./runtime.md):
  Returns runtime information on the internal state of Murex
- [`config`](./config.md):
  Query or define Murex runtime settings
- [`exec`](./exec.md):
  Runs an executable
- [`fid-list`](./fid-list.md):
  Lists all running functions within the current Murex session
- [`murex-docs`](./murex-docs.md):
  Displays the man pages for Murex builtins
- [`murex-update-exe-list`](./murex-update-exe-list.md):
  Forces Murex to rescan $PATH looking for executables
- [`runtime`](./runtime.md):
  Returns runtime information on the internal state of Murex
