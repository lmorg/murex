# _murex_ Shell Docs

## Command Reference: `cd`

> Change (working) directory

## Description

Changes current working directory.

## Usage

    cd path

## Examples

    # Home directory
    » cd: ~ 
    
    # Absolute path
    » cd: /etc/
    
    # Relative path
    » cd: Documents
    » cd: ./Documents

## Detail

`cd` updates an environmental variable, `$PWDHIST` with an array of paths.
You can then use that to change to a previous directory

    # View the working directory history
    » $PWDHIST
    
    # Change to a previous directory
    » cd $PWDHIST[0]

## See Also

* [commands/`source` ](../commands/source.md):
  Import _murex_ code from another file of code block
* [commands/pwd](../commands/pwd.md):
  