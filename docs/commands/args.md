# Define Function Arguments (`args`)

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

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex

<hr/>

This document was generated from [builtins/core/management/shell_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/shell_doc.yaml).