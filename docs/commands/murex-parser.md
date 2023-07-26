# `murex-parser`  - Command Reference

> Runs the Murex parser against a block of code 

## Description

`summary` define help text for a command. This is effectively like a tooltip
message that appears, by default, in blue in the interactive shell.

Normally this text is populated from the `man` pages or `murex-docs`, however
if neither exist or if you wish to override their text, then you can use
`summary` to define that text.

## Usage

```
<stdin> -> murex-parser -> <stdout>

murex-parser { code-block } -> <stdout>
```

## Detail

Please note this command is still very much in beta and is likely to change in incompatible ways in the future. If you do happen to like this command and/or have any suggestions on how to improve it, then please leave your feedback on the GitHub repository, https://github.com/lmorg/murex

## See Also

* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of Murex