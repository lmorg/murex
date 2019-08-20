# _murex_ Language Guide

## Command Reference: `os`

> Output the auto-detected OS name

### Description

Output the auto-detected OS name.

### Usage

    os -> <stdout>
    
    os string -> <stdout>
    ``` 

### Examples

    » os
    linux
    
Or if you want to check if the host is one of a number of platforms:

    # When run on Linux or FreeBSD
    » os linux freebsd
    true
    
    # When run on another platform, eg Windows or Darwin (OSX)
    # (exit number would also be `1`)
    » os linux freebsd
    false
    
`posix` is also supported:

    # When run on Linux, FreeBSD or Darwin (for example)
    » os posix
    true
    
    # When run on Windows or Plan 9
    # (exit number would also be `1`)
    » os posix
    false
    
Please note that although Plan 9 shares similarities with POSIX, it is not
POSIX-compliant. For that reason, `os` returns false with the `posix`
parameter when run on Plan 9. If you want to include Plan 9 in the check
then please write it as `os posix plan9`.

### See Also

* [`cpuarch`](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [`cpucount`](../commands/cpucount.md):
  Output the number of CPU cores available on your host