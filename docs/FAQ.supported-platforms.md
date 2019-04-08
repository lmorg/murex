# Supported Operating Systems

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux

**ArchLinux:**

This is my primary development platform so all features should work.

**Debian:**

Debian is automatically tested via Docker as part of the build tests.

**Ubuntu:**

_murex_ has been tested and works well.

**CentOS:**

CentOS is automatically tested via Docker as part of the build tests.

## OS X / Darwin

Darwin is included as part of the automatic build tests however there has been
no formal usability tests. This means _murex_ will compile and execute however I
cannot make any guarantees about if all the features will function.

If you do encounter any problems then please raise an issue on Github.

## FreeBSD

_murex_ has been tested inside a 10.3-RELEASE AMD64 jail

FreeBSD support is considered very good.

## OpenBSD

Tested on an earlier release of _murex_. `regression_test.sh` cannot be run
because of `timeout` dependency however `go test` should still work and OpenBSD
is included as part of the automated build tests.

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
compile for that platform. However I have not had access to a Windows machine
for some time so there maybe some new bugs introduced in recent versions. There
is also the caveat that without a broad range of command line utilities (eg GNU
coreutils) the usefulness of _murex_ is seriously diminished. There is some work
underway to replicate some of the basics of coreutils as _murex_ builtins but
that level of work is massive, thankless, and so obviously a low priority. Thus
my recommendation is to run _murex_ inside WSL (Windows Subsystem for Linux) on
Windows 10. However if native Windows is your preference then _murex_
*should* function.

## Plan 9

Plan 9 is included as part of the automated built tests however, due to the
differences in Plan 9's syscalls, there may be a few edge case where bugs exist
in the Plan 9 build which don't with the Linux / UNIX counterparts. I don't
personally perform any functional testing for Plan 9 beyond what is already
included as part of the build and unit tests.

If you do happen to run into any such bugs then I do welcome pull requests.

# Other CPU architectures

While there isn't any CPU specific code in _murex_ below is a breakdown
of state of support and testing on alternative architectures:

## 386

Untested but should function the same as AMD64. 64 bit data types are rarely
used so performance should be roughly equal. If you do still use 32 bit
platforms and notice any issues then please raise an issue ticket and I will
investigate.

## ARM

Only tested on Linux but full compatibility was present.

## PPC

Untested however the Go compiler does support various PPC architectures so
_murex_ might work.

## MIPS

Untested however the Go compiler does support various MIPS architectures so
_murex_ might work.

## SPARC

Unsupported.
