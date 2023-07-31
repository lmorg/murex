# `cd` - Command Reference

> Change (working) directory

## Description

Changes current working directory.

## Usage

    cd [path]

## Examples

**Home directory:**

    » cd: ~
    
Including `cd` without a parameter will also change to the current user's home
directory:

    » cd
    
**Absolute path:**

    » cd: /etc/
    
**Relative path:**

    » cd: Documents
    » cd: ./Documents

## Detail

`cd` updates an environmental variable, `$PWDHIST` with an array of paths.
You can then use that to change to a previous directory.

**View the working directory history:**

    » $PWDHIST
    
**Change to a previous directory:**

    » cd $PWDHIST[-1]
    
### auto-cd

Some people prefer to omit `cd` and just write the path, with their shell
automatically changing to that directory if the "command" is just a directory.
In Murex you can enable this behaviour by turning on "auto-cd":

    config: set shell auto-cd true

## See Also

* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`source` ](../commands/source.md):
  Import Murex code from another file of code block