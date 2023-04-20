# `man-get-flags`  - Command Reference

> Parses man page files for command line flags 

## Description

Sometimes you might want to programmatically search `man` pages for any
supported flag. Particularly if you're writing a dynamic autocompletion.
`man-get-flags` does this and returns a JSON document.

You can either pipe a man page to `man-get-flags`, or pass the name of
the command as a parameter.

## Usage

    <stdin> -> man-get-flags -> <stdout>
    
    man-get-flags command -> <stdout>

## See Also

* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for _murex_ builtins