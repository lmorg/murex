# Define Function Arguments: `args`

> Command line flag parser for Murex shell scripting

## Description

One of the nuisances of shell scripts is handling flags. More often than not
your script will be littered with `$1` still variables and not handle flags
shifting in placement amongst a group of parameters. `args` aims to fix that by
providing a common tool for parsing flags.

`args` takes a name of a variable to assign the result of the parsed parameters
as well as a JSON structure containing the result. It also returns a non-zero
exit number if there is an error when parsing.

## Usage

```
args var-name { json-block } -> <stdout>
```

## Examples

```
#!/usr/bin/env murex

# First we define what parameters to accept:
# Pass the `args` function a JSON string (because JSON objects share the same braces as murex block, you can enter JSON
# directly as unescaped values as parameters in murex).
#
# --str: str == string data type
# --num: num == numeric data type
# --bool: bool == flag used == true, missing == false
# -b: --bool == alias of --bool flag
args args %{
    AllowAdditional: true
    Flags: {
        --str:  str
        --num:  num
        --bool: bool
        -b: --bool
    }
}
catch {
    # Lets check for errors in the command line parameters. If they exist then
    # print the error and then exit.
    err $args.error
    exit 1
}

out "The structure of \$args is: ${$args->pretty}\n\n"


# Some example usage:
# -------------------

!if { $(args.Flags.--bool) } {
    out "Flag `--bool` was not set."
}

# `<!null>` redirects the STDERR to a named pipe. In this instance it's the 'null' pipe so equivalent to 2>/dev/null
# thus we are just suppressing any error messages.
try <!null> {
    $(args.Flags.--str) -> set fStr
    $(args.Flags.--num) -> set fNum

    out "Defined Flags:"
    out "  --str == $(fStr)"
    out "  --num == $(fNum)"
}

catch {
    err "Missing `--str` and/or `--num` flags."
}

$args[Additional] -> foreach flag {
    out "Additional argument (ie not assigned to a flag): `$(flag)`."
}
```

## Detail

### Flags vs Parameters

#### Flags

Flags are any values pass via `-` or `--` prefixed labels. For example
`--datatype str` would assign the value `str` to the flag name `--datatype`.

Flags can be `str`, `int` `num` `bool`. These values will be type checked and
`args` will return an error if a user passes (for example) alpha character to
a numeric flag.

Boolean flags do not require `true` nor `false` values to be included; the
absence of a boolean flag automatically sets it to _false_, while the presence
automatically sets it to _true_.

#### Parameters

Parameters are any values that are not assigned to a flag. For example
`cat file1.txt file2.txt file3.txt`.

#### Allowing Flag-like Parameters

Sometimes you'll want parameters that look like flags. For example
`calculator -3 + 2` where `-3` might normally be considered a flag.

There are two ways you can force `args` to read values as a parameter:

1. **User controlled**: the user can use the `--` flag to denote that everything
   which follows is a parameter. eg `calculator -- -3 + 2`.

2. **Developer controlled**: A better option is to enable `StrictFlagPlacement`
   and have the first parameter be non-flag option. Eg
   `calculator print -3 + 2`. With `StrictFlagPlacement`, anything including
   and proceeding a parameter will always be parsed as a parameter regardless
   of whether it looks like a flag or not.

(here, `calculator` is a hypothetical command)

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`bool`](../types/bool.md):
  Boolean (primitive)
* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`str` (string)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [builtins/core/management/shell_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/shell_doc.yaml).