# Supported Operating Systems

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux

### ArchLinux

This is one of the primary development platforms; all features should work.

### Debian

Extensively tested

### Ubuntu

Extensively tested

### CentOS

Extensively tested

## OS X / Darwin

This is one of the primary development platforms; all features work aside alt-
hotkeys.

## FreeBSD

Extensively tested inside a 10.3-RELEASE AMD64 jail

FreeBSD support is considered very good.

## OpenBSD

Tested on an earlier version of _murex_.

OpenBSD support is expected to be good but, as always, please log an issue via
Github if you do encounter problems.

## NetBSD

NetBSD is part of the automated build tests so _murex_ will compile on NetBSD.
However no functional testing has been conducted on that particular platform.

## DragonflyBSD

DragonflyBSD is part of the automated build tests so _murex_ will compile on
DragonflyBSD. However no functional testing has been conducted on that
particular platform.

## Windows

Windows is supported and part of the automated build tests so _murex_ will
compile for that platform however very little functional test on recent
versions. There is also the caveat that without a broad range of command line
utilities (eg GNU coreutils) the usefulness of _murex_ is seriously diminished.
There is some work underway to replicate some of the basics of coreutils as
_murex_ builtins but that level of work is massive, thankless, and so obviously
a low priority. Thus the recommendation is to run _murex_ inside WSL (Windows
Subsystem for Linux) on Windows 10. However if native Windows is your preference
then _murex_ *should* function.

Feature wise, job control isn't supported in Windows because Windows doesn't
support the SIGSTSP etc signals to stop processes.

## Plan 9

Plan 9 is included as part of the automated built tests however no functional
tests have been run.

If you do happen to run into any such bugs then I do welcome pull requests.

Feature wise, job control isn't supported in Plan 9 because Plan 9 doesn't
support all of the required signals.

## Other CPU architectures

_murex_ I developed on AMD64 and that is also the architecture which runs all
of the unit tests; however there is nothing CPU specific in _murex_'s source
and the CI pipeline does spit out binaries for 386, AMD64, ARMv7 (32bit) and
ARMv8 (64bit) so the shell should compile for those architectures.

If you do happen to run into any issues then please report them on the Github
issue tracker.
