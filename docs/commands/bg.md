# _murex_ Language Guide

## Command Reference: `bg`

> Run processes in the background

### Description

`bg` supports two modes: it can either be run as a function block which will
execute in the background, or it can take stopped processes and daemonize
them.

### Usage

Any operating system:

    bg { code block }
    
    <stdin> -> bg
    
POSIX only:

    bg { code block }
    
    <stdin> -> bg
    
    bg fid

### Examples

As a function:

    bg { sleep 5; out "Morning" }
    
As a method:

    » ({ sleep 5; out "Morning" }) -> bg

### Detail

The examples above will work on any system (Windows included). However the
`ctrl+z` usage of backgrounding a stopped process (like Bash) is only
supported on POSIX systems due to the limitation of required signals on
non-platforms. This means the usage described in the examples is cross
cross platform while `bg int` currently does not work on Windows nor Plan 9.

### See Also

* [`fg`](../commands/fg.md):
  Sends a background process into the foreground
* [fid-kill](../commands/fid-kill.md):
  
* [fid-killall](../commands/fid-killall.md):
  
* [fid-list](../commands/fid-list.md):
  