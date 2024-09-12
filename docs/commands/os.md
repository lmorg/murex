# Operating System (`os`)

> Output the auto-detected OS name

## Description

Output the auto-detected OS name.

## Usage

```
os -> <stdout>

os string -> <stdout>
``` 

## Examples

### Name current platform

```
» os
linux
```

### Check platforms

Or if you want to check if the host is one of a number of platforms:

```
# When run on Linux or FreeBSD
» os linux freebsd
true

# When run on another platform, eg Windows or Darwin (macOS)
# (exit number would also be `1`)
» os linux freebsd
false
```

The intention is this allows simple tests in scripts:

```
if { os windows } then {
    # run some Windows specific code
}
```

### POSIX

`posix` is also supported to check if Murex is running on a UNIX-like operating
system.

All Murex targets _apart_ from Windows and Plan 9 are considered POSIX.

```
# When run on Linux or macOS
» os posix
true

# When run on Windows or Plan 9
# (exit number would also be `1`)
» os posix
false
```

Please note that although Plan 9 shares similarities with POSIX, it is not
POSIX-compliant. For that reason, `os` returns false with the `posix`
parameter when run on Plan 9. If you want to include Plan 9 in the check
then please write it as `os posix plan9`.

## Synonyms

* `os`
* `sys.os`


## See Also

* [CPU Architecture (`cpuarch`)](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [CPU Count (`cpucount`)](../commands/cpucount.md):
  Output the number of CPU cores available on your host

<hr/>

This document was generated from [builtins/core/system/system_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/system/system_doc.yaml).